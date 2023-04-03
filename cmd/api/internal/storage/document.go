package storage

import "gorm.io/gorm"

type Document struct {
	gorm.Model
	Type              string `gorm:"size:100;" json:"type"`              // Тип документа
	IIN               string `gorm:"size:100;" json:"IIN"`               // ИИН
	Number            string `gorm:"size:100;" json:"number"`            // Номер
	DocumentIssueDate string `gorm:"size:100;" json:"documentIssueDate"` // Дата выдачи
	DocumentEndDate   string `gorm:"size:100;" json:"documentEndDate"`   // Дата истечения
	IssuingAuthority  string `gorm:"size:100;" json:"issuingAuthority"`  // Орган выдачи
	PlaceOfBirth      string `gorm:"size:100;" json:"placeOfBirth"`      // Место рождения
	ClientID          uint
}
