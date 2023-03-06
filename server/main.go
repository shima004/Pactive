package main

import (
	"fmt"
	"log"

	"github.com/shima004/pactive/db"
	"github.com/shima004/pactive/server"
)

func main() {
	connection, err := db.InitDB()
	if err != nil {
		panic(err)
	}
	db, err := connection.DB()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	log.Println("db connected")

	server.Run()

	fmt.Println(db)
}
