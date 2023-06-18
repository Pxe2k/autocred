package storage

import "gorm.io/gorm"

type BankApplication struct {
	gorm.Model
	BankID          uint  `gorm:"size:100;" json:"bankID"`
	TrenchesNumber  int   `json:"trenchesNumber"`            // Кол-во траншей
	CreditProductID *uint `json:"creditProductID,omitempty"` // Кредитный продукт
	KaskoID         *uint `json:"kaskoID,omitempty"`
	RoadHelpID      *uint `json:"roadHelpID,omitempty"`
	LifeInsuranceID *uint `json:"lifeInsuranceID,omitempty"`
	LoanAmount      int   `json:"loanAmount"` // Сумма займа
	ApplicationID   uint
	CreditProduct   *BankProduct   `json:"creditProduct,omitempty"`
	Kasko           *Kasko         `json:"kasko,omitempty"`
	RoadHelp        *RoadHelp      `json:"roadHelp,omitempty"`
	LifeInsurance   *LifeInsurance `json:"lifeInsurance,omitempty"`
	Bank            Bank           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"bank"`
	BankResponse    BankResponse   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"bankResponse"`
	BankProduct     BankProduct    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"bankProduct"`
}
