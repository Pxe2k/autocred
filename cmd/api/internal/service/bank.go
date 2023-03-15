package service

import (
	"autocredit/cmd/api/internal/storage"
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
)

func SignApplication(db *gorm.DB, body []byte) (storage.BankResponse, error) {
	bankResponse := storage.BankResponse{}
	err := json.Unmarshal(body, &bankResponse)
	if err != nil {
		fmt.Println(err)
		return storage.BankResponse{}, err
	}

	signedResponse, err := bankResponse.Save(db)
	if err != nil {
		fmt.Println(err)
		return storage.BankResponse{}, err
	}

	return *signedResponse, nil
}
