package storage

import "gorm.io/gorm"

type LifeInsurance struct {
	gorm.Model
	Title       string  `gorm:"size:255;" json:"title"`
	Price       int     `json:"price"`
	Percent     float64 `json:"percent"`
	InsuranceID uint
}

func (l *LifeInsurance) Save(db *gorm.DB) (*LifeInsurance, error) {
	err := db.Debug().Create(&l).Error
	if err != nil {
		return &LifeInsurance{}, err
	}

	return l, nil
}

func (l *LifeInsurance) Update(db *gorm.DB, id int) (*LifeInsurance, error) {
	err := db.Debug().Model(&LifeInsurance{}).Where("id = ?", id).Updates(&l).Error
	if err != nil {
		return nil, err
	}
	return l, nil
}
