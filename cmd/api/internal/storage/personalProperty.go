package storage

import "gorm.io/gorm"

type PersonalProperty struct {
	gorm.Model

	ClientID uint
}
