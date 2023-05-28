package storage

import "gorm.io/gorm"

type RelationWithBank struct {
	gorm.Model
	AmountOfCredits  string `gorm:"size:100;" json:"amountOfCredits"` // Количество текущих кредитов
	AmountOfPayments int    `json:"amountOfPayments"`                 // Сумма ежемесячных платежей
	Delay            string `gorm:"size:100;" json:"delay"`           // Текущая просрочка свыше 7 дней
	ClientID         uint
}
