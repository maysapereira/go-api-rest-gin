package main

import (
	"github.com/maysapereira/go-api-rest-gin/database"
	"github.com/maysapereira/go-api-rest-gin/models"
	"github.com/maysapereira/go-api-rest-gin/routes"
)

func main() {
	database.ConectaBancoDeDados()
	models.Alunos = []models.Aluno{
		{Nome: "Maysa Pereira", CPF: "123.456.789-01", RG: "98.765.43"},
		{Nome: "Kasper", CPF: "987.654.321-09", RG: "12.345.67"},
	}
	routes.HandleRequests()
}