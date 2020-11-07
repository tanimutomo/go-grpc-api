package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/namsral/flag"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"

	"github.com/tanimutomo/grpcapi-go-server/pkg/server"
)

const serviceName = "grpcapi-go-server"

var (
	port = flag.Int("grpcPort", 50051, "grpc port")
	gs   *grpc.Server
)

func main() {
	flag.Parse()

	log.Println("starting app.")

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(interrupt)

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		addr := fmt.Sprintf(":%d", *port)
		ln, err := net.Listen("tcp", addr)
		if err != nil {
			log.Printf("gRPC server: failed to listen: %v\n", err)
			os.Exit(2)
		}

		gs, err = server.InitGrpcServer()
		if err != nil {
			log.Fatal(err)
			os.Exit(2)
		}

		log.Printf("gRPC server serving at %s\n", addr)

		return gs.Serve(ln)
	})

	select {
	case <-interrupt:
		break
	case <-ctx.Done():
		break
	}

	log.Println("received shutdown signal")

	cancel()

	_, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	if gs != nil {
		log.Println("graceful stop server")
		gs.GracefulStop()
	}
	err := g.Wait()
	if err != nil {
		log.Printf("server returning an error %v\n", err)
		os.Exit(2)
	}

	log.Println("all processes are finished")
}
