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
	server         *gin.Engine
	userController controllers.UserController
	ctx            context.Context
	dbClient       *mongo.Client
	err            error
)

func routes() {
	basePath := server.Group("/api")

	userController = controllers.NewUserController(ctx, dbClient)
	userController.RegisterUserRoutes(basePath)
}

func init() {
	ctx = context.TODO()
	mongoConn := options.Client().ApplyURI("mongodb+srv://")
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
