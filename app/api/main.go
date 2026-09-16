package main

import (
	"api/db"
	"api/router"
	"fmt"
	"os"
)

func main() {
	databaseURL, isPresent := os.LookupEnv("DATABASE_URL")
	if !isPresent {
		panic("DB URL not defined!")
	}

	DB := db.ConnectDatabase(databaseURL)
	r := router.Init()

	r.Run(":3000")
	fmt.Println(DB)
}
