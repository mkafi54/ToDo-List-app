package models

type User struct {
	Id       string `gorm:"primaryKey;default:uuid_generate_v4()" json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (e *User) TableName() string {
	return "user"
}
