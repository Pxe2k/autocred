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
		return nil, err
	}
	return c, nil
}

func (c *CarBrand) All(db *gorm.DB) (*[]CarBrand, error) {
	var carBrands []CarBrand

	err := db.Debug().Model(&CarBrand{}).Preload(clause.Associations).Find(&carBrands).Error
	if err != nil {
		return nil, err
	}

	return &carBrands, nil
}

func (c *CarBrand) Update(db *gorm.DB, id int) (*CarBrand, error) {
	err := db.Debug().Model(&CarBrand{}).Where("id = ?", id).Updates(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (c *CarBrand) Get(db *gorm.DB, id uint) (*CarBrand, error) {
	err := db.Debug().Model(&CarBrand{}).Where("id = ?", id).First(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (c *CarBrand) SoftDelete(db *gorm.DB, id uint) (int64, error) {
	err := db.Debug().Model(&CarBrand{}).Where("id = ?", id).Take(&CarBrand{}).Select(clause.Associations).Delete(&CarBrand{})
	if err != nil {
		return 0, err.Error
	}

	if err.Error != nil {
		return 0, err.Error
	}

	return err.RowsAffected, nil
}
