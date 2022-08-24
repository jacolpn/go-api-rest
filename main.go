package main

import (
	"fmt"

	"github.com/jacolpn/go-api-rest/database"
	"github.com/jacolpn/go-api-rest/models"
	"github.com/jacolpn/go-api-rest/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome 1", Historia: "História 1"},
		{Id: 2, Nome: "Nome 2", Historia: "História 2"},
	}

	database.ConectaComBancoDeDados()
	fmt.Println("Iniciando o servidor Rest em Go")
	routes.HandleRequest()
}
