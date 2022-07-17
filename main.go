package main

import (
	"github.com/maysapereira/go-api-rest-gin/database"
	"github.com/maysapereira/go-api-rest-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}