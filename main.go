package main

import (
	"github.com/ropehapi/controle-financeiro-go/config"
	"github.com/ropehapi/controle-financeiro-go/routes"
)

func main() {
	config.GetConn()
	routes.HandleRequests()
}
