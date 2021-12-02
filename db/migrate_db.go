package db

import (
	"fortune-cookies/entity"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&entity.User{})
	db.AutoMigrate(&entity.Cart{})
	db.AutoMigrate(&entity.Product{})
	db.AutoMigrate(&entity.Cart_Products{})
	db.AutoMigrate(&entity.History{})
	db.AutoMigrate(&entity.History_products{})
	db.AutoMigrate(&entity.Admin{})
}