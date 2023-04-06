package storage

import "gorm.io/gorm"

type BankApplication struct {
	gorm.Model
	Bank               string `gorm:"size:100;" json:"bank"`
	CreditProduct      string `gorm:"size:100;" json:"creditProduct"` // Кредитный продукт
	ReLoan             bool   `json:"reLoan"`                         // Повторный займ
	CarPrice           int    `json:"carPrice"`
	LoanAmount         int    `json:"loanAmount"`     // Сумма займа
	InitFee            int    `json:"initFee"`        // Первоначалка
	LoanPercentage     int    `json:"loanPercentage"` // Процент кредита
	TrenchesNumber     int    `json:"trenchesNumber"` // Кол-во траншей
	KaskoTitle         string `gorm:"size:100;" json:"kaskoTitle"`
	KaskoPrice         int    `json:"kaskoPrice"`
	KaskoTerm          int    `json:"kaskoTerm"`
	RoadHelpTitle      string `gorm:"size:100;" json:"roadHelpTitle"`
	RoadHelpPrice      int    `json:"roadHelpPrice"`
	RoadHelpTerm       int    `json:"roadHelpTerm"`
	LifeInsuranceTitle string `gorm:"size:100;" json:"lifeInsuranceTitle"`
	LifeInsurancePrice int    `json:"lifeInsurancePrice"`
	LifeInsuranceTerm  int    `json:"lifeInsuranceTerm"`
	ApplicationID      uint
	BankResponses      *[]BankResponse `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"bankResponses,omitempty"`
}
