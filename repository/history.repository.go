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

func GetAllHistory(db *gorm.DB) echo.HandlerFunc{
	return func(c echo.Context) error {
		var history []entity.HistoryList
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

		resultHistory := db.Raw("select id, created_at , address, total_price from histories where user_id = ? order by created_at desc;", user.Id).Scan(&history)
		if resultHistory.Error != nil {
			return c.JSON(http.StatusOK, helper.ResultResponse(true, "Error Occured While Querying SQL", result.Error))
		}
		return c.JSON(http.StatusOK, helper.ResultResponse(false, "Fetch history success", &history))
	}
}

func GetDetailHistoryProducts(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var products []entity.DetailHistory
		
		//Query
		result := db.Raw(`select products.product_name, products.product_image, products.product_price, history_products.amount from history_products
		join products
		on products.id = history_products.product_id
		where history_products.history_id = ?`, c.Param("history_id")).Scan(&products)
		if result.Error != nil {
			return c.JSON(http.StatusOK, helper.ResultResponse(true, "Error Occured While Querying SQL", result.Error))
		}
		return c.JSON(http.StatusOK, helper.ResultResponse(false, "Fetch Detail history success", &products))
	}
}

func Purchase(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var user entity.User
		var history entity.History

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

		//Query create history
		history.UserId = user.Id
		history.Address = c.FormValue("address")
		history.Total_price,_ = strconv.Atoi(c.FormValue("total_price"))
		resultHistory := db.Create(&history)
		if resultHistory.Error != nil {
			return c.JSON(http.StatusOK, helper.ResultResponse(true, "Error Occured While Querying SQL", resultHistory.Error))
		}

		//Query create history products
		array := strings.Split(c.FormValue("products_id"), ",")
		amounts := strings.Split(c.FormValue("amounts"), ",")
		var products = []entity.History_products{}
		for index, item := range array {
			var conv,_ = strconv.Atoi(item)
			var convAmount,_ = strconv.Atoi(amounts[index])
			products = append(products, entity.History_products{
				HistoryID: int(history.ID),
				ProductID: conv,
				Amount: convAmount,
			})
		}
		resultAdd := db.Create(&products)
		if resultAdd.Error != nil {
			return c.JSON(http.StatusOK, helper.ResultResponse(true, "Error Occured While Querying SQL", resultAdd.Error))
		}

		//Query
		resultDelete := db.Exec("delete from cart_products where cart_id = ?", user.Id)
		if resultDelete.Error != nil {
			return c.JSON(http.StatusOK, helper.ResultResponse(true, "Error Occured While Querying SQL", resultDelete.Error))
		}

		return c.JSON(http.StatusOK, helper.ResultResponse(false, "Purchase Success", &history))
	}
}

func CountHistory(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var count int64

		//QUERY
		result := db.Table("histories").Count(&count)
		if result.Error != nil {
			return c.JSON(http.StatusOK, helper.ResultResponse(true, "Error Occured While Querying SQL", result.Error.Error()))
		}
		return c.JSON(http.StatusOK, helper.ResultResponse(false, "Count Histories Success", &count))
	}
}