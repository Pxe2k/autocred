package storage

import "gorm.io/gorm"

type DocumentBusiness struct {
	gorm.Model
	DocumentType              string `gorm:"size:100" json:"documentType"`      // Тип документа
	IIN                       string `gorm:"size:100" json:"IIN"`               // ИИН
	IssuingAuthority          string `gorm:"size:100" json:"issuingAuthority"`  // Орган выдачи
	DocumentNumber            string `gorm:"size:100" json:"documentNumber"`    // Номер документа
	DocumentIssueDate         string `gorm:"size:100" json:"documentIssueDate"` // Дата выдачи
	DocumentEndDate           string `gorm:"size:100" json:"documentEndDate"`   // Дата истечения
	BeneficialOwnerBusinessID uint
}
