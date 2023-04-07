package storage

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IssuingAuthority struct {
	ID    uint   `json:"ID"`
	Title string `json:"Title"`
}

func (i *IssuingAuthority) All(db *gorm.DB) (*[]IssuingAuthority, error) {
	var issuingAuthorities []IssuingAuthority
	err := db.Debug().Model(&IssuingAuthority{}).Preload(clause.Associations).Limit(100).Find(&issuingAuthorities).Error
	if err != nil {
		return nil, err
	}

	return &issuingAuthorities, nil
}
