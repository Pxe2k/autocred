package storage

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CarBrand struct {
	gorm.Model
	Title     string     `gorm:"size:100;" json:"title"`
	CarModels []CarModel `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"carModels"`
}

func (c *CarBrand) Save(db *gorm.DB) (*CarBrand, error) {
	err := db.Debug().Create(&c).Error
	if err != nil {
		return &CarBrand{}, err
	}
	return c, nil
}

func (c *CarBrand) All(db *gorm.DB) (*[]CarBrand, error) {
	var carBrands []CarBrand

	err := db.Debug().Model(&CarBrand{}).Preload(clause.Associations).Limit(100).Find(&carBrands).Error
	if err != nil {
		return nil, err
	}

	return &carBrands, nil
}
