package main

import (
	"log"

	"github.com/luytbq/go-angular-reactive-crud-example/api"
	"github.com/luytbq/go-angular-reactive-crud-example/config"
	"github.com/luytbq/go-angular-reactive-crud-example/database"
)

func main() {
	db, err := database.NewPostgresDB()

	if err != nil {
		log.Fatal(err)
	}

	server := api.NewAPIServer(
		config.App.SERVER_PORT,
		config.App.SERVER_API_PREFIX,
		db,
	)

	err = server.Run()

	if err != nil {
		log.Fatal(err)
	}

	log.Println(db)
}
