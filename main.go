package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/rajajamal/resto-backend/configs"
	"github.com/rajajamal/resto-backend/models"
	"github.com/rajajamal/resto-backend/routes"
)

func init() {
	godotenv.Load()
	configs.Load()
	configs.Db.AutoMigrate(
		&models.Customer{},
		// &models.Menu{},
		// &models.Order{},
		// &models.OrderDetail{},
	)
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("مرحبا بالعالم")
	})

	(routes.Customer{}).RegisterRoutes(app)

	app.Listen(fmt.Sprintf(":%d", configs.Env.AppPort))
}
