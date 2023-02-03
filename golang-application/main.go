package main

import (
	"log"

	"github.com/shatvik/week2-GL1-CipherSchools.git/database"
	"github.com/shatvik/week2-GL1-CipherSchools.git/routers"
)

// init func always called before main:

func init() {
	database.Setup()
}

// NOte: = We can have multiple init functions ->
// All the init funcs will execute one by one ina sequence:

// 1st
// func init() {
// 	fmt.Println(1)
// }
// 2nd
// func init() {
// 	fmt.Println(2)
// }
// 3rd
// func init() {
// 	fmt.Println(3)
// }

func main() {

	engine := routers.Router()
	err := engine.Run("127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
}
