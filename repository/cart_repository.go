package repository

import (
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
		result := db.Raw("select cart_products.id, products.product_name, products.product_price, products.product_image, cart_products.amount from cart_products join products on products.id = cart_products.product_id join carts on carts.id = cart_products.cart_id where carts.user_id = ?;", c.Param("id")).
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
		addCart.CartID, _ = strconv.ParseUint(c.FormValue("cart_id"), 10, 64)
		addCart.ProductID, _ = strconv.ParseUint(c.FormValue("product_id"), 10, 64)
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
		updateItem.CartID, _ = strconv.ParseUint(c.FormValue("cart_id"), 10, 64)
		updateItem.ProductID, _ = strconv.ParseUint(c.FormValue("product_id"), 10, 64)
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
		result := db.Raw("DELETE FROM users WHERE cart_id = ? AND product_id = ?", c.Param("id"), c.FormValue("product_id"))
		if result.Error != nil {
			return c.JSON(http.StatusOK, helper.ResultResponse(true, "Delete Item Cart Failed", &deleteItem))
		}
		return c.JSON(http.StatusOK, helper.ResultResponse(false, "Delete Item Cart Success", &deleteItem))
	}
}
