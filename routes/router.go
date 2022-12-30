package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ropehapi/controle-financeiro-go/controllers"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/account", controllers.Index)
	r.GET("/account/:id", controllers.Show)
	r.POST("/account", controllers.Store)
	r.PATCH("/account/:id", controllers.Update)
	r.DELETE("/account/:id", controllers.Delete)
	r.Run()
}
