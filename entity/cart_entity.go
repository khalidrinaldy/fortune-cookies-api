package entity

type Cart struct {
	ID       int `gorm:"primarykey"`
	UserID   int
	Products []Product `gorm:"many2many:cart_products;"`
}

type CartProductsList struct {
	ID            int
	Product_Id    int
	Product_Name  string
	Product_Price int
	Product_Image string
	Amount        int
}
