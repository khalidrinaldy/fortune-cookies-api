package entity

import (
	"time"

	"gorm.io/gorm"
)

type History struct {
	gorm.Model
	UserId      int       `gorm:"not null"`
	Products    []Product `gorm:"many2many:history_products;"`
	Tanggal     time.Time `gorm:"not null"`
	Address     string    `gorm:"not null"`
	Total_price int       `gorm:"not null"`
}
