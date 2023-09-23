package server

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type server struct {
	app *fiber.App
	db  *sql.DB
}

func New(db *sql.DB) *server {
	return &server{
		app: fiber.New(),
		db:  db,
	}
}

func (s *server) Run() error {
	s.app.Use(recover.New())
	s.app.Use(logger.New())

	s.MapRoutes()

	port := "8080"
	if p := os.Getenv("SERVER_PORT"); p != "" {
		port = p
	}

	log.Printf("server is running on port %s\n", port)
	return s.app.Listen(fmt.Sprintf(":%s", port))
}
