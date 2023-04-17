package storage

import "gorm.io/gorm"

type Document struct {
	gorm.Model

	ClientID uint
}

func (d *Document) Update(db gorm.DB, document Document) (*Document, error) {
	err := db.Debug().Model(&Document{}).Where("id = ?", document.ID).Updates(document).Error
	if err != nil {
		return nil, err
	}

	return d, nil
}
