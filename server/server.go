package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc/credentials"

	greetpb "github.com/nicewook/grpc-ssltls/proto"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Greet(ctx context.Context, req *greetpb.GreetReq) (*greetpb.GreetRes, error) {
	name := req.GetName()
	result := &greetpb.GreetRes{
		Greeting: "Hi, Nice to meet you, " + name,
	}
	return result, nil
}

func serve() {
	l, err := net.Listen("tcp", "127.0.0.1:50055")
	if err != nil {
		log.Fatalf("fail to net listen: %v", err)
	}

	var opts []grpc.ServerOption
	tls := true
	if tls {
		fmt.Println("In Secure mode")
		serverCrt := "keys/server.crt"
		serverPem := "keys/server.pem"
		creds, sslErr := credentials.NewServerTLSFromFile(serverCrt, serverPem)
		if sslErr != nil {
			log.Fatalf("fail to load creds: %v", sslErr)
		}
		opts = append(opts, grpc.Creds(creds))
	}
	s := grpc.NewServer(opts...)

	greetpb.RegisterGreetServiceServer(s, &server{})
	serveErr := s.Serve(l)
	if serveErr != nil {
		log.Fatalf("fail to serve gRPC server: %v", serveErr)
	}
}
func main() {
	fmt.Println("Greet gRPC server starts...")
	serve()
}
