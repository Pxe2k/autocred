package storage

import (
	"errors"
	"gorm.io/gorm"
)

type Insurance struct {
	gorm.Model
	BankID        uint
	Kasko         *[]Kasko         `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"kasko,omitempty"`
	RoadHelp      *[]RoadHelp      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"roadHelp,omitempty"`
	LifeInsurance *[]LifeInsurance `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"lifeInsurance,omitempty"`
}

func (i *Insurance) Save(db *gorm.DB) (*Insurance, error) {
	err := db.Debug().Create(&i).Error
	if err != nil {
		errors.New("testsd")
		return &Insurance{}, err
	}

	return i, nil
}

func (i *Insurance) All(db *gorm.DB) (*[]Insurance, error) {
	var insurances []Insurance
	err := db.Debug().Model(&Insurance{}).Preload("Kasko").Preload("RoadHelp").Preload("LifeInsurance").Limit(100).Find(&insurances).Error
	if err != nil {
		return nil, err
	}

	return &insurances, nil
}
