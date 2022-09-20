package main

import (
	"fmt"

	"github.com/CaioLuColaco/api-golang-rest/database"
	"github.com/CaioLuColaco/api-golang-rest/models"
	"github.com/CaioLuColaco/api-golang-rest/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome 1", Historia: "Historia 1"},
		{Id: 2, Nome: "Nome 2", Historia: "Historia 2"},
	}
	database.ConectaComBancoDeDados()
	fmt.Println("Iniciando servidor Rest com go")
	routes.HandleRequest()
}