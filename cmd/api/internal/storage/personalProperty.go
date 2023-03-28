package storage

import "gorm.io/gorm"

type PersonalProperty struct {
	gorm.Model
	Type           string `gorm:"size:100;" json:"type"`           // Вид имущества
	CarBrand       string `gorm:"size:100;" json:"brand"`          // Марка
	CarModel       string `gorm:"size:100;" json:"model"`          // Модель
	YearOfIssue    string `gorm:"size:100;" json:"yearOfIssue"`    // Год выпуска
	Price          int    `json:"price"`                           // Стоимость
	DocumentSerial string `gorm:"size:100;" json:"documentSerial"` // Серия документа
	DocumentNumber string `gorm:"size:100;" json:"documentNumber"` // Номер документа
	PurchaseMethod string `json:"purchaseMethod"`                  // Способ приобретения
	Description    string `json:"description"`                     // Описание
	Document       string `gorm:"size:100;" json:"document"`       // Документ
	ClientID       uint
}
