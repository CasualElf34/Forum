package main

import (
	"log"
	"net/http"

	"myforum/database"
	"myforum/handlers"
)

func main() {
	database.InitDB()

	http.HandleFunc("/register", handlers.RegisterHandler)
	http.HandleFunc("/login", handlers.LoginHandler)

	log.Println("Serveur lancé sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}