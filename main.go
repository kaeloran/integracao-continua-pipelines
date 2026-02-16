package main

import (
	"integracao-continua-pipelines/database"
	"integracao-continua-pipelines/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequest()
}
