package repository

import (
	"fmt"
	"fortune-cookies/entity"
	"fortune-cookies/helper"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func GetCartList(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var cart []entity.CartProductsList
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

		//Query
		resultCart := db.Raw("select cart_products.id, cart_products.product_id, products.product_name, products.product_price, products.product_image, cart_products.amount from cart_products join products on products.id = cart_products.product_id join carts on carts.id = cart_products.cart_id where carts.user_id = ?;", user.Id).
				Scan(&cart)
		if resultCart.Error != nil {
			return c.JSON(http.StatusOK, helper.ResultResponse(true, "Get Cart List Failed", resultCart.Error))
		}
		return c.JSON(http.StatusOK, helper.ResultResponse(false, "Get Cart List Success", &cart))
	}
}

func AddToCart(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var addCart entity.Cart_Products
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

		addCart.CartID = user.Id
		addCart.ProductID, _ = strconv.Atoi(c.FormValue("product_id"))
		addCart.Amount, _ = strconv.Atoi(c.FormValue("amount"))

		//Query
		resultAdd := db.Create(&addCart)
		if resultAdd.Error != nil {
			return c.JSON(http.StatusOK, helper.ResultResponse(true, "Add To Cart Failed", resultAdd.Error))
		}
		return c.JSON(http.StatusOK, helper.ResultResponse(false, "Add To Cart Success", &addCart))
	}
}

func UpdateItemCart(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var updateItem entity.Cart_Products
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

		updateItem.CartID = user.Id
		updateItem.ProductID, _ = strconv.Atoi(c.FormValue("product_id"))
		updateItem.Amount, _ = strconv.Atoi(c.FormValue("amount"))

		//Query
		resultEdit := db.Model(&updateItem).Where("cart_id = ?", updateItem.CartID).Updates(entity.Cart_Products{
			CartID: updateItem.CartID,
			ProductID: updateItem.ProductID,
			Amount: updateItem.Amount,
		})
		if resultEdit.Error != nil {
			return c.JSON(http.StatusOK, helper.ResultResponse(true, "Update Cart Failed", resultEdit.Error))
		}
		return c.JSON(http.StatusOK, helper.ResultResponse(false, "Update Cart Success", &updateItem))
	}
}

func DeleteItemCart(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var deleteItem entity.Cart_Products
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

		deleteItem.CartID = user.Id
		deleteItem.ProductID,_ = strconv.Atoi(c.Param("product_id"))

		//Query
		query := fmt.Sprintf("DELETE FROM cart_products WHERE cart_id = %d AND product_id = %d", deleteItem.CartID, deleteItem.ProductID)
		resultDelete := db.Exec(query)
		if resultDelete.Error != nil {
			return c.JSON(http.StatusOK, helper.ResultResponse(true, "Delete Item Cart Failed", resultDelete.Error))
		}
		return c.JSON(http.StatusOK, helper.ResultResponse(false, "Delete Item Cart Success", &deleteItem))
	}
}
