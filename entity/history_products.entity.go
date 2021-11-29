package entity

type History_products struct {
	Id        int `gorm:"primarykey"`
	HistoryID int
	ProductID int
	Amount    int
}
