package storage

import "gorm.io/gorm"

type Country struct {
	ID     uint   `json:"ID"`
	Title  string `json:"Title" json:"Title,omitempty"`
	Prefix string `json:"Prefix,omitempty"`
}

func (c *Country) All(db *gorm.DB) (*[]Country, error) {
	var countries []Country
	err := db.Debug().Model(&Country{}).Find(&countries).Error
	if err != nil {
		return nil, err
	}

	return &countries, nil
}
