package entity

type Cart struct {
	ID       uint64 `gorm:"primarykey"`
	UserID   uint64
	Products []Product `gorm:"many2many:cart_products;"`
}

type CartProductsList struct {
	ID uint64
	Product_Name string
	Product_Price	int
	Product_Image string
	Amount int
}