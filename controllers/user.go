package controllers

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/rubichandrap/hello-go/database"
	"github.com/rubichandrap/hello-go/dtos"
	"github.com/rubichandrap/hello-go/repositories"
	"github.com/rubichandrap/hello-go/utils"
)

type UserController struct{}

func (m *UserController) Get(c *fiber.Ctx) error {
	db, _ := database.Connect()
	res := dtos.Response{}

	r := repositories.NewUserRepository(db)

	users, dbResult := r.Get()

	if dbResult.Error != nil {
		c.Status(500).JSON(res.InternalServerError(dtos.ResponsePayload{
			Message: dbResult.Error.Error(),
		}))
	}

	for i := 0; i < len(users); i++ {
		users[i].Password = nil
	}

	return c.JSON(res.OK(dtos.ResponsePayload{
		Data: users,
	}))
}

func (m *UserController) GetByID(c *fiber.Ctx) error {
	db, _ := database.Connect()
	res := dtos.Response{}

	userID := c.Params("userID")

	id, err := strconv.Atoi(userID)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(res.BadRequest(dtos.ResponsePayload{
			Message: fmt.Sprintf("ID is not valid: %v", err),
		}))
	}

	var user dtos.User

	r := repositories.NewUserRepository(db)

	dbResult := r.GetByID(&user, &id)

	if dbResult.Error != nil {
		return c.Status(dbResult.StatusCode).JSON(res.Send(dbResult.StatusCode, dtos.ResponsePayload{
			Message: dbResult.Error.Error(),
		}))
	}

	user.Password = nil

	return c.JSON(res.OK(dtos.ResponsePayload{
		Data: user,
	}))
}

func (m *UserController) Store(c *fiber.Ctx) error {
	db, _ := database.Connect()
	res := dtos.Response{}
	validationRes := dtos.ValidationResponse{}

	user := new(dtos.User)

	err := c.BodyParser(user)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(res.BadRequest(dtos.ResponsePayload{
			Message: fmt.Sprintf("failed when parsing the body: %v", err),
		}))
	}

	validationErrors := utils.ValidateStruct(*user)

	if validationErrors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(validationRes.Send(validationErrors))
	}

	r := repositories.NewUserRepository(db)

	isEmailExists, dbResult := r.IsFieldExists("email", user.Email)

	if dbResult.Error != nil {
		return c.Status(dbResult.StatusCode).JSON(res.Send(dbResult.StatusCode, dtos.ResponsePayload{
			Message: dbResult.Error.Error(),
		}))
	}

	if isEmailExists {
		return c.Status(fiber.StatusBadRequest).JSON(res.BadRequest(dtos.ResponsePayload{
			Message: "Email is already exists",
		}))
	}

	hashedPassword, err := utils.HashPassword(*user.Password)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(res.BadRequest(dtos.ResponsePayload{
			Message: fmt.Sprintf("error when hashing password: %v", err),
		}))
	}

	user.Password = &hashedPassword

	dbResult = r.Store(user)

	if dbResult.Error != nil {
		return c.Status(dbResult.StatusCode).JSON(res.Send(dbResult.StatusCode, dtos.ResponsePayload{
			Message: dbResult.Error.Error(),
		}))
	}

	user.Password = nil

	return c.Status(fiber.StatusCreated).JSON(res.Created(dtos.ResponsePayload{
		Data: user,
	}))
}

func (m *UserController) Update(c *fiber.Ctx) error {
	db, _ := database.Connect()
	res := dtos.Response{}
	validationRes := dtos.ValidationResponse{}

	userID := c.Params("userID")

	id, err := strconv.Atoi(userID)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(res.BadRequest(dtos.ResponsePayload{
			Message: fmt.Sprintf("ID is not valid: %v", err),
		}))
	}

	r := repositories.NewUserRepository(db)

	isExists, dbResult := r.IsFieldExists("id", id)

	if dbResult.Error != nil {
		return c.Status(dbResult.StatusCode).JSON(res.Send(dbResult.StatusCode, dtos.ResponsePayload{
			Message: dbResult.Error.Error(),
		}))
	}

	if !isExists {
		return c.Status(fiber.StatusNotFound).JSON(res.NotFound(dtos.ResponsePayload{
			Message: "Data is not found",
		}))
	}

	var user dtos.User

	var updateUserData dtos.UpdateUser

	err = c.BodyParser(&updateUserData)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(res.BadRequest(dtos.ResponsePayload{
			Message: fmt.Sprintf("failed when parsing the body: %v", err),
		}))
	}

	validationErrors := utils.ValidateStruct(updateUserData)

	if validationErrors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(validationRes.Send(validationErrors))
	}

	dbResult = r.Update(&user, &id, &updateUserData)

	if dbResult.Error != nil {
		return c.Status(dbResult.StatusCode).JSON(res.Send(dbResult.StatusCode, dtos.ResponsePayload{
			Message: dbResult.Error.Error(),
		}))
	}

	user.Password = nil

	return c.JSON(res.OK(dtos.ResponsePayload{
		Data: user,
	}))
}

func (m *UserController) UpdatePassword(c *fiber.Ctx) error {
	db, _ := database.Connect()
	res := dtos.Response{}
	validationRes := dtos.ValidationResponse{}

	userID := c.Params("userID")

	id, err := strconv.Atoi(userID)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(res.BadRequest(dtos.ResponsePayload{
			Message: fmt.Sprintf("ID is not valid: %v", err),
		}))
	}

	r := repositories.NewUserRepository(db)

	isExists, dbResult := r.IsFieldExists("id", id)

	if dbResult.Error != nil {
		return c.Status(dbResult.StatusCode).JSON(res.Send(dbResult.StatusCode, dtos.ResponsePayload{
			Message: dbResult.Error.Error(),
		}))
	}

	if !isExists {
		return c.Status(fiber.StatusNotFound).JSON(res.NotFound(dtos.ResponsePayload{
			Message: "Data is not found",
		}))
	}

	var updateUserPassword dtos.UpdateUserPassword

	err = c.BodyParser(&updateUserPassword)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(res.BadRequest(dtos.ResponsePayload{
			Message: fmt.Sprintf("failed when parsing the body: %v", err),
		}))
	}

	validationErrors := utils.ValidateStruct(updateUserPassword)

	if validationErrors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(validationRes.Send(validationErrors))
	}

	dbResult = r.UpdatePassword(&id, &updateUserPassword)

	if dbResult.Error != nil {
		return c.Status(dbResult.StatusCode).JSON(res.Send(dbResult.StatusCode, dtos.ResponsePayload{
			Message: dbResult.Error.Error(),
		}))
	}

	return c.JSON(res.OK(dtos.ResponsePayload{}))
}
