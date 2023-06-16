package storage

import "gorm.io/gorm"

type BankResponse struct {
	gorm.Model
	Status            string `json:"status"`
	Description       string `json:"description"`
	ApplicationID     string `json:"applicationID"`
	BankApplicationID uint   `json:"bankApplicationID"`
}

func (b *BankResponse) Save(db *gorm.DB) (*BankResponse, error) {
	err := db.Debug().Create(&b).Error
	if err != nil {
		return nil, err
	}

	return b, nil
}
