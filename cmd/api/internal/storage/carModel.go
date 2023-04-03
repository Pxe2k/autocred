package storage

import "gorm.io/gorm"

type CarModel struct {
	gorm.Model
	Title      string `gorm:"size:100;" json:"title"`
	CarBrandID uint
}
