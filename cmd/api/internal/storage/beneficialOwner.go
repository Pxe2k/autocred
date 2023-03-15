package storage

import "gorm.io/gorm"

type BeneficialOwner struct {
	gorm.Model
	FullName        string `gorm:"size:100;"`
	Bin             string `gorm:"size:100;"`
	Sex             string `gorm:"size:100;"`
	BirthDate       string `gorm:"size:100;"`
	TypeOfOwnership string `gorm:"size:100;"`
	ClientID        uint
}
