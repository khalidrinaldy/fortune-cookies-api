package repository

import (
	"fortune-cookies/entity"
	"fortune-cookies/helper"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func GetAllProducts(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var products []entity.Product
		result := db.Find(&products)
		if result.Error != nil {
			return c.JSON(http.StatusOK, helper.ResultResponse(true, "Get All Products Failed", &products))
		}
		return c.JSON(http.StatusOK, helper.ResultResponse(false, "Get All Products Success", &products))
	}
}

func GetProductsByCategory(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var products []entity.Product
		result := db.Order("id asc").Where("product_category = ?", c.Param("category")).Find(&products)
		if result.Error != nil {
			return c.JSON(http.StatusOK, helper.ResultResponse(true, "Get Product By Category Failed", &products))
		}
		return c.JSON(http.StatusOK, helper.ResultResponse(false, "Get Product By Category Success", &products))
	}
}

func GetOneProduct(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var product entity.Product
		result := db.Where("id = ?", c.Param("id")).Find(&product)
		if result.Error != nil {
			return result.Error
		}
		return c.JSON(http.StatusOK, helper.ResultResponse(false, "Get One Product By Id Success", &product))
	}
}

func AddProduct(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var product entity.Product
		var admin entity.Admin

		//Get token
		headerToken := c.Request().Header.Get("Authorization")
		headerToken = strings.ReplaceAll(headerToken, "Bearer", "")
		headerToken = strings.ReplaceAll(headerToken, " ", "")

		//Check Is Admin
		resultAdmin := db.First(&admin, "token = ?", headerToken)
		if resultAdmin.Error != nil {
			return c.JSON(http.StatusOK, helper.ResultResponse(true, "Error Occured While Querying SQL", resultAdmin.Error.Error()))
		}
		if resultAdmin.RowsAffected == 0 {
			return c.JSON(http.StatusOK, helper.ResultResponse(true, "Admin Token Not Found", ""))
		}

		product.Product_Name = c.FormValue("product_name")
		product.Product_Category = c.FormValue("product_category")
		product.Product_Price, _ = strconv.Atoi(c.FormValue("product_price"))
		product.Product_Image = c.FormValue("product_image")
		product.Product_Description = c.FormValue("product_description")

		result := db.Create(&product)
		if result.Error != nil {
			return c.JSON(http.StatusOK, helper.ResultResponse(true, "Error Occured While Querying SQL", result.Error.Error()))
		}
		return c.JSON(http.StatusOK, &product)
	}
}

func UpdateProduct(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var product entity.Product
		var admin entity.Admin

		//Get token
		headerToken := c.Request().Header.Get("Authorization")
		headerToken = strings.ReplaceAll(headerToken, "Bearer", "")
		headerToken = strings.ReplaceAll(headerToken, " ", "")

		//Check Is Admin
		resultAdmin := db.First(&admin, "token = ?", headerToken)
		if resultAdmin.Error != nil {
			return c.JSON(http.StatusOK, helper.ResultResponse(true, "Error Occured While Querying SQL", resultAdmin.Error.Error()))
		}
		if resultAdmin.RowsAffected == 0 {
			return c.JSON(http.StatusOK, helper.ResultResponse(true, "Admin Token Not Found", ""))
		}

		product.ID, _ = strconv.Atoi(c.Param("product_id"))
		product.Product_Name = c.FormValue("product_name")
		product.Product_Category = c.FormValue("product_category")
		product.Product_Price, _ = strconv.Atoi(c.FormValue("product_price"))
		product.Product_Image = c.FormValue("product_image")
		product.Product_Description = c.FormValue("product_description")

		result := db.Model(&product).Where("id = ?", c.Param("product_id")).Updates(map[string]interface{}{
			"product_name":        product.Product_Name,
			"product_category":    product.Product_Category,
			"product_price":       product.Product_Price,
			"product_image":       product.Product_Image,
			"product_description": product.Product_Description,
		})
		// query := fmt.Sprintf(`update products set product_name = '%s', product_category = '%s', product_price = %d, product_image = '%s', product_description = '%s' where id = %d;`,
		// 	product.Product_Name,
		// 	product.Product_Category,
		// 	product.Product_Price,
		// 	product.Product_Image,
		// 	product.Product_Description,
		// 	product.ID,
		// )
		// result := db.Exec(query)
		if result.Error != nil {
			return c.JSON(http.StatusOK, helper.ResultResponse(true, c.Param("product_id"), result.Error.Error()))
		}
		return c.JSON(http.StatusOK, &product)
	}
}

func DeleteProduct(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var product entity.Product
		var admin entity.Admin

		//Get token
		headerToken := c.Request().Header.Get("Authorization")
		headerToken = strings.ReplaceAll(headerToken, "Bearer", "")
		headerToken = strings.ReplaceAll(headerToken, " ", "")

		//Check Is Admin
		resultAdmin := db.First(&admin, "token = ?", headerToken)
		if resultAdmin.Error != nil {
			return c.JSON(http.StatusOK, helper.ResultResponse(true, "Error Occured While Querying SQL", resultAdmin.Error.Error()))
		}
		if resultAdmin.RowsAffected == 0 {
			return c.JSON(http.StatusOK, helper.ResultResponse(true, "Admin Token Not Found", ""))
		}

		result := db.Delete(&product, c.Param("id"))
		if result.Error != nil {
			return c.JSON(http.StatusOK, helper.ResultResponse(true, "Error Occured While Querying SQL", result.Error.Error()))
		}
		return c.JSON(http.StatusOK, product)
	}
}
