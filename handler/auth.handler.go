package handler

import (
	"belajar_golang_fiber/database"
	"belajar_golang_fiber/model/entity"
	"belajar_golang_fiber/model/request"
	"belajar_golang_fiber/utils"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func LoginHandler(ctx *fiber.Ctx) error {
	loginRequest := new(request.LoginRequest)

	if err := ctx.BodyParser(loginRequest); err != nil {
		return err
	}

	//VALIDASI REQUEST
	validate := validator.New()
	errValidate := validate.Struct(loginRequest)

	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	var user entity.User
	err := database.DB.First(&user, "email = ?", loginRequest.Email).Error
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "wrong credential",
		})
	}

	isValid := utils.CheckPasswordHash(loginRequest.Password, user.Password)
	
	if !isValid {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "wrong credential",
		})
	}

	// GENERATE JWT
	claims := jwt.MapClaims{}
	claims["name"] = user.Name
	claims["email"] = user.Email
	claims["address"] = user.Address
	claims["exp"] = time.Now().Add(time.Minute * 2).Unix()

	if user.Email == "example@gmail.com" {
		claims["role"] = "admin"
	}else {
		claims["role"] = "user"
	}

	token, errGenerateToken := utils.GenerateToken(&claims)

	if errGenerateToken != nil {
		log.Println(errGenerateToken)
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "wrong credential",
		})
	}

	return ctx.JSON(fiber.Map{
		"token": token,
	})
}
