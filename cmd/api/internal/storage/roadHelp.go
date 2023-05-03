package storage

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type RoadHelp struct {
	gorm.Model
	Title   string  `gorm:"size:255;" json:"title"`
	Price   int     `json:"price"`
	Percent float64 `json:"percent"`
	BankID  uint    `json:"bankID"`
}

func (r *RoadHelp) Save(db *gorm.DB) (*RoadHelp, error) {
	err := db.Debug().Create(&r).Error
	if err != nil {
		return &RoadHelp{}, err
	}

	return r, nil
}

func (r *RoadHelp) Update(db *gorm.DB, id int) (*RoadHelp, error) {
	err := db.Debug().Model(&RoadHelp{}).Where("id = ?", id).Updates(&r).Error
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (r *RoadHelp) Get(db *gorm.DB, id uint) (*RoadHelp, error) {
	err := db.Debug().Model(&RoadHelp{}).Where("id = ?", id).First(&r).Error
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (r *RoadHelp) SoftDelete(db *gorm.DB, id uint) (int64, error) {
	err := db.Debug().Model(&RoadHelp{}).Where("id = ?", id).Take(&RoadHelp{}).Select(clause.Associations).Delete(&RoadHelp{})
	if err != nil {
		return 0, err.Error
	}

	if err.Error != nil {
		return 0, err.Error
	}

	return err.RowsAffected, nil
}
