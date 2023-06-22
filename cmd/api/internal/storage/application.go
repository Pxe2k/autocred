package storage

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Application struct {
	gorm.Model
	CarBrand                string                   `gorm:"size:100;" json:"carBrand"`
	CarModel                string                   `gorm:"size:100;" json:"carModel"`
	OriginCountry           string                   `gorm:"size:100;" json:"originCountry"`
	YearIssue               string                   `gorm:"size:100;" json:"yearIssue"`
	Condition               bool                     `gorm:"size:100;" json:"condition"`
	LoanPurpose             string                   `gorm:"size:100;" json:"loanPurpose"` // Цель кредита
	Subsidy                 bool                     `json:"subsidy"`                      // Субсудия
	CarPrice                int                      `json:"carPrice"`                     // Цена авто
	InitFee                 int                      `json:"initFee"`                      // Первоначалка
	LoanPercentage          int                      `json:"loanPercentage"`               // Процент кредита
	BankApplications        []BankApplication        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"bankApplications"`
	UserID                  uint                     `json:"userID"`
	IndividualClientID      uint                     `json:"individualClientID,omitempty"`
	BusinessClientID        *uint                    `json:"businessClientID,omitempty"`
	IndividualClient        *IndividualClient        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"individualClient,omitempty"`
	BusinessClient          *BusinessClient          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"businessClient,omitempty"`
	BankProcessingDocuments []BankProcessingDocument `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"bankProcessingDocuments"`
}

func (a *Application) Save(db *gorm.DB) (*Application, error) {
	err := db.Debug().Create(&a).Error
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *Application) All(db *gorm.DB, uid uint) (*[]Application, error) {
	var applications []Application
	err := db.Debug().Model(&Application{}).Where("user_id = ?", uid).Order("created_at DESC").Preload("BankApplications").Preload("BankApplications.BankResponse").Preload("BankApplications.BankProduct").Preload("BankApplications.Bank").Preload("IndividualClient").Limit(100).Find(&applications).Error
	if err != nil {
		return nil, err
	}

	return &applications, nil
}

func (a *Application) Get(db *gorm.DB, id uint) (*Application, error) {
	err := db.Debug().Model(&Application{}).Preload("IndividualClient").Where("id = ?", id).Preload(clause.Associations).Take(a).Error
	if err != nil {
		return nil, err
	}

	return a, nil
}
