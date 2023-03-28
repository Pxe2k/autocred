package storage

import "gorm.io/gorm"

type BankProduct struct {
	gorm.Model
	Title  string `gorm:"size:100;" json:"title"`
	BankID uint   `json:"bankID"`
}

func (b *BankProduct) Save(db *gorm.DB) (*BankProduct, error) {
	err := db.Debug().Create(&b).Error
	if err != nil {
		return &BankProduct{}, err
	}

	return b, nil
}
