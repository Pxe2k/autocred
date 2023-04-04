package storage

import "gorm.io/gorm"

type BankResponse struct {
	gorm.Model
	Sign              bool   `json:"sign"`
	Comment           string `json:"comment"`
	BankID            uint   `json:"bankID"`
	BankApplicationID uint   `json:"bankApplicationID"`
}

func (r *BankResponse) Save(db *gorm.DB) (*BankResponse, error) {
	err := db.Debug().Create(&r).Error
	if err != nil {
		return &BankResponse{}, err
	}

	return r, nil
}
