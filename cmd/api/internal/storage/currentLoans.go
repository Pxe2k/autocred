package storage

import "gorm.io/gorm"

type CurrentLoans struct {
	gorm.Model
	BankTitle           string `gorm:"size:100;" json:"bankTitle"`
	AgreementNumber     string `gorm:"size:100;" json:"agreementNumber"`
	Sum                 int    `json:"sum"`
	BalanceOwed         int    `json:"balanceOwed"`
	MonthlyPayment      int    `json:"monthlyPayment"`
	InterestRate        string `json:"interestRate"`
	PresenceOfDelays    bool   `json:"presenceOfDelays"`
	FrequencyOfPayments string `gorm:"size:100;" json:"frequencyOfPayments"`
	DateLastPayment     string `gorm:"size:100;" json:"dateLastPayment"`
	ClientID            uint
}
