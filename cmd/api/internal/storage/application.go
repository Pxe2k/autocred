package storage

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Application struct {
	gorm.Model
	Bank               string          `gorm:"size:100;" json:"bank"`
	CreditProduct      string          `gorm:"size:100;" json:"creditProduct"` // Кредитный продукт
	ReLoan             bool            `json:"reLoan"`                         // Повторный займ
	LoanAmount         int             `json:"loanAmount"`                     // Сумма займа
	InitFee            int             `json:"initFee"`                        // Первоначалка
	LoanPercentage     int             `json:"loanPercentage"`                 // Процент кредита
	Subsidy            bool            `json:"subsidy"`                        // Субсудия
	LoanPurpose        string          `gorm:"size:100;" json:"loanPurpose"`   // Цель кредита
	TrenchesNumber     int             `json:"trenchesNumber"`                 // Кол-во траншей
	KaskoTitle         string          `gorm:"size:100;" json:"kaskoTitle"`
	KaskoPrice         int             `json:"kaskoPrice"`
	KaskoTerm          int             `json:"kaskoTerm"`
	RoadHelpTitle      string          `gorm:"size:100;" json:"roadHelpTitle"`
	RoadHelpPrice      int             `json:"roadHelpPrice"`
	RoadHelpTerm       int             `json:"roadHelpTerm"`
	LifeInsuranceTitle string          `gorm:"size:100;" json:"lifeInsuranceTitle"`
	LifeInsurancePrice int             `json:"lifeInsurancePrice"`
	LifeInsuranceTerm  int             `json:"lifeInsuranceTerm"`
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
	if !creditor {
		query.Preload(clause.Associations)
	}

	err := query.Take(a).Error
	if err != nil {
		return nil, err
	}

	return a, nil
}
