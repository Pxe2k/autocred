package storage

import "gorm.io/gorm"

type BankApplication struct {
	gorm.Model
	Bank            string `gorm:"size:100;" json:"bank"`
	CreditProductID uint   `json:"creditProductID"` // Кредитный продукт
	ReLoan          bool   `json:"reLoan"`          // Повторный займ
	CarPrice        int    `json:"carPrice"`        // Цена авто
	LoanAmount      int    `json:"loanAmount"`      // Сумма займа
	InitFee         int    `json:"initFee"`         // Первоначалка
	LoanPercentage  int    `json:"loanPercentage"`  // Процент кредита
	TrenchesNumber  int    `json:"trenchesNumber"`  // Кол-во траншей
	KaskoID         uint   `json:"kaskoID"`
	RoadHelpID      uint   `json:"roadHelpID"`
	LifeInsuranceID uint   `json:"lifeInsuranceID"`
	ApplicationID   uint
	CreditProduct   *BankProduct    `json:"creditProduct,omitempty"`
	Kasko           *Kasko          `json:"kasko,omitempty"`
	RoadHelp        *RoadHelp       `json:"roadHelp,omitempty"`
	LifeInsurance   *LifeInsurance  `json:"lifeInsurance,omitempty"`
	BankResponses   *[]BankResponse `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"bankResponses,omitempty"`
}
