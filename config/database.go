package config

import (
	"github.com/ropehapi/controle-financeiro-go/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func GetConn() {
	dsn := "root:159357@tcp(localhost:3306)/db_controle_financeiro_go?charset=utf8&parseTime=True&loc=Local"
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	DB.AutoMigrate(&models.Account{})
}
