package storage

import (
	"autocredit/cmd/api/helpers/requests"
	"gorm.io/gorm"
)

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

func (b *BankResponse) UpdateStatus(db *gorm.DB, data requests.UpdateBCCStatus) error {
	err := db.Debug().Model(BankResponse{}).Where("application_id = ?", data.ApplicationID).Updates(map[string]interface{}{"description": data.Description, "status": data.Status}).Error
	if err != nil {
		return err
	}

	return nil
}
