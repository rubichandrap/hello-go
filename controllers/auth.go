package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/rubichandrap/hello-go/database"
	"github.com/rubichandrap/hello-go/dtos"
	"github.com/rubichandrap/hello-go/repositories"
	"github.com/rubichandrap/hello-go/utils"
)

type AuthController struct{}

func (m *AuthController) Register(c *fiber.Ctx) error {
	l := UserController{}
	return l.Store(c)
}

func (m *AuthController) Login(c *fiber.Ctx) error {
	db, _ := database.Connect()
	res := dtos.Response{}
	validationRes := dtos.ValidationResponse{}

	var loginPayload dtos.Login

	err := c.BodyParser(&loginPayload)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(res.BadRequest(dtos.ResponsePayload{
			Message: fmt.Sprintf("failed when parsing the body: %v", err),
		}))
	}

	validationErrors := utils.ValidateStruct(loginPayload)

	if validationErrors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(validationRes.Send(validationErrors))
	}

	r := repositories.NewAuthRepository(db)

	dbResult := r.Login(&loginPayload)

	if dbResult.Error != nil {
		return c.Status(dbResult.StatusCode).JSON(res.Send(dbResult.StatusCode, dtos.ResponsePayload{
			Message: dbResult.Error.Error(),
		}))
	}

	return c.JSON(res.OK(dtos.ResponsePayload{
		Data: dbResult.Data,
	}))
}
