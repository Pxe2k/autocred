package storage

import "gorm.io/gorm"

type ClientComment struct {
	gorm.Model
	Comment  string `json:"comment"`
	ClientID uint
}
