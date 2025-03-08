package main

import (
    e "engine/server"
    "net/http"
)

func main() {
    e.InitDB()
    css := http.FileServer(http.Dir("css"))
	http.Handle("/css/", http.StripPrefix("/css/", css))
    var forum e.User
	e.Run(&forum)
}   