package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ropehapi/controle-financeiro-go/controllers"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/", controllers.HelloWorld)
	r.Run()
}
