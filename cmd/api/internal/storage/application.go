package storage

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Application struct {
	gorm.Model
	LoanPurpose      string            `gorm:"size:100;" json:"loanPurpose"` // Цель кредита
	Subsidy          bool              `json:"subsidy"`                      // Субсудия
	BankApplications []BankApplication `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"bankApplications"`
	UserID           uint              `json:"userID"`
	ClientID         uint              `json:"clientID"`
	Client           *Client           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"client,omitempty"`
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
