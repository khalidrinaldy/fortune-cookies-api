package entity

import (
	"time"

	"gorm.io/gorm"
)

type History struct {
	gorm.Model
	UserId      int       `gorm:"not null"`
	Products    []Product `gorm:"many2many:history_products;"`
	Address     string    `gorm:"not null"`
	Total_price int       `gorm:"not null"`
}

type HistoryList struct {
	Id          int
	CreatedAt   time.Time
	Address     string
	Total_price int
}

type DetailHistory struct {
	Product_name  string
	Product_image string
	Product_price int
	Amount        int
}
