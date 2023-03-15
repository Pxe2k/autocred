package storage

import "gorm.io/gorm"

type ClientContact struct {
	gorm.Model
	RelationType string `gorm:"size:100;"`
	FullName     string `gorm:"size:100;"`
	Phone        string `gorm:"size:100;"`
	ClientID     uint
}
