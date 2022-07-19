package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rajajamal/resto-backend/configs"
	"github.com/rajajamal/resto-backend/handlers"
	"github.com/rajajamal/resto-backend/repositories"
	"github.com/rajajamal/resto-backend/services"
)

type Menu struct {
}

func (Menu) RegisterRoutes(router fiber.Router) {
	menuRepository := repositories.Menu{Storage: configs.Db}
	menuService := services.Menu{Repository: menuRepository}

	menu := handlers.Menu{Service: menuService}

	router.Get("/menus", menu.GetAll)
	router.Get("/menus/:id", menu.Get)
}
