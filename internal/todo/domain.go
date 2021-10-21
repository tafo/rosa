package todo

type Item struct {
	Id    int    `gorm:"id, primarykey, autoincrement"`
	Content string `gorm:"content"`
	IsCompleted bool `gorm:"is_completed"`
}