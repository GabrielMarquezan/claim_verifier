package main

import (
	"api/config"
	"api/db"
	"api/router"
	"context"
)

func main() {
	cfg := config.Load()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	database := db.ConnectDatabase(cfg)
	r := router.Init(&ctx, database, cfg)

	r.Run(":3000")
}
