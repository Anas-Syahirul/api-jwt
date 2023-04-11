package main

import (
	"api-jwt/database"
	"api-jwt/routers"
	"log"
)

func main() {
	database.StartDB()
	r := routers.StartApp()
	log.Println("starting app...")
	r.Run(":8080")
}
