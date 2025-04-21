package model

type Task struct {
	ID          string `json:"id" gorm:"primary_key"`
	Description string `json:"description"`
}
