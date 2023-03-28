package storage

import "gorm.io/gorm"

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
