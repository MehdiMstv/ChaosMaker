package main

import (
	"flag"
	"google.golang.org/protobuf/compiler/protogen"
	"strings"
)

func main() {
	var flags flag.FlagSet
	protogen.Options{
		ParamFunc: flags.Set,
	}.Run(func(gen *protogen.Plugin) error {
		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}
			generateChaosFile(gen, f)
		}
		return nil
	})
}

func generateChaosFile(gen *protogen.Plugin, file *protogen.File) {
	filename := file.GeneratedFilenamePrefix + "_chaos.pb.go"
	g := gen.NewGeneratedFile(filename, file.GoImportPath)

	service := file.Services[0]
	g.P("// Code generated by protoc-gen-chaos-maker. DO NOT EDIT.")
	g.P("// source: ", file.Desc.Path())
	g.P("package ", file.GoPackageName)
	g.P()
	g.P("import (")
	g.P(`    "context"`)
	g.P(`    "encoding/json"`)
	g.P(`    "flag"`)
	g.P(`    "fmt"`)
	g.P(`    "log"`)
	g.P(`    "net/http"`)
	g.P(`    "time"`)
	g.P()
	g.P(`    "github.com/gin-gonic/gin"`)
	g.P(`    "go.mongodb.org/mongo-driver/bson"`)
	g.P(`    "go.mongodb.org/mongo-driver/bson/primitive"`)
	g.P(`    "go.mongodb.org/mongo-driver/mongo"`)
	g.P(`    "go.mongodb.org/mongo-driver/mongo/options"`)
	g.P(`    "google.golang.org/grpc"`)
	g.P(`    "google.golang.org/grpc/credentials/insecure"`)
	g.P()
	g.P("	", file.GoPackageName, " ", file.GoImportPath)
	g.P(")")
	g.P()
	g.P(`var serviceName = "chaos"`)
	g.P()
	g.P("type flagData struct {")
	g.P("	MongodbURI string `json:\"mongodb_uri\"`")
	g.P("}")
	g.P()
	for _, method := range service.Methods {
		g.P("type ", toLower(method.GoName), "RequestEntry struct {")
		g.P("	Timestamp primitive.DateTime            `bson:\"timestamp\"`")
		g.P("	Request   *calculator.", method.Input.GoIdent, " `bson:\"request\"`")
		g.P("}")
		g.P()
	}
	g.P("var (")
	g.P("	addr = flag.String(\"addr\", \"localhost:50051\", \"the address to connect to\")")
	g.P(")")
	g.P()
	g.P("func RunServer(db *mongo.Client, conn *grpc.ClientConn) {")
	g.P("	router := gin.Default()")
	for _, method := range service.Methods {
		g.P("	router.POST(\"/start_", method.GoName, "_chaos\", handle", method.GoName, "Chaos(db, conn))")
	}
	g.P("	router.Run(\":8080\")")
	g.P("}")
	g.P()
	g.P("func getRequests(db *mongo.Client, methodName string) (*mongo.Cursor, error){")
	g.P(`	requests, err := db.Database("logs").Collection(methodName).Find(context.Background(), bson.D{})`)
	g.P("	if err != nil {")
	g.P("		fmt.Println(err)")
	g.P("		return nil, err")
	g.P("	}")
	g.P("return requests, nil")
	g.P("}")
	g.P()
	for _, method := range service.Methods {
		g.P("func handle", method.GoName, "Chaos(db *mongo.Client, conn *grpc.ClientConn) gin.HandlerFunc {")
		g.P("	fn := func (c *gin.Context){")
		g.P(`		chaosID := c.Request.FormValue("id")`)
		g.P(`		http.Get(fmt.Sprintf("http://127.0.0.1:9033/api/chaos?chaos_id=%s", chaosID))`)
		g.P("		var data []", method.GoName, "RequestEntry")
		g.P(`		filters, _ := getRequests(db, "handle`, method.GoName, `Chaos")`)
		g.P("		filters.All(context.Background(), &data)")
		g.P("		client := ", file.GoPackageName, ".New", service.GoName, "Client(conn)")
		g.P(`		http.Get(fmt.Sprintf("http://127.0.0.1:9033/api/chaos?chaos_id=%s", chaosID))`)
		g.P("		for _, v := range data {")
		g.P("			response, _ := client.", method.GoName, "(context.Background(), v.Request)")
		g.P("			fmt.Println(response)")
		g.P("		}")
		g.P(`		http.Get(fmt.Sprintf("http://127.0.0.1:9033/api/chaos?chaos_id=%s", chaosID))`)
		g.P(`		http.Get(fmt.Sprintf("http://127.0.0.1:9033/api/chaos?chaos_id=%s", chaosID))`)
		g.P(`		c.String(http.StatusOK, "Chaos created")`)
		g.P("	}")
		g.P("	return fn")
		g.P("	}")
		g.P()
	}
	g.P("func readFlags(flags *flagData) {")
	g.P("	for {")
	g.P(`		response, err := http.Get(fmt.Sprintf("http://127.0.0.1:9033/api/flags?service_name=%s", serviceName))`)
	g.P("		if err != nil {")
	g.P("			time.Sleep(10 * time.Second)")
	g.P("			continue")
	g.P("		}")
	g.P("		decoder := json.NewDecoder(response.Body)")
	g.P("		err = decoder.Decode(&flags)")
	g.P("		if err != nil {")
	g.P("			time.Sleep(10 * time.Second)")
	g.P("			continue")
	g.P("		}")
	g.P("		time.Sleep(10 * time.Second)")
	g.P("	}")
	g.P("}")
	g.P("func main() {")
	g.P("	flags := &flagData{")
	g.P(`		MongodbURI: "mongodb://127.0.0.1:27017/",`)
	g.P("	}")
	g.P("	go readFlags(flags)")
	g.P(`	mongoClient, err := mongo.Connect(context.Background(), options.Client().ApplyURI(flags.MongodbURI))`)
	g.P("	if err != nil {")
	g.P("		log.Fatal(err)")
	g.P("	}")
	g.P("	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))")
	g.P("	if err != nil {")
	g.P("		log.Fatal(err)")
	g.P("}")
	g.P("RunServer(mongoClient, conn)")
	g.P("}")
}

func toLower(s string) string { return strings.ToLower(s[:1]) + s[1:] }
