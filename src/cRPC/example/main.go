package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/MehdiMstv/ChaosMaker/src/cRPC/example/interface/calculator"
)

var serviceName = "chaos"

type flagData struct {
	MongodbURI string `json:"mongodb_uri"`
	F1         bool   `json:"F1"`
}

type Calculate1RequestEntry struct {
	Timestamp primitive.DateTime             `bson:"timestamp"`
	Request   *calculator.Calculator1Request `bson:"request"`
}

type Calculate2RequestEntry struct {
	Timestamp primitive.DateTime             `bson:"timestamp"`
	Request   *calculator.Calculator2Request `bson:"request"`
}

func RunServer(db *mongo.Client, conn *grpc.ClientConn) {
	router := gin.Default()
	router.POST("/start_Calculate1_chaos", handleCalculate1Chaos(db, conn))
	router.POST("/start_Calculate2_chaos", handleCalculate2Chaos(db, conn))
	router.Run(":8080")
}

func getRequests(db *mongo.Client, methodName string) (*mongo.Cursor, error) {
	requests, err := db.Database("logs").Collection(methodName).Find(context.Background(), bson.D{})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return requests, nil
}

func handleCalculate1Chaos(db *mongo.Client, conn *grpc.ClientConn) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		chaosID := c.Request.FormValue("id")
		fmt.Println("Test1", chaosID)
		get, err := http.Get(fmt.Sprintf("http://127.0.0.1:9033/api/chaoses?chaos_id=%s", chaosID))
		fmt.Println(get)
		if err != nil {
			fmt.Println(err)
			return
		}
		var data []Calculate1RequestEntry
		filters, _ := getRequests(db, "handleCalculate1Chaos")
		fmt.Println("Test2")
		filters.All(context.Background(), &data)
		client := calculator.NewCalculatorClient(conn)
		http.Get(fmt.Sprintf("http://127.0.0.1:9033/api/chaoses?chaos_id=%s", chaosID))
		for _, v := range data {
			response, _ := client.Calculate1(context.Background(), v.Request)
			fmt.Println(response)
		}
		http.Get(fmt.Sprintf("http://127.0.0.1:9033/api/chaoses?chaos_id=%s", chaosID))
		http.Get(fmt.Sprintf("http://127.0.0.1:9033/api/chaoses?chaos_id=%s", chaosID))
		c.String(http.StatusOK, "Chaos Done")
	}
	return fn
}

func handleCalculate2Chaos(db *mongo.Client, conn *grpc.ClientConn) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var data []Calculate2RequestEntry
		filters, _ := getRequests(db, "handleCalculate2Chaos")
		filters.All(context.Background(), &data)
		client := calculator.NewCalculatorClient(conn)
		for _, v := range data {
			response, _ := client.Calculate2(context.Background(), v.Request)
			fmt.Println(response)
		}
		c.String(http.StatusOK, "Chaos created")
	}
	return fn
}

func readFlags(flags *flagData) {
	for {
		response, err := http.Get(fmt.Sprintf("http://127.0.0.1:9033/api/flags?service_name=%s", serviceName))
		if err != nil {
			time.Sleep(10 * time.Second)
			continue
		}
		decoder := json.NewDecoder(response.Body)
		err = decoder.Decode(&flags)
		if err != nil {
			time.Sleep(10 * time.Second)
			continue
		}
		fmt.Println(flags.MongodbURI)
		fmt.Println(flags.F1)
		time.Sleep(10 * time.Second)
	}
}

func main() {
	flags := &flagData{
		MongodbURI: "mongodb://127.0.0.1:27017/",
	}
	go readFlags(flags)
	mongoClient, err := mongo.Connect(context.Background(), options.Client().ApplyURI(flags.MongodbURI))
	if err != nil {
		log.Fatal(err)
	}
	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	RunServer(mongoClient, conn)
}
