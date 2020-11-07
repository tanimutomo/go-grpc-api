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

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/namsral/flag"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/tanimutomo/grpcapi-go-server/pkg/adapter/grpc/article"
)

const serviceName = "grpcapi-go-server"

var (
	port   = flag.Int("grpcPort", 50051, "grpc port")
	server *grpc.Server
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

		server, err = initServer()
		if err != nil {
			log.Fatal(err)
			os.Exit(2)
		}

		log.Printf("gRPC server serving at %s\n", addr)

		return server.Serve(ln)
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

	if server != nil {
		log.Println("graceful stop server")
		server.GracefulStop()
	}
	err := g.Wait()
	if err != nil {
		log.Printf("server returning an error %v\n", err)
		os.Exit(2)
	}

	log.Println("all processes are finished")
}

func initServer() (s *grpc.Server, err error) {
	logger, err := zap.NewDevelopment()
	if err != nil {
		return s, errors.New("failed to create zap logger")
	}
	s = grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_opentracing.UnaryServerInterceptor(),
			grpc_prometheus.UnaryServerInterceptor,
			grpc_zap.UnaryServerInterceptor(logger),
			grpc_recovery.UnaryServerInterceptor(),
		)),
	)
	reflection.Register(s)

	article.SetHandler(s)

	return s, nil
}
