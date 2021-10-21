package auth

type Account struct {
	Id        float64 `gorm:"id, primarykey, autoincrement"`
	Email     string  `gorm:"email; unique"`
	Password  string  `gorm:"password"`
	CreatedAt int64   `gorm:"created_at"`
}
