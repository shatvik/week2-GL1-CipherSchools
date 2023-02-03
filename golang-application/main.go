package main

import (
	"log"

	"github.com/shatvik/week2-GL1-CipherSchools.git/database"
	"github.com/shatvik/week2-GL1-CipherSchools.git/routers"
)

func main() {
	database.Setup()
	engine := routers.Router()
	err := engine.Run("127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
}
