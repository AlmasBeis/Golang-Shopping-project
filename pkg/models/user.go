package models

// Модель User для таблицы User
type User struct {
	ID       int     `json:"id" gorm:"primaryKey"`
	Username string  `json:"username"`
	Password string  `json:"password"`
	Balance  float32 `json:"balance"`
}
