package storage

import "gorm.io/gorm"

type CurrentLoans struct {
	gorm.Model
	BankTitle           string `gorm:"size:100;" json:"bankTitle"`           // Название банка
	AgreementNumber     string `gorm:"size:100;" json:"agreementNumber"`     // Номер договора
	Sum                 int    `json:"sum"`                                  // Общая сумма займа
	OverdraftAmount     int    `json:"overdraftAmount"`                      // Сумма c НДС
	BalanceOwed         int    `json:"balanceOwed"`                          // Остаток долг
	MonthlyPayment      int    `json:"monthlyPayment"`                       // Ежемесячный платеж
	InterestRate        string `json:"interestRate"`                         // Процентная ставка
	FrequencyOfPayments string `gorm:"size:100;" json:"frequencyOfPayments"` // Переодичность выплат
	LastPayment         string `gorm:"size:100;" json:"lastPayment"`         // Дата последнего платежа
	PresenceOfDelays    bool   `json:"presenceOfDelays"`                     // Наличие просрочек
	ClientID            uint
}
