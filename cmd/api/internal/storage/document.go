package storage

import "gorm.io/gorm"

type Document struct {
	gorm.Model
	Type              string `gorm:"size:100;" json:"type"`              // Тип документа
	Serial            string `gorm:"size:100;" json:"serial"`            // Серия
	Number            string `gorm:"size:100;" json:"number"`            // Номер
	DocumentIssueDate string `gorm:"size:100;" json:"documentIssueDate"` // Дата выдачи
	DocumentEndDate   string `gorm:"size:100;" json:"documentEndDate"`   // Дата истечения
	PlaceOfBirth      string `gorm:"size:100;" json:"placeOfBirth"`      // Место рождения
	IssuingAuthority  string `gorm:"size:100;" json:"issuingAuthority"`  // Орган выдачи
	ClientID          uint
}
