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
	app := fiber.New(
		fiber.Config{
			Views: engine,
		},
	)
	app.Static("/static", "./static")
	server := &FiberServer{
		App: app,
		db:  database.New(),
	}

	return server
}
