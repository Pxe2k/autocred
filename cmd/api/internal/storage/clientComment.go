package storage

import "gorm.io/gorm"

type ClientComment struct {
	gorm.Model

	ClientID uint
}
