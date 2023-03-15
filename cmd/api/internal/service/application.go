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

func GetApplication(db *gorm.DB, id uint, uid uint) (storage.Application, error) {
	user := storage.User{}
	userGotten, err := user.Get(db, uid)
	if err != nil {
		return storage.Application{}, err
	}

	creditor := false
	if userGotten.BankID != nil {
		creditor = true
	}

	application := storage.Application{}
	applicationGotten, err := application.Get(db, id, creditor)
	if err != nil {
		return storage.Application{}, err
	}

	return *applicationGotten, nil
}
