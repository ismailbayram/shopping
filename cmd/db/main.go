package main

import (
	"github.com/ismailbayram/shopping/config"
	mediaInfrastructure "github.com/ismailbayram/shopping/internal/media/infrastructure/db"
	"github.com/ismailbayram/shopping/pkg/database"
)

func main() {
	cfg := config.Init()

	db := database.New(&cfg.Database)
	db.Migrate([]interface{}{
		mediaInfrastructure.ImageDB{},
	})
}
