package storage

import "gorm.io/gorm"

type Kasko struct {
	gorm.Model
	Title       string  `gorm:"size:255;" json:"title"`
	Price       int     `json:"price"`
	Percent     float64 `json:"percent"`
	InsuranceID uint    `json:"insuranceID"`
}

func (k *Kasko) Save(db *gorm.DB) (*Kasko, error) {
	err := db.Debug().Create(&k).Error
	if err != nil {
		return &Kasko{}, err
	}

	return k, nil
}

func (k *Kasko) Update(db *gorm.DB, id int) (*Kasko, error) {
	err := db.Debug().Model(&Kasko{}).Where("id = ?", id).Updates(&k).Error
	if err != nil {
		return nil, err
	}
	return k, nil
}
