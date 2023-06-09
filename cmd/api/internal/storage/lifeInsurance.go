package storage

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type LifeInsurance struct {
	gorm.Model
	Title   string  `gorm:"size:255;" json:"title"`
	Price   int     `json:"price"`
	Percent float64 `json:"percent"`
	BankID  uint    `json:"bankID"`
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

func (l *LifeInsurance) Get(db *gorm.DB, id uint) (*LifeInsurance, error) {
	err := db.Debug().Model(&LifeInsurance{}).Where("id = ?", id).First(&l).Error
	if err != nil {
		return nil, err
	}
	return l, nil
}

func (l *LifeInsurance) SoftDelete(db *gorm.DB, id uint) (int64, error) {
	err := db.Debug().Model(&LifeInsurance{}).Where("id = ?", id).Take(&LifeInsurance{}).Select(clause.Associations).Delete(&LifeInsurance{})
	if err != nil {
		return 0, err.Error
	}

	if err.Error != nil {
		return 0, err.Error
	}

	return err.RowsAffected, nil
}
