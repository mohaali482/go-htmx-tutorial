package server

import (
	"github.com/gofiber/fiber/v2"
)

func (s *FiberServer) RegisterFiberRoutes() {
	s.App.Get("/", s.index)
	s.App.Get("/health", s.healthHandler)
}

func (s *FiberServer) index(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{})
}

func (s *FiberServer) healthHandler(c *fiber.Ctx) error {
	return c.JSON(s.db.Health())
}
