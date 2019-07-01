package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc/credentials"

	greetpb "github.com/nicewook/grpc-ssltls/proto"

	"google.golang.org/grpc"
)

func main() {

	fmt.Println("Greet gRPC client starts")

	opts := grpc.WithInsecure()
	tls := true
	if tls {
		fmt.Println("In Secure mode")
		rootCACert := "keys/rootca.crt"
		creds, sslErr := credentials.NewClientTLSFromFile(rootCACert, "")
		if sslErr != nil {
			log.Fatalf("fail to load CA trust certificate: %v", sslErr)
		}
		opts = grpc.WithTransportCredentials(creds)
	}

	cc, err := grpc.Dial("localhost:50055", opts)
	if err != nil {
		log.Fatalf("fail to dial to gRPC server: %v", err)
	}
	c := greetpb.NewGreetServiceClient(cc)
	res, err := c.Greet(context.Background(), &greetpb.GreetReq{Name: "JHS"})
	if err != nil {
		log.Fatalf("Fail to greet: %v", err)
	}
	fmt.Println(res.GetGreeting())
}
