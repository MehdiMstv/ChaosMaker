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
	g.P("// Code generated by protoc-gen-chaos-maker.")
	g.P("// source: ", file.Desc.Path())
	g.P("package ", file.GoPackageName)
	g.P()
	g.P("import (")
	g.P(`    "context"`)
	g.P(`    "encoding/json"`)
	g.P(`    "fmt"`)
	g.P(`    "log"`)
	g.P(`    "net/http"`)
	g.P()
	g.P(`    "github.com/gin-gonic/gin"`)
	g.P(`    "go.mongodb.org/mongo-driver/bson"`)
	g.P(`    "go.mongodb.org/mongo-driver/bson/primitive"`)
	g.P(`    "go.mongodb.org/mongo-driver/mongo"`)
	g.P(`    "go.mongodb.org/mongo-driver/mongo/options"`)
	g.P(`    "google.golang.org/grpc"`)
	g.P(`    "google.golang.org/grpc/credentials/insecure"`)
	g.P(`    "google.golang.org/grpc/status"`)
	g.P(")")
	g.P()
	g.P("type config struct {")
	g.P("	LoggerMongodbURI string")
	g.P("	ControlPlaneURL string")
	g.P("	StagingAddress string `json:\"staging_address\"`")
	g.P("	ServiceName string")
	g.P("	Port string")
	g.P("}")
	g.P()
	for _, method := range service.Methods {
		g.P("type ", toLower(method.GoName), "RequestEntry struct {")
		g.P("	Timestamp primitive.DateTime            `bson:\"timestamp\"`")
		g.P("	Request   *calculator.", method.Input.GoIdent, " `bson:\"request\"`")
		g.P("}")
		g.P()
	}
	g.P()
	g.P("func RunServer(db *mongo.Client, conn *grpc.ClientConn, c *config) {")
	g.P("	router := gin.Default()")
	for _, method := range service.Methods {
		g.P("	router.POST(\"/start_", method.GoName, "_chaos\", handle", method.GoName, "Chaos(db, conn, c))")
	}
	g.P("	err := router.Run(\":\" + c.Port)")
	g.P("	if err != nil {")
	g.P("		return ")
	g.P("	}")
	g.P("}")
	g.P()
	g.P("func getRequests(db *mongo.Client, methodName string) (*mongo.Cursor, error){")
	g.P(`	requests, err := db.Database("`, service.GoName, `").Collection(methodName).Find(context.Background(), bson.D{})`)
	g.P("	if err != nil {")
	g.P("		return nil, err")
	g.P("	}")
	g.P("return requests, nil")
	g.P("}")
	g.P()
	for _, method := range service.Methods {
		g.P("func handle", method.GoName, "Chaos(db *mongo.Client, conn *grpc.ClientConn, config *config) gin.HandlerFunc {")
		g.P("	fn := func (c *gin.Context){")
		g.P(`		chaosID := c.Request.FormValue("id")`)
		g.P("		var data []", toLower(method.GoName), "RequestEntry")
		g.P()
		g.P(`		http.Post(fmt.Sprintf("http://%s/api/chaos?id=%s", config.ControlPlaneURL, chaosID), "application/json", nil)`)
		g.P()
		g.P(`		filters, _ := getRequests(db, "`, method.GoName, `")`)
		g.P("		err := filters.All(context.Background(), &data)")
		g.P("		if err != nil {")
		g.P("			c.JSON(500, gin.H{\"error\": err.Error()})")
		g.P("			return")
		g.P("		}")
		g.P()
		g.P("		client := ", file.GoPackageName, ".New", service.GoName, "Client(conn)")
		g.P(`		http.Post(fmt.Sprintf("http://%s/api/chaos?id=%s", config.ControlPlaneURL, chaosID), "application/json", nil)`)
		g.P()
		g.P(`		resultData := make(map[string]int)`)
		g.P("		for _, v := range data {")
		g.P("			_, err := client.", method.GoName, "(context.Background(), v.Request)")
		g.P("			if err != nil {")
		g.P("				if s, ok := status.FromError(err); ok {")
		g.P("					resultData[s.Code().String()] = resultData[s.Code().String()] + 1")
		g.P("				} else {")
		g.P(`					resultData["Unknown"] = resultData["Unknown"] + 1`)
		g.P("				}")
		g.P("				continue")
		g.P("			}")
		g.P(`			resultData["Success"] = resultData["Success"] + 1`)
		g.P("		}")
		g.P()
		g.P("		jsonString, _ := json.Marshal(resultData)")
		g.P(`		http.Post(fmt.Sprintf("http://%s/api/chaos?id=%s", config.ControlPlaneURL, chaosID), "application/json", bytes.NewBuffer(jsonString))`)
		g.P(`		c.String(http.StatusOK, "Chaos Done")`)
		g.P("	}")
		g.P()
		g.P("	return fn")
		g.P("	}")
		g.P()
	}
	g.P("func getStagingAddress(c *config) error {")
	g.P(`	response, err := http.Get(fmt.Sprintf("http://%s/api/service/staging_address?name=%s", c.ControlPlaneURL, c.ServiceName))`)
	g.P("	if err != nil {")
	g.P(`		return err`)
	g.P("	}")
	g.P("	if response.StatusCode != 200 {")
	g.P(`		return fmt.Errorf("failed to get staging address")`)
	g.P("	}")
	g.P()
	g.P("	body, err := io.ReadAll(response.Body)")
	g.P("	if err != nil {")
	g.P("		return err")
	g.P("	}")
	g.P()
	g.P("	err = json.Unmarshal(body, &c)")
	g.P("	if err != nil {")
	g.P("		return err")
	g.P("	}")
	g.P()
	g.P(`	return nil`)
	g.P("}")
	g.P()
	g.P("// main function of program, you can change config values here")
	g.P("func main() {")
	g.P("	c := &config{")
	g.P(`		LoggerMongodbURI: "mongodb://127.0.0.1:27017/",`)
	g.P(`		ControlPlaneURL: "127.0.0.1:9033",`)
	g.P(`		ServiceName: "`, service.GoName, `",`)
	g.P(`		Port: "8082",`)
	g.P("	}")
	g.P()
	g.P("	err := getStagingAddress(c)")
	g.P("	if err != nil {")
	g.P("		log.Fatal(err)")
	g.P("		return")
	g.P("	}")
	g.P()
	g.P(`	mongoClient, err := mongo.Connect(context.Background(), options.Client().ApplyURI(c.LoggerMongodbURI))`)
	g.P("	if err != nil {")
	g.P("		log.Fatal(err)")
	g.P("		return")
	g.P("	}")
	g.P()
	g.P("	conn, err := grpc.Dial(c.StagingAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))")
	g.P("	if err != nil {")
	g.P("		log.Fatal(err)")
	g.P("		return")
	g.P("}")
	g.P()
	g.P("RunServer(mongoClient, conn, c)")
	g.P("}")
}

func toLower(s string) string { return strings.ToLower(s[:1]) + s[1:] }
