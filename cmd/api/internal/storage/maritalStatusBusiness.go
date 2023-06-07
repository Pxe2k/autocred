package storage

import "gorm.io/gorm"

type MaritalStatusBusiness struct {
	gorm.Model
	Status                    string `gorm:"size:100" json:"status"`             // Семейное положение
	FamilyPartnerName         string `gorm:"size:100" json:"familyPartnerName"`  // ФИО партнера
	FamilyPartnerPhone        string `gorm:"size:100" json:"familyPartnerPhone"` // Телефон
	MinorChildren             string `gorm:"size:100" json:"minorChildren"`      // Кол-во несовершеннолетних детей
	BeneficialOwnerBusinessID uint
}
