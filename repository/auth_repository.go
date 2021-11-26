package repository

import (
	"fortune-cookies/config"
	"fortune-cookies/entity"
	"fortune-cookies/helper"
	"net/http"
	"strings"

	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Registration(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var user entity.User
		user.Email = c.FormValue("email")
		user.Password = c.FormValue("password")

		//Check user exist
		result := db.Where("email = ?", user.Email).Find(&user)
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected > 0 {
			return c.JSON(http.StatusOK, helper.ResultResponse(true, "Email already used", map[string]interface{}{}))
		}

		//Hashing password
		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 5)
		if err != nil {
			return c.JSON(http.StatusOK, helper.ResultResponse(true, "Error while hashing password", map[string]interface{}{}))
		}

		//Token
		cfg, _ := config.NewConfig(".env")
		user.Password = string(hash)
		user.Token = helper.JwtGenerator(user.Email, cfg.JWTConfig.SecretKey)

		//Post Registration
		regisResult := db.Create(&user)
		if regisResult.Error != nil {
			return c.JSON(http.StatusOK, helper.ResultResponse(true, "Error when registration", map[string]interface{}{}))
		}
		var cart entity.Cart
		cart.UserID = user.Id
		regisCart := db.Create(&cart)
		if regisCart.Error != nil {
			return c.JSON(http.StatusOK, helper.ResultResponse(true, "Error when registration", map[string]interface{}{}))
		}

		user.Password = "hidden"
		return c.JSON(http.StatusOK, helper.ResultResponse(false, "Registration Success", &user))
	}
}

func Login(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var userInput entity.User
		var userResult entity.User
		userInput.Email = c.FormValue("email")
		userInput.Password = c.FormValue("password")

		//Login
		resLogin := db.Where("email = ?", userInput.Email).Find(&userResult)
		if resLogin.Error != nil {
			return resLogin.Error
		}

		//Check user exist
		if resLogin.RowsAffected == 0 {
			return c.JSON(http.StatusOK, helper.ResultResponse(true, "Invalid Email", map[string]interface{}{}))
		}

		//Check password
		checkPass := bcrypt.CompareHashAndPassword([]byte(userResult.Password), []byte(userInput.Password))
		if checkPass != nil {
			return c.JSON(http.StatusOK, helper.ResultResponse(true, "Invalid Password", map[string]interface{}{}))
		}

		//Token
		cfg, _ := config.NewConfig(".env")
		userResult.Token = helper.JwtGenerator(userResult.Email, cfg.JWTConfig.SecretKey)

		//Login success result
		userResult.Password = "hidden"
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
		
		user.Password = "hidden"
		return c.JSON(http.StatusOK, helper.ResultResponse(false, "Get User By Id Success", &user))
	}
}

func GetUserByToken(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var user entity.User

		//Get token
		headerToken := c.Request().Header.Get("Authorization")
		headerToken = strings.ReplaceAll(headerToken, "Bearer", "")
		headerToken = strings.ReplaceAll(headerToken, " ", "")
		
		//Query
		result := db.First(&user, "token = ?", headerToken)
		if result.Error != nil {
			return c.JSON(http.StatusOK, helper.ResultResponse(true, "Error Occured While Querying SQL", result.Error))
		}
		if result.RowsAffected == 0 {
			return c.JSON(http.StatusOK, helper.ResultResponse(true, "Token Not Found", ""))
		}

		user.Password = "hidden"
		return c.JSON(http.StatusOK, helper.ResultResponse(false, "Fetch Data Success", &user))
	}
}
