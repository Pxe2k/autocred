package storage

import "gorm.io/gorm"

type Document struct {
	gorm.Model
	Type              string `gorm:"size:100;"`
	Serial            string `gorm:"size:100;"`
	Number            string `gorm:"size:100;"`
	IssuingAuthority  string `gorm:"size:100;"`
	DocumentIssueDate string `gorm:"size:100;"`
	DocumentEndDate   string `gorm:"size:100;"`
	PlaceOfBirth      string `gorm:"size:100;"`
	ClientID          uint
}
