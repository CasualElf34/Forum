package engine

import (
    "log"
    "net/http"
)

func Run(forum *User) {
    http.HandleFunc("/register", RegisterHandler)
    http.HandleFunc("/login", LoginHandler)

    log.Println("Serveur lancé sur http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}