package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"

	"google.golang.org/grpc/reflection"

	calculator "github.com/MehdiMstv/ChaosMaker/src/cRPC/example/interface/calculator"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50052, "The server port")
)

type server struct {
	calculator.UnimplementedCalculatorServer
	flags *calculator.FlagData
}

func (s *server) Calculate(ctx context.Context, in *calculator.CalculateRequest) (*calculator.CalculateResponse, error) {
	result := int64(0)
	switch in.GetOperation().String() {
	case calculator.CalculateRequest_SUM.String():
		result = in.GetFirstNumber() + in.GetSecondNumber()
	case calculator.CalculateRequest_SUB.String():
		result = in.GetFirstNumber() - in.GetSecondNumber()
	case calculator.CalculateRequest_MUL.String():
		result = in.GetFirstNumber() * in.GetSecondNumber()
	case calculator.CalculateRequest_DIV.String():
		result = in.GetFirstNumber() / in.GetSecondNumber()
	default:
		return nil, errors.New("invalid operation")
	}
	return &calculator.CalculateResponse{Result: result}, nil
}
func (s *server) GetRandom(ctx context.Context, in *calculator.GetRandomRequest) (*calculator.GetRandomResponse, error) {
	time.Sleep(s.flags.SleepTime * time.Second)
	return &calculator.GetRandomResponse{Random: int64(rand.Int())}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	reflection.Register(s)
	c := calculator.CRPCConfig{
		FlagData:        &calculator.FlagData{},
		IsStaging:       true,
		ServiceName:     "Calculator",
		ControlPlaneURL: "127.0.0.1:9033",
	}
	calculator.RegisterCalculatorCRPCServer(s, &server{flags: c.FlagData}, &c)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
