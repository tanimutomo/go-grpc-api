package main

import (
	"fmt"
	"log"
	"net"

	"github.com/pkg/errors"
	"google.golang.org/grpc"

	"github.com/tanimutomo/go-grpc-api/cmd/api/adapter/grpc/article"
)

const PORT = ":50051"

func main() {
	fmt.Println("start server")
	if err := set(); err != nil {
		log.Fatalf("%v", err)
	}
}

func set() error {
	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		return errors.Wrap(err, "failed to set port")
	}

	s := grpc.NewServer()
	article.SetHandler(s)
	if err := s.Serve(lis); err != nil {
		return errors.Wrap(err, "failed to start server")
	}

	return nil
}
