package main

import (
	"log"

	"deliverables/config"
	"deliverables/controllers"
	"deliverables/routes"
	"deliverables/services"
	"deliverables/usecases"
)

func main() {
	config := config.ReadConfig()
	s, err := services.New(config.ExternalUrl.GetPokemonById, config.ExternalUrl.TimeOut)
	if err != nil {
		log.Fatal(err)
	}
	u := usecases.New(s)
	c := controllers.New(u)
	routes.New(c)
}
