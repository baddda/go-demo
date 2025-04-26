package model

type Task struct {
	ID          uint   `gorm:"primarykey"`
	Description string `json:"description"`
}
