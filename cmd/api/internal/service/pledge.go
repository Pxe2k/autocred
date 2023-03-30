package service

import (
	"autocredit/cmd/api/internal/storage"
	"encoding/json"

	"gorm.io/gorm"
)

func CreatePledgeService(db *gorm.DB, body []byte) (*storage.Pledge, error) {
	pledge := storage.Pledge{}
	err := json.Unmarshal(body, &pledge)
	if err != nil {
		return &storage.Pledge{}, err
	}

	createdPledge, err := pledge.Save(db)
	if err != nil {
		return &storage.Pledge{}, err
	}

	return createdPledge, nil
}
