package storage

import "gorm.io/gorm"

type BankApplication struct {
	gorm.Model
	Bank            string `gorm:"size:100;" json:"bank"`
	TrenchesNumber  int    `json:"trenchesNumber"`  // Кол-во траншей
	CreditProductID uint   `json:"creditProductID"` // Кредитный продукт
	KaskoID         *uint  `json:"kaskoID,omitempty"`
	RoadHelpID      *uint  `json:"roadHelpID,omitempty"`
	LifeInsuranceID *uint  `json:"lifeInsuranceID,omitempty"`
	LoanAmount      int    `json:"loanAmount"` // Сумма займа
	ApplicationID   uint
	CreditProduct   *BankProduct   `json:"creditProduct,omitempty"`
	Kasko           *Kasko         `json:"kasko,omitempty"`
	RoadHelp        *RoadHelp      `json:"roadHelp,omitempty"`
	LifeInsurance   *LifeInsurance `json:"lifeInsurance,omitempty"`
	BankResponses   []BankResponse `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"bankResponses"`
}
