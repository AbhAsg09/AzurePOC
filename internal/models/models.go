package models

type Video struct {
	ID    uint `gorm:"primaryKey"`
	Name  string
	Age   string
	City  string
	Phone string
}
