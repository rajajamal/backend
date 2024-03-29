package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/iancoleman/strcase"
	"github.com/rajajamal/resto-backend/models"
	"github.com/rajajamal/resto-backend/services"
)

type Customer struct {
	Service services.Customer
}

// Get Customer by ID

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

//Register customer
func (c Customer) Register(ctx *fiber.Ctx) error {
	model := models.Customer{}
	ctx.BodyParser(&model)
	model.ID = 0

	validate := validator.New()
	err := validate.Struct(model)
	if err != nil {
		messages := []string{}
		for _, err := range err.(validator.ValidationErrors) {
			messages = append(messages, fmt.Sprintf("%s is %s", strcase.ToSnake(err.Field()), err.Tag))
		}

		ctx.JSON(map[string]interface{}{
			"message": messages,
		})

		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	model, err = c.Service.Save(model)
	if err != nil {
		ctx.JSON(map[string]string{
			"message": err.Error(),
		})

		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	ctx.JSON(model)

	return ctx.SendStatus(fiber.StatusCreated)
}
