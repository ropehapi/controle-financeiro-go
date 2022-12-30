package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ropehapi/controle-financeiro-go/models"
)

func Index(c *gin.Context) {
	account := models.Account{Name:   "Inter",Amount: 2684.15,}
	c.JSON(200, account)
}
