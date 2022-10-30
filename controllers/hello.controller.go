package controllers

import (
	"net/http"

	"github.com/NeoTRAN001/go-ginframework/services"
	"github.com/gin-gonic/gin"
)

type HelloController struct {
	helloService services.HelloService
}

func NewHelloController() HelloController {
	return HelloController{
		helloService: services.NewHelloService(),
	}
}

func (c *HelloController) getHello(ctx *gin.Context) {

	message, err := c.helloService.GetHello()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, message)
}

func (c *HelloController) RegisterHelloRoutes(rg *gin.RouterGroup) {
	route := rg.Group("/hello")

	route.GET("", c.getHello)
}
