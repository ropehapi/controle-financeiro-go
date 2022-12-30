package models

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	Name   string  `json:"name"`
	Amount float64 `json:"amount"`
}

var Accounts []Account
