package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ropehapi/controle-financeiro-go/database"
	"github.com/ropehapi/controle-financeiro-go/models"
)

func Index(c *gin.Context) {
	db := database.GetConn()

	var accounts []models.Account
	db.Find(&accounts)
	c.JSON(200, accounts)
}

func Show(c *gin.Context){
	db := database.GetConn()
	
	var account models.Account
	id := c.Params.ByName("id")
	db.First(&account, id)

	if account.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Account not found"})
		return
	}

	c.JSON(http.StatusOK, account)
}

func Store(c *gin.Context) {
	db := database.GetConn()

	var account models.Account
	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	db.Create(&account)
	c.JSON(http.StatusOK, account)
}

func Update(c *gin.Context){
	db := database.GetConn()

	var account models.Account
	id := c.Params.ByName("id")
	db.First(&account, id)

	if account.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Account not found"})
		return
	}

	if err := c.ShouldBindJSON(&account); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	db.Model(&account).Where(account.ID, id).UpdateColumns(account)
	c.JSON(http.StatusOK, account)
}

 func Delete(c *gin.Context){
	db := database.GetConn()

	var account models.Account
	id := c.Params.ByName("id")
	db.First(&account, id)

	if account.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Account not found"})
		return
	}

	db.Delete(&account, id)
	c.JSON(http.StatusOK, account)
 }