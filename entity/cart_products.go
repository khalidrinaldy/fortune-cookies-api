package entity

type Cart_Products struct {
	ID        int `gorm:"primarykey"`
	CartID    int 
	ProductID int 
	Amount    int
}
