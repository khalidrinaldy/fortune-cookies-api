package repository

import (
	"fortune-cookies/entity"
	"fortune-cookies/helper"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func GetAdminByToken(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var admin entity.Admin

		//Get token
		headerToken := c.Request().Header.Get("Authorization")
		headerToken = strings.ReplaceAll(headerToken, "Bearer", "")
		headerToken = strings.ReplaceAll(headerToken, " ", "")

		//Query
		result := db.First(&admin, "token = ?", headerToken)
		if result.Error != nil {
			return c.JSON(http.StatusOK, helper.ResultResponse(true, "Error Occured While Querying SQL", result.Error))
		}
		if result.RowsAffected == 0 {
			return c.JSON(http.StatusOK, helper.ResultResponse(true, "Token Not Found", ""))
		}

		admin.Password = "hidden"
		return c.JSON(http.StatusOK, helper.ResultResponse(false, "Fetch Data Success", &admin))
	}
}

func RegisterAdmin(db *gorm.DB) echo.HandlerFunc{
	return func(c echo.Context) error {
		var admin entity.Admin

		//Get Value
		admin.Email = c.FormValue("email")
		admin.Password = c.FormValue("password")

		//Check Admin exist
		result := db.Where("email = ?", admin.Email).Find(&admin)
		if result.Error != nil {
			return c.JSON(http.StatusOK, helper.ResultResponse(true, "Error Querying SQL", result.Error))
		}
		if result.RowsAffected > 0 {
			return c.JSON(http.StatusOK, helper.ResultResponse(true, "Email already used", map[string]interface{}{}))
		}

		//Hashing password
		hash, err := bcrypt.GenerateFromPassword([]byte(admin.Password), 5)
		if err != nil {
			return c.JSON(http.StatusOK, helper.ResultResponse(true, "Error while hashing password", map[string]interface{}{}))
		}
		admin.Password = string(hash)

		//Token
		admin.Token = helper.JwtGenerator(admin.Email, os.Getenv("SECRET_KEY"))

		//Post Registration
		regisResult := db.Create(&admin)
		if regisResult.Error != nil {
			return c.JSON(http.StatusOK, helper.ResultResponse(true, "Error when registration", map[string]interface{}{}))
		}
		admin.Password = "hidden"
		return c.JSON(http.StatusOK, helper.ResultResponse(false, "Registration Success", &admin))
	}
}

func LoginAdmin(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var adminInput entity.Admin
		var adminResult entity.Admin

		//Get Value
		adminInput.Email = c.FormValue("email")
		adminInput.Password = c.FormValue("password")

		//Check email exist
		result := db.Where("email = ?", adminInput.Email).Find(&adminResult)
		if result.Error != nil {
			return c.JSON(http.StatusOK, helper.ResultResponse(true, "Error Querying SQL", result.Error))
		}
		if result.RowsAffected == 0 {
			return c.JSON(http.StatusOK, helper.ResultResponse(true, "Invalid Email", map[string]interface{}{}))
		}

		//Check Password
		checkPass := bcrypt.CompareHashAndPassword([]byte(adminResult.Password), []byte(adminInput.Password))
		if checkPass != nil {
			return c.JSON(http.StatusOK, helper.ResultResponse(true, "Invalid Password", map[string]interface{}{}))
		}

		adminResult.Password = "hidden"
		return c.JSON(http.StatusOK, helper.ResultResponse(false, "Registration Success", &adminResult))
	}
}