package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc.com/learn/procpb"
	"log"
	"net"
)

type server struct{}

func (c *server) Greet(ctx context.Context, request *procpb.GreetRequest) (*procpb.GreetResponse, error) {

	fmt.Println("Greet was invoked by client")

	name := request.GetGreeting().GetName() // name = 1;

	result := &procpb.GreetResponse{
		Result: "Welcome " + name,
	}

	return result, nil
}

func main() {

	fmt.Println("Server has started")

	conn, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		fmt.Println("cannot connect by TCP")
	}

	s := grpc.NewServer()

	procpb.RegisterProcServiceServer(s, &server{})


	if err := s.Serve(conn); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
	}
}
