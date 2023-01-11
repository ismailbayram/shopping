package main

import (
	"fmt"
	"github.com/ismailbayram/shopping/config"
	"github.com/ismailbayram/shopping/pkg/api"
	"log"
)

func main() {
	cfg := config.Init()

	router := api.NewRouter()
	err := router.Run(fmt.Sprintf(":%s", cfg.Server.Port))
	if err != nil {
		log.Fatalln(err)
	}
}
