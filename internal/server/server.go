package server

import (
	"htmx-tutorial/internal/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/django/v3"
)

type FiberServer struct {
	*fiber.App
	db database.Service
}

func New() *FiberServer {
	engine := django.New("./internal/templates", ".html")
	server := &FiberServer{
		App: fiber.New(
			fiber.Config{
				Views: engine,
			},
		),
		db: database.New(),
	}

	return server
}
