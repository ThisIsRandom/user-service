package main

import (
	"fmt"
	"net"

	"os"

	"github.com/thisisrandom/user-service/proto"
	"google.golang.org/grpc"
)

func main() {
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", PORT))

	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	proto.RegisterTokenServer(s, &proto.Service{})

	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
