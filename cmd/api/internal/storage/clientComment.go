package storage

import "gorm.io/gorm"

type ClientComment struct {
	gorm.Model
	WorkInfo   string `json:"workInfo"`
	Comment    string `json:"comment"`
	RiskStatus string `json:"riskStatus"`
	ClientID   uint
}
