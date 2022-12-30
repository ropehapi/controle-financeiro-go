package database

import (
	"github.com/ropehapi/controle-financeiro-go/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetConn() *gorm.DB{
	dsn := "root:159357@tcp(localhost:3306)/db_controle_financeiro_go?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&models.Account{})

	return db
}
