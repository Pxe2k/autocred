package storage

import "gorm.io/gorm"

type Application struct {
	gorm.Model
	CreditProduct      string `gorm:"size:100;" json:"creditProduct"`
	ReLoan             bool   `json:"reLoan"`
	LoanAmount         int    `json:"loanAmount"`
	Score              int    `json:"score"`
	Subsidy            int    `json:"subsidy"`
	LoanPurpose        string `gorm:"size:100;" json:"loanPurpose"`
	PrincipalAmount    int    `json:"principalAmount"`
	PeriodLength       int    `json:"periodLength"`
	TrenchesNumber     int    `json:"trenchesNumber"`
	CarryoverContracts bool   `json:"carryoverContracts"`
	InterestRate       int    `json:"interestRate"`
	InitialFee         int    `json:"initialFee"`
	LoanType           string `gorm:"size:100;" json:"loanType"`
	LoanStage          int    `json:"loanStage"`
	Income             int    `json:"income"`
	Payment            int    `json:"payment"`
	DebtLoad           int    `json:"debtLoad"`
	UserID             uint   `json:"userID"`
	CreditorID         uint   `json:"creditorID"`
	ClientID           uint   `json:"clientID"`
	Client             Client `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"client"`
}

func (a *Application) Save(db *gorm.DB) (*Application, error) {
	err := db.Debug().Create(&a).Error
	if err != nil {
		return &Application{}, err
	}

	return a, nil
}
