package storage

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type AutoDealer struct {
	gorm.Model
	Title   string `gorm:"size:100;" json:"title"`
	Address string `gorm:"size:100;" json:"address"`
	Users   []User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"users"`
}

func (a *AutoDealer) Save(db *gorm.DB) (*AutoDealer, error) {
	err := db.Debug().Create(&a).Error
	if err != nil {
		return &AutoDealer{}, err
	}

	return a, nil
}

func (a *AutoDealer) Update(db *gorm.DB, id int) (*AutoDealer, error) {
	err := db.Debug().Model(&AutoDealer{}).Where("id = ?", id).Updates(&a).Error
	if err != nil {
		return nil, err
	}
	return a, nil
}

func (a *AutoDealer) Get(db *gorm.DB, id uint) (*AutoDealer, error) {
	err := db.Debug().Model(&AutoDealer{}).Preload("Users").Where("id = ?", id).First(&a).Error
	if err != nil {
		return nil, err
	}
	return a, nil
}

func (a *AutoDealer) SoftDelete(db *gorm.DB, id uint) (int64, error) {
	err := db.Debug().Model(&AutoDealer{}).Where("id = ?", id).Take(&AutoDealer{}).Select(clause.Associations).Delete(&AutoDealer{})
	if err != nil {
		return 0, err.Error
	}

	if err.Error != nil {
		return 0, err.Error
	}

	return err.RowsAffected, nil
}

func (a *AutoDealer) All(db *gorm.DB) (*[]AutoDealer, error) {
	var autoDealers []AutoDealer
	err := db.Debug().Model(&AutoDealer{}).Limit(100).Find(&autoDealers).Error
	if err != nil {
		return nil, err
	}

	return &autoDealers, nil
}
