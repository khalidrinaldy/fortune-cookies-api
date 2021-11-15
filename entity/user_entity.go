package entity

type User struct {
	ID       uint64 `gorm:"primarykey"`
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token,omitempty"`
	Address  string `json:"address"`
	Cart     Cart
}
