package storage

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Application struct {
	gorm.Model
	CreditProduct      string          `gorm:"size:100;" json:"creditProduct"` // Кредитный продукт
	ReLoan             bool            `json:"reLoan"`                         // Повторный займ
	LoanAmount         int             `json:"loanAmount"`                     // Сумма займа
	Score              int             `json:"score"`
	Subsidy            int             `json:"subsidy"`                      // Субсудия
	LoanPurpose        string          `gorm:"size:100;" json:"loanPurpose"` // Цель кредита
	PrincipalAmount    int             `json:"principalAmount"`              // Сумма ОС второго периода
	PeriodLength       int             `json:"periodLength"`                 // Длительность между погашениями
	TrenchesNumber     int             `json:"trenchesNumber"`               // Кол-во траншей
	CarryoverContracts bool            `json:"carryoverContracts"`           // Переносить проценты по пролонгированным контрактам
	InterestRate       int             `json:"interestRate"`                 // Процентная ставка
	InitialFee         int             `json:"initialFee"`                   // Первоначальный взнос
	LoanType           string          `gorm:"size:100;" json:"loanType"`    // Тип займа
	LoanStage          int             `json:"loanStage"`                    // Ступень займа
	Income             int             `json:"income"`                       // Среднемесячный доход
	Payment            int             `json:"payment"`                      // Среднемесячный платеж
	DebtLoad           int             `json:"debtLoad"`                     // Показатель долговой нагрузки
	UserID             uint            `json:"userID"`
	ClientID           uint            `json:"clientID"`
	Client             *Client         `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"client,omitempty"`
	BankResponses      *[]BankResponse `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"bankResponses,omitempty"`
}

func (a *Application) Save(db *gorm.DB) (*Application, error) {
	err := db.Debug().Create(&a).Error
	if err != nil {
		return &Application{}, err
	}

	return a, nil
}

func (a *Application) All(db *gorm.DB, uid uint) (*[]Application, error) {
	var applications []Application
	err := db.Debug().Model(&Application{}).Where("user_id = ?", uid).Limit(100).Find(&applications).Error
	if err != nil {
		return nil, err
	}

	return &applications, nil
}

func (a *Application) Get(db *gorm.DB, id uint, creditor bool) (*Application, error) {
	query := db.Debug().Model(&Application{}).Preload("Client").Where("id = ?", id)
	if creditor == false {
		query.Preload(clause.Associations)
	}

	err := query.Take(a).Error
	if err != nil {
		return nil, err
	}

	return a, nil
}
