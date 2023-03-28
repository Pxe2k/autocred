package storage

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type City struct {
	gorm.Model
	Name      string     `gorm:"size:100;" json:"name"`
	Districts []District `json:"districts"`
}

func (c *City) Save(db *gorm.DB) (*City, error) {
	err := db.Debug().Create(&c).Error
	if err != nil {
		return &City{}, err
	}

	return c, nil
}

func (c *City) All(db *gorm.DB) (*[]City, error) {
	var cities []City
	err := db.Debug().Model(&City{}).Preload(clause.Associations).Limit(100).Find(&cities).Error
	if err != nil {
		return nil, err
	}

	return &cities, nil
}
