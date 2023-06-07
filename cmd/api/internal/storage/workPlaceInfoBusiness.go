package storage

import "gorm.io/gorm"

type WorkPlaceInfoBusiness struct {
	gorm.Model
	JobTitle                  string `gorm:"size:100" json:"jobTitle"`       // Должность
	Stake                     string `gorm:"size:100" json:"stake"`          // Доля в уставном капитале
	WorkExperience            string `gorm:"size:100" json:"workExperience"` // Опыт работы
	BeneficialOwnerBusinessID uint
}
