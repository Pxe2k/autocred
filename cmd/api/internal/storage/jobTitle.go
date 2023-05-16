package storage

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type JobTitle struct {
	gorm.Model
	Title             string `json:"title"`
	WorkingActivityID uint   `json:"workingActivityID"`
}

func (j *JobTitle) Save(db *gorm.DB) (*JobTitle, error) {
	err := db.Debug().Create(&j).Error
	if err != nil {
		return &JobTitle{}, err
	}

	return j, nil
}

func (j *JobTitle) Get(db *gorm.DB, id uint) (*JobTitle, error) {
	err := db.Debug().Model(&JobTitle{}).Where("id = ?", id).Take(j).Error
	if err != nil {
		return nil, err
	}

	return j, nil
}

func (j *JobTitle) Update(db *gorm.DB, id int) (*JobTitle, error) {
	err := db.Debug().Model(&JobTitle{}).Where("id = ?", id).Updates(&j).Error
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (j *JobTitle) SoftDelete(db *gorm.DB, id uint) (int64, error) {
	err := db.Debug().Model(&JobTitle{}).Where("id = ?", id).Take(&JobTitle{}).Select(clause.Associations).Delete(&JobTitle{})
	if err != nil {
		return 0, err.Error
	}

	if err.Error != nil {
		return 0, err.Error
	}

	return err.RowsAffected, nil
}
