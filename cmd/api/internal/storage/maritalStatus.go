package storage

import "gorm.io/gorm"

type MaritalStatus struct {
	gorm.Model
	Status            string `gorm:"size:100;" json:"status"`            // Семейное положение
	FamilyPartnerName string `gorm:"size:100;" json:"familyPartnerName"` // ФИО партнера
	Phone             string `gorm:"size:100;" json:"phone"`             // Телефон
	MinorChildren     string `gorm:"size:100;" json:"minorChildren"`     // Кол-во несовершеннолетних детей
	Bin               string `gorm:"size:100;" json:"bin"`               // ИИН
	ClientID          uint
}
