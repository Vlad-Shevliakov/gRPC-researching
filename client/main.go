package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc.com/learn/procpb"
)

func main() {
	fmt.Println("client here")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		fmt.Println("fail to dial connection")
	}

	defer conn.Close()

	client := procpb.NewProcServiceClient(conn)

	request := &procpb.GreetRequest{
		Greeting: &procpb.Greeting{
			Name:    "Vlad",
			Network: "Facebook",
		},
	}

	response, err := client.Greet(context.Background(), request)
	if err != nil {
		fmt.Println("fail to ger Greet response")
	}

	fmt.Printf("%v\n", response.Result)

}
