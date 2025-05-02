package models

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Username string `gorm:"unique" json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Role     string `json:"role" gorm:"default:'user'"` // Добавляем поле для роли, по умолчанию 'user'
}
