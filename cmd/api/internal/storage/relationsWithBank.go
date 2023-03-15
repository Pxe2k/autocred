package storage

import "gorm.io/gorm"

type RelationWithBank struct {
	gorm.Model
	AmountOfCredits  string `gorm:"size:100;"`
	AmountOfPayments int
	OverdraftAmount  int
	Delay            string `gorm:"size:100;"`
	ClientID         uint
}
