package main

import (    
    e "engine/server"
    "myforum/server/database"

)

func main() {
    database.InitDB()
    var forum e.User
	e.Run(&forum)
}   