package models

type Account struct {
	Id    int    `gorm:"id, primarykey, autoincrement" json:"id"`
	Email string `gorm:"email; unique" json:"email"`
	Password  string `gorm:"password" json:"password"`
	CreatedAt int64  `gorm:"created_at"`
}
