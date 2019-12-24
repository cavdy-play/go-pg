package main

import (
	"log"
	db "github.com/cavdy-play/go_pg/db"
)

func main() {
	log.Printf("Hello World...!!!\n")
	db.Connect()
}
