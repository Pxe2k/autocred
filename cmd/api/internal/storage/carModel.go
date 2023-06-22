package storage

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CarModel struct {
	gorm.Model
	Title      string `gorm:"size:100;" json:"title"`
	CarBrandID uint   `json:"carBrandID"`
}

func (c *CarModel) Save(db *gorm.DB) (*CarModel, error) {
	err := db.Debug().Create(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (c *CarModel) Update(db *gorm.DB, id int) (*CarModel, error) {
	err := db.Debug().Model(&CarModel{}).Where("id = ?", id).Updates(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (c *CarModel) Get(db *gorm.DB, id uint) (*CarModel, error) {
	err := db.Debug().Model(&CarModel{}).Where("id = ?", id).First(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (c *CarModel) SoftDelete(db *gorm.DB, id uint) (int64, error) {
	err := db.Debug().Model(&CarModel{}).Where("id = ?", id).Take(&CarModel{}).Select(clause.Associations).Delete(&CarModel{})
	if err != nil {
		return 0, err.Error
	}

	if err.Error != nil {
		return 0, err.Error
	}

	return err.RowsAffected, nil
}
