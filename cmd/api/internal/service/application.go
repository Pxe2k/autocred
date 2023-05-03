package service

import (
	"autocredit/cmd/api/internal/storage"
	"encoding/json"

	"gorm.io/gorm"
)

func CreateApplicationService(db *gorm.DB, body []byte, uid uint) (*storage.Application, error) {
	application := storage.Application{}
	err := json.Unmarshal(body, &application)
	if err != nil {
		return &storage.Application{}, err
	}

	application.UserID = uid

	createdApplication, err := application.Save(db)
	if err != nil {
		return &storage.Application{}, err
	}

	return createdApplication, nil
}

func GetApplication(db *gorm.DB, id uint) (storage.Application, error) {
	application := storage.Application{}
	applicationGotten, err := application.Get(db, id)
	if err != nil {
		return storage.Application{}, err
	}

	return *applicationGotten, nil
}
