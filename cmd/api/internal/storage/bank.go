package storage

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Bank struct {
	gorm.Model
	Title     string         `gorm:"size:100;" json:"title"`
	Image     *string        `gorm:"size:100;" json:"image,omitempty"`
	Products  *[]BankProduct `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"products,omitempty"`
	Insurance *Insurance     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"insurance,omitempty"`
}

func (b *Bank) Save(db *gorm.DB) (*Bank, error) {
	err := db.Debug().Create(&b).Error
	if err != nil {
		return &Bank{}, err
	}

	return b, nil
}

func (b *Bank) Update(db *gorm.DB, id int) (*Bank, error) {
	err := db.Debug().Model(&Bank{}).Where("id = ?", id).Updates(&b).Error
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (b *Bank) All(db *gorm.DB) (*[]Bank, error) {
	var banks []Bank
	err := db.Debug().Model(&Bank{}).Preload(clause.Associations).Preload("Insurance.Kasko").Preload("Insurance.RoadHelp").Preload("Insurance.LifeInsurance").Limit(100).Find(&banks).Error
	if err != nil {
		return nil, err
	}

	return &banks, nil
}

func (b *Bank) SoftDelete(db *gorm.DB, id uint) (int64, error) {
	err := db.Debug().Model(&Bank{}).Where("id = ?", id).Take(&Bank{}).Select(clause.Associations).Delete(&Bank{})
	if err != nil {
		return 0, err.Error
	}

	if err.Error != nil {
		return 0, err.Error
	}

	return err.RowsAffected, nil
}
