package storage

import "gorm.io/gorm"

type MaritalStatus struct {
	gorm.Model
	Status             string `gorm:"size:100;" json:"status"`            // Семейное положение
	FamilyPartnerName  string `gorm:"size:100;" json:"familyPartnerName"` // ФИО партнера
	Phone              string `gorm:"size:100;" json:"phone"`             // Телефон
	MinorChildren      string `gorm:"size:100;" json:"minorChildren"`     // Кол-во несовершеннолетних детей
	IndividualClientID uint
}

func (m *MaritalStatus) Update(db gorm.DB, status MaritalStatus) (*MaritalStatus, error) {
	err := db.Debug().Model(&MaritalStatus{}).Where("id = ?", status.ID).Updates(status).Error
	if err != nil {
		return nil, err
	}

	return m, nil
}
