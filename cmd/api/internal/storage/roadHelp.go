package storage

import "gorm.io/gorm"

type RoadHelp struct {
	gorm.Model
	Title       string  `gorm:"size:255;" json:"title"`
	Price       int     `json:"price"`
	Percent     float64 `json:"percent"`
	InsuranceID uint
}

func (r *RoadHelp) Save(db *gorm.DB) (*RoadHelp, error) {
	err := db.Debug().Create(&r).Error
	if err != nil {
		return &RoadHelp{}, err
	}

	return r, nil
}
