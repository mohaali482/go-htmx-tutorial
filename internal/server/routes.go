package server

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (s *FiberServer) RegisterFiberRoutes() {
	s.App.Get("/", s.index)
	s.App.Post("/create", s.createCount)
	s.App.Delete("/delete/:id", s.deleteCount)
	s.App.Post("/increment/:id", s.incrementCount)
	s.App.Post("/decrement/:id", s.decrementCount)
}

func (s *FiberServer) index(c *fiber.Ctx) error {
	counts, _ := s.db.GetAllCounts()
	return c.Render("index", fiber.Map{"counts": counts})
}

func (s *FiberServer) createCount(c *fiber.Ctx) error {
	count := s.db.CreateCount()
	return c.Render("count", fiber.Map{"count": count})
}

func (s *FiberServer) incrementCount(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	s.db.IncrementCount(id)
	count, _ := s.db.GetCount(id)
	return c.Render("count", fiber.Map{"count": count})
}

func (s *FiberServer) decrementCount(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	err = s.db.DecrementCount(id)
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	count, err := s.db.GetCount(id)
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	return c.Render("count", fiber.Map{"count": count})
}

func (s *FiberServer) deleteCount(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	err = s.db.DeleteCount(id)
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	return c.SendStatus(fiber.StatusOK)
}
