package main

import (
	"fmt"
	"github.com/ismailbayram/shopping/config"
)

func main() {
	cfg := config.Init()
	fmt.Println(cfg.Database.Port)
	fmt.Println("Hello from Server!")
}
