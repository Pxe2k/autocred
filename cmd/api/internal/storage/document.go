package storage

import "gorm.io/gorm"

type Document struct {
	gorm.Model
	Type               string `gorm:"size:100" json:"type"`              // Тип документа
	IIN                string `gorm:"size:100;unique" json:"IIN"`        // ИИН
	Number             string `gorm:"size:100" json:"number"`            // Номер
	IssuingAuthority   string `gorm:"size:100" json:"issuingAuthority"`  // Орган выдачи
	PlaceOfBirth       string `gorm:"size:100" json:"placeOfBirth"`      // Место рождения
	DocumentIssueDate  string `gorm:"size:100" json:"documentIssueDate"` // Дата выдачи
	DocumentEndDate    string `gorm:"size:100" json:"documentEndDate"`   // Дата истечения
	IndividualClientID uint
}

func (d *Document) Update(db *gorm.DB, document *Document, clientID uint) error {
	err := db.Debug().Model(&Document{}).Where("individual_client_id = ?", clientID).Updates(document).Error
	if err != nil {
		return err
	}

	return nil
}
