package storage

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Title string `gorm:"size:100;" json:"title"`
}
