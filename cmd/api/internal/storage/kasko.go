package storage

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Kasko struct {
	gorm.Model
	Title   string  `gorm:"size:255;" json:"title"`
	Price   int     `json:"price"`
	Percent float64 `json:"percent"`
	BankID  uint    `json:"bankID"`
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

func (k *Kasko) Get(db *gorm.DB, id uint) (*Kasko, error) {
	err := db.Debug().Model(&Kasko{}).Where("id = ?", id).First(&k).Error
	if err != nil {
		return nil, err
	}
	return k, nil
}

func (k *Kasko) SoftDelete(db *gorm.DB, id uint) (int64, error) {
	err := db.Debug().Model(&Kasko{}).Where("id = ?", id).Take(&Kasko{}).Select(clause.Associations).Delete(&Kasko{})
	if err != nil {
		return 0, err.Error
	}

	if err.Error != nil {
		return 0, err.Error
	}

	return err.RowsAffected, nil
}
