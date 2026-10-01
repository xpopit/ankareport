package main

import (
	"log"

	"backend/db"
	"backend/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	db.Connect()

	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())

	v1 := app.Group("/api/v1")
	reports := v1.Group("/reports")

	reports.Get("/", handlers.GetReports)
	reports.Post("/", handlers.CreateReport)
	reports.Get("/:id", handlers.GetReport)
	reports.Put("/:id", handlers.UpdateReport)
	reports.Delete("/:id", handlers.DeleteReport)
	reports.Post("/:id/clone", handlers.CloneReport)

	app.Static("/", "../frontend/dist")
	app.Get("*", func(c *fiber.Ctx) error {
		return c.SendFile("../frontend/dist/index.html")
	})

	log.Fatal(app.Listen(":3000"))
}
