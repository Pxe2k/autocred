package storage

import "gorm.io/gorm"

type MaritalStatus struct {
	gorm.Model

	ClientID uint
}

func (m *MaritalStatus) Update(db gorm.DB, status MaritalStatus) (*MaritalStatus, error) {
	err := db.Debug().Model(&MaritalStatus{}).Where("id = ?", status.ID).Updates(status).Error
	if err != nil {
		return nil, err
	}

	return m, nil
}
