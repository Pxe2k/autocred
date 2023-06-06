package storage

import "gorm.io/gorm"

type BusinessClient struct {
	gorm.Model
	TypeOfClient string `gorm:"size:100" json:"typeOfClient"` // Тип клиента
	BIN          string `gorm:"size:100" json:"BIN"`          // БИН
	CompanyName  string `gorm:"size:100" json:"companyName"`  // Название организации
	CompanyPhone string `gorm:"size:100" json:"companyPhone"`
}
