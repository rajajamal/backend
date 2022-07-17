package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rajajamal/resto-backend/configs"
	"github.com/rajajamal/resto-backend/handlers"
	"github.com/rajajamal/resto-backend/repositories"
	"github.com/rajajamal/resto-backend/services"
)

type Customer struct {
}

func (Customer) RegisterRoutes(router fiber.Router) {
	customerRepository := repositories.Customer{Storage: configs.Db}
	// orderRepository := repositories.Order{Storage: configs.Db}

	customerService := services.Customer{
		Repository: customerRepository,
		// Order:      &orderRepository,

	}

	customer := handlers.Customer{Service: customerService}

	// router.Post("/customers", customer.Register)
	router.Get("/customers/:id", customer.Get)
	// router.Post("/customers/:id/reservation", customer.Reservation)
}
