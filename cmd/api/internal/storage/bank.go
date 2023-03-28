package storage

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Bank struct {
	gorm.Model
	Title    string         `gorm:"size:100;" json:"title"`
	Image    *string        `gorm:"size:100;" json:"image,omitempty"`
	Products *[]BankProduct `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"products,omitempty"`
}

func (b *Bank) Save(db *gorm.DB) (*Bank, error) {
	err := db.Debug().Create(&b).Error
	if err != nil {
		return &Bank{}, err
	}

	return b, nil
}

func (b *Bank) All(db *gorm.DB) (*[]Bank, error) {
	var banks []Bank
	err := db.Debug().Model(&Bank{}).Preload(clause.Associations).Limit(100).Find(&banks).Error
	if err != nil {
		return nil, err
	}

	return &banks, nil
}
