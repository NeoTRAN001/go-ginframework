package main

import (
	"context"
	"fmt"
	"log"

	"github.com/NeoTRAN001/go-ginframework/controllers"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server          *gin.Engine
	ctx             context.Context
	dbClient        *mongo.Client
	err             error
	userController  controllers.UserController
	helloController controllers.HelloController
)

func routes() {
	basePath := server.Group("/api")

	userController = controllers.NewUserController(ctx, dbClient)
	userController.RegisterUserRoutes(basePath)

	helloController = controllers.NewHelloController()
	helloController.RegisterHelloRoutes(basePath)
}

func init() {
	ctx = context.TODO()
	mongoConn := options.Client().ApplyURI("mongodb+srv://neotran:123abc@cluster0.wlfkhkk.mongodb.net/?retryWrites=true&w=majority")
	dbClient, err = mongo.Connect(ctx, mongoConn)

	if err != nil {
		log.Fatal(err)
	}

	if err := dbClient.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Mongo connection established")

	server = gin.Default()
}

func main() {
	defer dbClient.Disconnect(ctx)

	routes()

	log.Fatal(server.Run(":9090"))
}
