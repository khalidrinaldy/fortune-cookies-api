package repository

import (
	"fmt"
	"fortune-cookies/config"
	"fortune-cookies/entity"
	"fortune-cookies/helper"
	"net/http"

	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Registration(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var user entity.User
		user.Username = c.FormValue("username")
		user.Password = c.FormValue("password")
		fmt.Println(user.Username)

		//Check user exist
		result := db.Where("username = ?", user.Username).Find(&user)
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected > 0 {
			return c.JSON(http.StatusOK, helper.ResultResponse(true, "username already used", map[string]interface{}{}))
		}

		//Hashing password
		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 5)
		if err != nil {
			return c.JSON(http.StatusOK, helper.ResultResponse(true, "Error while hashing password", map[string]interface{}{}))
		}

		//Token
		cfg, _ := config.NewConfig(".env")
		user.Username = c.FormValue("username")
		user.Password = string(hash)
		user.Token = helper.JwtGenerator(user.Username, cfg.JWTConfig.SecretKey)

		//Post Registration
		regisResult := db.Create(&user)
		if regisResult.Error != nil {
			return c.JSON(http.StatusOK, helper.ResultResponse(true, "Error when registration", map[string]interface{}{}))
		}
		var cart entity.Cart
		cart.UserID = user.ID
		regisCart := db.Create(&cart)
		if regisCart.Error != nil {
			return c.JSON(http.StatusOK, helper.ResultResponse(true, "Error when registration", map[string]interface{}{}))
		}

		return c.JSON(http.StatusOK, helper.ResultResponse(false, "Registration Success", &user))
	}
}

func Login(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var userInput entity.User
		var userResult entity.User
		userInput.Username = c.FormValue("username")
		userInput.Password = c.FormValue("password")

		//Login
		resLogin := db.Where("username = ?", userInput.Username).Find(&userResult)
		if resLogin.Error != nil {
			return resLogin.Error
		}

		//Check user exist
		if resLogin.RowsAffected == 0 {
			return c.JSON(http.StatusOK, helper.ResultResponse(true, "Invalid Username", map[string]interface{}{}))
		}

		//Check password
		checkPass := bcrypt.CompareHashAndPassword([]byte(userResult.Password), []byte(userInput.Password))
		if checkPass != nil {
			return c.JSON(http.StatusOK, helper.ResultResponse(true, "Invalid Password", map[string]interface{}{}))
		}

		//Token
		cfg, _ := config.NewConfig(".env")
		userResult.Token = helper.JwtGenerator(userResult.Username, cfg.JWTConfig.SecretKey)

		//Login success result
		return c.JSON(http.StatusOK, helper.ResultResponse(false, "Login Success", &userResult))
	}
}

func GetUser(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var user entity.User
		result := db.Where("id = ?", c.Param("id")).Find(&user)
		if result.Error != nil {
			return result.Error
		}
		return c.JSON(http.StatusOK, helper.ResultResponse(false, "Get User By Id Success", &user))
	}
}
