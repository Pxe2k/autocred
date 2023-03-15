package storage

import "gorm.io/gorm"

type MaritalStatus struct {
	gorm.Model
	Status             string `gorm:"size:100;"`
	FamilyPartnerName  string `gorm:"size:100;"`
	Phone              string `gorm:"size:100;"`
	NumberOfDependents string `gorm:"size:100;"`
	MinorChildren      string `gorm:"size:100;"`
	ClientID           uint
}
