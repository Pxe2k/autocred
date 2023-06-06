package storage

import "gorm.io/gorm"

type ClientContact struct {
	gorm.Model
	RelationType       string `gorm:"size:100;" json:"relationType"`
	FullName           string `gorm:"size:100;" json:"fullName"`
	Phone              string `gorm:"size:100;" json:"phone"`
	IndividualClientID uint
}
