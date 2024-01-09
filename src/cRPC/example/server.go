package main

import (
	"context"
	"flag"
	"fmt"
	calculator2 "github.com/MehdiMstv/ChaosMaker/src/cRPC/example/interface/calculator"
	"log"
	"net"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	calculator2.UnimplementedCalculatorServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) Calculate1(ctx context.Context, in *calculator2.CalculatorRequest) (*calculator2.CalculatorResponse, error) {
	return &calculator2.CalculatorResponse{Result: 10}, nil
}
func (s *server) Calculate2(ctx context.Context, in *calculator2.CalculatorRequest) (*calculator2.CalculatorResponse, error) {
	return &calculator2.CalculatorResponse{Result: 10}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	calculator2.RegisterCalculatorServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
