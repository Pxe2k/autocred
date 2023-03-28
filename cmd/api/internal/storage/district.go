package storage

import "gorm.io/gorm"

type District struct {
	gorm.Model
	Name   string `gorm:"size:100;" json:"name"`
	Kato   string `gorm:"size:100;" json:"kato"`
	CityID uint   `json:"cityID"`
}

func (d *District) Save(db *gorm.DB) (*District, error) {
	err := db.Debug().Create(&d).Error
	if err != nil {
		return &District{}, err
	}

	return d, nil
}
