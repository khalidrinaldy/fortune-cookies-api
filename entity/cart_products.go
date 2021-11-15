package entity

type Cart_Products struct {
	ID        uint64 `gorm:"primarykey"`
	CartID    uint64 `gorm:"primaryKey"`
	ProductID uint64 `gorm:"primaryKey"`
	Amount    int
}
