package entity

type Cart_Products struct {
	ID        int `gorm:"primarykey"`
	CartID    int `gorm:"primaryKey"`
	ProductID int `gorm:"primaryKey"`
	Amount    int
}
