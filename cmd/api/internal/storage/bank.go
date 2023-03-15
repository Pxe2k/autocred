package storage

import "gorm.io/gorm"

type Bank struct {
	gorm.Model
	Title string  `gorm:"size:100;" json:"title"`
	Image *string `gorm:"size:100;" json:"image,omitempty"`
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
	err := db.Debug().Model(&Bank{}).Limit(100).Find(&banks).Error
	if err != nil {
		return nil, err
	}

	return &banks, nil
}
