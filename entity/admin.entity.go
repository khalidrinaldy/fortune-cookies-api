package entity

type Admin struct {
	Id       int    `gorm:"primarykey"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token,omitempty"`
}
