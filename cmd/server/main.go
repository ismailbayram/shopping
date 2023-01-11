package main

import (
	"fmt"
	"github.com/ismailbayram/shopping/config"
	"github.com/ismailbayram/shopping/internal/database"
	"github.com/ismailbayram/shopping/internal/users"
	"github.com/ismailbayram/shopping/pkg/api"
	"log"
	"net/http"
	"time"
)

func main() {
	cfg := config.Init()

	db := database.New(&cfg.Database)

	app := &api.App{
		Users: users.New(db.Conn),
	}

	s := &http.Server{
		Addr:           fmt.Sprintf(":%s", cfg.Server.Port),
		Handler:        api.NewRouter(app),
		ReadTimeout:    time.Duration(cfg.Server.Timeout) * time.Second,
		WriteTimeout:   time.Duration(cfg.Server.Timeout) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := s.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}
