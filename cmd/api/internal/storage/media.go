package storage

import "gorm.io/gorm"

type Media struct {
	gorm.Model
	Title              string `gorm:"size:255;" json:"title"`
	File               string `gorm:"size:255;" json:"file"`
	IndividualClientID uint   `json:"clientID,omitempty"`
	BusinessClientID   *uint  `json:"businessClient,omitempty"`
}

func (m *Media) Save(db *gorm.DB) (*Media, error) {
	err := db.Debug().Create(&m).Error
	if err != nil {
		return &Media{}, err
	}
	return m, nil
}

func (m *Media) SoftDelete(db *gorm.DB, id uint) (int64, error) {
	err := db.Debug().Model(&Media{}).Where("id = ?", id).Take(&Media{}).Delete(&Media{})
	if err != nil {
		return 0, err.Error
	}

	if err.Error != nil {
		return 0, err.Error
	}

	return err.RowsAffected, nil
}
