package storage

import "gorm.io/gorm"

type Media struct {
	gorm.Model
	Title    string `gorm:"size:255;" json:"title"`
	File     string `gorm:"size:255;" json:"file"`
	ClientID uint
}

func (m *Media) SaveMedia(db *gorm.DB) (*Media, error) {
	err := db.Debug().Create(&m).Error
	if err != nil {
		return &Media{}, err
	}
	return m, nil
}