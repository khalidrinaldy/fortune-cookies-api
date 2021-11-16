package repository

import (
	"fmt"
	"fortune-cookies/entity"
	"fortune-cookies/helper"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func GetCartList(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var cart []entity.CartProductsList
		result := db.Raw("select cart_products.id, cart_products.product_id, products.product_name, products.product_price, products.product_image, cart_products.amount from cart_products join products on products.id = cart_products.product_id join carts on carts.id = cart_products.cart_id where carts.user_id = ?;", c.Param("id")).
				Scan(&cart)
		if result.Error != nil {
			return c.JSON(http.StatusOK, helper.ResultResponse(true, "Get Cart List Failed", ""))
		}
		return c.JSON(http.StatusOK, helper.ResultResponse(false, "Get Cart List Success", &cart))
	}
}

func AddToCart(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var addCart entity.Cart_Products
		addCart.CartID, _ = strconv.Atoi(c.FormValue("cart_id"))
		addCart.ProductID, _ = strconv.Atoi(c.FormValue("product_id"))
		addCart.Amount, _ = strconv.Atoi(c.FormValue("amount"))
		result := db.Create(&addCart)
		if result.Error != nil {
			return c.JSON(http.StatusOK, helper.ResultResponse(true, "Add To Cart Failed", &addCart))
		}
		return c.JSON(http.StatusOK, helper.ResultResponse(false, "Add To Cart Success", &addCart))
	}
}

func UpdateItemCart(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var updateItem entity.Cart_Products
		updateItem.CartID, _ = strconv.Atoi(c.FormValue("cart_id"))
		updateItem.ProductID, _ = strconv.Atoi(c.FormValue("product_id"))
		updateItem.Amount, _ = strconv.Atoi(c.FormValue("amount"))
		result := db.Model(&updateItem).Where("id = ?", c.Param("id")).Updates(entity.Cart_Products{
			CartID: updateItem.CartID,
			ProductID: updateItem.ProductID,
			Amount: updateItem.Amount,
		})
		if result.Error != nil {
			return c.JSON(http.StatusOK, helper.ResultResponse(true, "Update Cart Failed", &updateItem))
		}
		return c.JSON(http.StatusOK, helper.ResultResponse(false, "Update Cart Success", &updateItem))
	}
}

func DeleteItemCart(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var deleteItem entity.Cart_Products
		deleteItem.CartID,_ = strconv.Atoi(c.Param("id"))
		deleteItem.ProductID,_ = strconv.Atoi(c.FormValue("product_id"))
		query := fmt.Sprintf("DELETE FROM cart_products WHERE cart_id = %d AND product_id = %d", deleteItem.CartID, deleteItem.ProductID)
		result := db.Exec(query)
		if result.Error != nil {
			return c.JSON(http.StatusOK, helper.ResultResponse(true, "Delete Item Cart Failed", result.Error))
		}
		return c.JSON(http.StatusOK, helper.ResultResponse(false, "Delete Item Cart Success", &deleteItem))
	}
}
