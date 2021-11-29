package entity

type User struct {
	Id       int       `gorm:"primarykey"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Token    string    `json:"token,omitempty"`
	Cart     Cart      `json:"omitempty"`
	History  []History `gorm:"foreignKey:UserId;references:Id" json:",omitempty"`
}
