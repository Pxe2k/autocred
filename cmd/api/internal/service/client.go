package service

import (
	"autocredit/cmd/api/internal/storage"
	"encoding/json"
	"gorm.io/gorm"
)

func CreateClientService(db *gorm.DB, body []byte, uid uint) (*storage.Client, error) {
	client := storage.Client{}
	err := json.Unmarshal(body, &client)
	if err != nil {
		return &storage.Client{}, err
	}

	client.UserID = uid

	createdClient, err := client.Save(db)
	if err != nil {
		return &storage.Client{}, err
	}

	return createdClient, nil
}
