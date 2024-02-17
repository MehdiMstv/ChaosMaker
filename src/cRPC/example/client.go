package main

import (
	"context"
	"flag"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	calculator "github.com/MehdiMstv/ChaosMaker/src/cRPC/example/interface/calculator"
)

const (
	defaultName = "world"
)

type LogEntry struct {
	Timestamp primitive.DateTime           `bson:"timestamp"`
	Request   *calculator.CalculateRequest `bson:"request"`
}

var (
	addr = flag.String("addr", "localhost:50052", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	mongoClient, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://127.0.0.1:27017/"))
	if err != nil {
		log.Fatal(err)
	}
	c := calculator.NewCalculatorcRPCClient(conn, mongoClient)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	rc, err := c.Calculate(ctx, &calculator.CalculateRequest{
		Operation:    calculator.CalculateRequest_MUL,
		FirstNumber:  2,
		SecondNumber: 3,
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	rr, err := c.GetRandom(ctx, &calculator.GetRandomRequest{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Greeting: %v", rc.GetResult())
	log.Printf("Greeting: %v", rr.GetRandom())
}
