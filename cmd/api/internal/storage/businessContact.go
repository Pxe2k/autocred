package storage

import "gorm.io/gorm"

type BusinessContact struct {
	gorm.Model
	RelationType              string `gorm:"size:100;" json:"relationType"`
	FullName                  string `gorm:"size:100;" json:"fullName"`
	Phone                     string `gorm:"size:100;" json:"phone"`
	BeneficialOwnerBusinessID uint
}
