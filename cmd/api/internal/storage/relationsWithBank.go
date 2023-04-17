package storage

import "gorm.io/gorm"

type RelationWithBank struct {
	gorm.Model

	ClientID uint
}
