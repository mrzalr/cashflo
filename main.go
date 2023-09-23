package main

import (
	"log"

	"github.com/mrzalr/cashflo/internal/server"
	"github.com/mrzalr/cashflo/pkg/db/mysql"
)

func main() {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	db, err := mysql.New()
	if err != nil {
		log.Fatal(err)
	}

	s := server.New(db)

	err = s.Run()
	if err != nil {
		log.Fatal(err)
	}
}
