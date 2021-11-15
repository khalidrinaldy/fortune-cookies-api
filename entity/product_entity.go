package entity

type Product struct {
	ID                  uint64 `gorm:"primarykey"`
	Product_Name        string `json:"product_name"`
	Product_Price       int    `json:"product_price"`
	Product_Category    string `json:"product_category"`
	Product_Image       string `json:"product_image"`
	Product_Description string `json:"product_description"`
}
