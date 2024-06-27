package handler

import (
	"belajar_golang_fiber/database"
	"belajar_golang_fiber/model/entity"
	"belajar_golang_fiber/model/request"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func UserHandlerGetAll(ctx *fiber.Ctx) error {

	var users []entity.User
	result := database.DB.Find(&users)

	if result.Error != nil {
		log.Println(result.Error)
	}

	return ctx.JSON(users)
}

func UserHandlerCreate(ctx *fiber.Ctx) error {
	user := new(request.UserCreateRequest)

	if err := ctx.BodyParser(user); err != nil {
		return err
	}
	validate := validator.New()
	errValidate := validate.Struct(user)

	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	newUser := entity.User{
		Name:    user.Name,
		Email:   user.Email,
		Address: user.Address,
		Phone:   user.Phone,
	}

	errCreateUser := database.DB.Create(&newUser).Error

	if errCreateUser != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Failed to store data",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Success",
		"data":    newUser,
	})
}

func UserHandlerGetByID(ctx *fiber.Ctx) error {
	userId := ctx.Params("id")

	var user entity.User
	err := database.DB.First(&user, "id = ?", userId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	// userResponse := response.UserResponse{
	// 	ID: user.ID,
	// 	Name: user.Name,
	// 	Email: user.Email,
	// 	Address: user.Address,
	// 	Phone: user.Phone,
	// 	CreatedAt: user.CreatedAt,
	// 	UpdatedAt: user.UpdatedAt,
	// }
	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    user,
	})

}
