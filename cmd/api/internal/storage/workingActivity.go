package storage

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type WorkingActivity struct {
	gorm.Model
	Title     string      `json:"title"`
	JobTitles *[]JobTitle `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"jobTitles,omitempty"`
}

func (w *WorkingActivity) Save(db *gorm.DB) (*WorkingActivity, error) {
	err := db.Debug().Create(&w).Error
	if err != nil {
		return &WorkingActivity{}, err
	}

	return w, nil
}

func (w *WorkingActivity) All(db *gorm.DB) (*[]WorkingActivity, error) {
	var workingActivities []WorkingActivity
	err := db.Debug().Model(&WorkingActivity{}).Preload(clause.Associations).Limit(100).Find(&workingActivities).Error
	if err != nil {
		return nil, err
	}

	return &workingActivities, nil
}

func (w *WorkingActivity) Update(db *gorm.DB, id int) (*WorkingActivity, error) {
	err := db.Debug().Model(&WorkingActivity{}).Where("id = ?", id).Updates(&w).Error
	if err != nil {
		return nil, err
	}
	return w, nil
}

func (w *WorkingActivity) SoftDelete(db *gorm.DB, id uint) (int64, error) {
	err := db.Debug().Model(&WorkingActivity{}).Where("id = ?", id).Take(&WorkingActivity{}).Select(clause.Associations).Delete(&WorkingActivity{})
	if err != nil {
		return 0, err.Error
	}

	if err.Error != nil {
		return 0, err.Error
	}

	return err.RowsAffected, nil
}
