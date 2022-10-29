package controllers

import (
	"context"
	"net/http"

	"github.com/NeoTRAN001/go-ginframework/models"
	"github.com/NeoTRAN001/go-ginframework/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

/*
	El controlador va a implementar los servicios que sean necesarios, estos declarados
	en el estruct, usando inyección de dependencias
*/

type UserController struct {
	UserService services.UserService
}

func NewUserController(ctxDB context.Context, mongoClient *mongo.Client) UserController {
	return UserController{
		UserService: services.NewUserService(mongoClient, ctxDB),
	}
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := c.UserService.CreateUser(&user); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (c *UserController) GetUser(ctx *gin.Context) {

	username := ctx.Param("name")

	user, err := c.UserService.GetUser(&username)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (c *UserController) GetAll(ctx *gin.Context) {
	users, err := c.UserService.GetAll()

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func (c *UserController) UpdateUser(ctx *gin.Context) {

	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := c.UserService.UpdateUser(&user); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (c *UserController) DeleteUser(ctx *gin.Context) {
	username := ctx.Param("name")

	if err := c.UserService.DeleteUser(&username); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

/*
Función para crear rutas de un controlador
*/
func (c *UserController) RegisterUserRoutes(rg *gin.RouterGroup) {
	userRoute := rg.Group("/user")
	userRoute.GET("/get/:name", c.GetUser)
	userRoute.GET("/getAll", c.GetAll)
	userRoute.POST("/create", c.CreateUser)
	userRoute.PATCH("/update", c.UpdateUser)
	userRoute.DELETE("/delete/:name", c.DeleteUser)
}
