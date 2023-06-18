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

func (m *Media) MultipleSave(db *gorm.DB, media []Media) ([]Media, error) {
	err := db.Debug().Create(media).Error
	if err != nil {
		return []Media{}, err
	}
	return media, nil
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

func (m *Media) All(db *gorm.DB, uid uint) ([]Media, error) {
	var media []Media
	err := db.Debug().Model(&Media{}).Where("individual_client_id - ?", uid).Limit(100).Find(&media).Error
	if err != nil {
		return nil, err
	}

	return media, nil
}
