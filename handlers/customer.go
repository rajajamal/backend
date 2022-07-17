package handlers

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/rajajamal/resto-backend/services"
)

type Customer struct {
	Service services.Customer
}

func (c Customer) Get(ctx *fiber.Ctx) error {
	customerId, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		ctx.JSON(map[string]string{
			"message": "id is not number",
		})

		return ctx.SendStatus(http.StatusBadRequest)
	}

	customer, err := c.Service.Get(customerId)
	if err != nil {
		ctx.JSON(map[string]string{
			"message": "customer not found",
		})

		return ctx.SendStatus(fiber.StatusNotFound)
	}

	return ctx.JSON(customer)
}
