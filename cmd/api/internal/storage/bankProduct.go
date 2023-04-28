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

func (b *BankProduct) Update(db *gorm.DB, id int) (*BankProduct, error) {
	err := db.Debug().Model(&BankProduct{}).Where("id = ?", id).Updates(&b).Error
	if err != nil {
		return nil, err
	}
	return b, nil
}
