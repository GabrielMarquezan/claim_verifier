package main

import (
	"api/db"
	"api/router"
	"context"
	"os"
)

func main() {
	databaseURL, isPresent := os.LookupEnv("DATABASE_URL")
	if !isPresent {
		panic("DB URL not defined!")
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	database := db.ConnectDatabase(databaseURL)
	r := router.Init(&ctx, database)

	r.Run(":3000")
}
