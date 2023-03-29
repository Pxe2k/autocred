package storage

import (
	"gorm.io/gorm"
)

type City struct {
	gorm.Model
	Name     string  `gorm:"size:100;" json:"name"`
	Kato     string  `gorm:"size:100;" json:"kato"`
	ParentID *uint   `json:"parentID,omitempty"`
	Area     *[]City `gorm:"foreignkey:ParentID" json:"area,omitempty"`
}

func (c *City) All(db *gorm.DB) (*[]City, error) {
	var cities []City
	err := db.Debug().Model(&City{}).Where("parent_id IS NULL").Limit(100).Find(&cities).Error
	if err != nil {
		return nil, err
	}

	return &cities, nil
}

func (c *City) Find(db *gorm.DB, id uint) (*[]City, error) {
	var cities []City
	err := db.Debug().Model(&City{}).Where("parent_id = ?", id).Limit(100).Find(&cities).Error
	if err != nil {
		return nil, err
	}

	return &cities, nil
}
