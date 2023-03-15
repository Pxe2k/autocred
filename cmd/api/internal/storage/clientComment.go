package storage

import "gorm.io/gorm"

type ClientComment struct {
	gorm.Model
	WorkInfo   string
	Comment    string
	RiskStatus string
	ClientID   uint
}
