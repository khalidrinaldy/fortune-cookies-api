package entity

type User struct {
	Id       int    `gorm:"primarykey"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token,omitempty"`
	Address  string `json:"address"`
	Cart     Cart
}
