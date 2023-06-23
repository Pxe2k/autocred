package storage

import "gorm.io/gorm"

type BankProcessingDocument struct {
	gorm.Model
	Title         string `json:"title"`
	Image         string `json:"image"`
	IsSign        bool   `json:"isSign"`
	File          string `json:"file"`
	BankID        uint   `json:"bankID"`
	ApplicationID uint   `json:"applicationID"`
}

func (b *BankProcessingDocument) Save(db *gorm.DB) (*BankProcessingDocument, error) {
	err := db.Debug().Create(&b).Error
	if err != nil {
		return &BankProcessingDocument{}, err
	}
	return b, nil
}

func (b *BankProcessingDocument) MultipleSave(db *gorm.DB, media []BankProcessingDocument) ([]BankProcessingDocument, error) {
	err := db.Debug().Create(media).Error
	if err != nil {
		return []BankProcessingDocument{}, err
	}
	return media, nil
}

func (b *BankProcessingDocument) SoftDelete(db *gorm.DB, id uint) (int64, error) {
	err := db.Debug().Model(&BankProcessingDocument{}).Where("id = ?", id).Take(&BankProcessingDocument{}).Delete(&BankProcessingDocument{})
	if err != nil {
		return 0, err.Error
	}

	if err.Error != nil {
		return 0, err.Error
	}

	return err.RowsAffected, nil
}

func (b *BankProcessingDocument) All(db *gorm.DB, uid uint) ([]BankProcessingDocument, error) {
	var media []BankProcessingDocument
	err := db.Debug().Model(&BankProcessingDocument{}).Where("application_id = ?", uid).Limit(100).Find(&media).Error
	if err != nil {
		return nil, err
	}

	return media, nil
}
