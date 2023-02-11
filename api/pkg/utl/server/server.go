package server

import (
	"github.com/docker/docker/api/server/middleware"
	"github.com/gofiber/fiber/middleware"
	"github.com/gofiber/fiber/v2"
)

func New() *fiber.App {

	f := fiber.New()
	f.Use(middleware.Logger(), middleware.Recover())
}
