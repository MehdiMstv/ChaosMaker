// Code generated by protoc-gen-crpc. DO NOT EDIT.
// source: example/interface/calculator.proto
package calculator

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

type calculate1RequestEntry struct {
	Timestamp time.Time           `bson:"timestamp"`
	Request   *Calculator1Request `bson:"request"`
}

type calculate2RequestEntry struct {
	Timestamp time.Time           `bson:"timestamp"`
	Request   *Calculator2Request `bson:"request"`
}

type CalculatorcRPCClient interface {
	Calculate1(ctx context.Context, in *Calculator1Request, opts ...grpc.CallOption) (*CalculatorResponse, error)
	Calculate2(ctx context.Context, in *Calculator2Request, opts ...grpc.CallOption) (*CalculatorResponse, error)
}
type calculatorcRPCClient struct {
	client CalculatorClient
	db     *mongo.Client
}

func NewCalculatorcRPCClient(cc grpc.ClientConnInterface, db *mongo.Client) CalculatorcRPCClient {
	client := NewCalculatorClient(cc)
	return &calculatorcRPCClient{
		client: client,
		db:     db,
	}
}

func (s *calculatorcRPCClient) Calculate1(ctx context.Context, req *Calculator1Request, opts ...grpc.CallOption) (*CalculatorResponse, error) {
	// Log the request to MongoDB async
	go s.db.Database("Calculator").Collection("Calculate1").InsertOne(ctx, &calculate1RequestEntry{
		Timestamp: time.Now(),
		Request:   req,
	})

	// Invoke the original RPC method
	resp, err := s.client.Calculate1(ctx, req)

	// Handle response and error
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *calculatorcRPCClient) Calculate2(ctx context.Context, req *Calculator2Request, opts ...grpc.CallOption) (*CalculatorResponse, error) {
	// Log the request to MongoDB async
	go s.db.Database("Calculator").Collection("Calculate2").InsertOne(ctx, &calculate2RequestEntry{
		Timestamp: time.Now(),
		Request:   req,
	})

	// Invoke the original RPC method
	resp, err := s.client.Calculate2(ctx, req)

	// Handle response and error
	if err != nil {
		return nil, err
	}
	return resp, nil
}
