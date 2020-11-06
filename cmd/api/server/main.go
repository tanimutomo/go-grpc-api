package main

import (
	"fmt"
	"log"
	"net"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/tanimutomo/grpcapi-go-server/cmd/api/adapter/grpc/article"
)

const port = ":50051"

func main() {
	fmt.Println("start server")
	if err := set(); err != nil {
		log.Fatalf("%v", err)
	}
}

func set() error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		return errors.Wrap(err, "failed to set port")
	}

	s := grpc.NewServer()
	article.SetHandler(s)
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		return errors.Wrap(err, "failed to start server")
	}

	return nil
}
