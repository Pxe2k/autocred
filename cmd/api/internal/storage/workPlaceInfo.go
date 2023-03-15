package storage

import "gorm.io/gorm"

type WorkPlaceInfo struct {
	gorm.Model
	Address          string `gorm:"size:100;"`
	MonthlyIncome    int
	OrganizationName string `gorm:"size:100;"`
	JobTitle         string `gorm:"size:100;"`
	BonusIncome      int
	ClientID         uint
}
