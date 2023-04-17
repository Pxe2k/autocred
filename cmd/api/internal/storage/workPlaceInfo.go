package storage

import "gorm.io/gorm"

type WorkPlaceInfo struct {
	gorm.Model

	ClientID uint
}

func (w *WorkPlaceInfo) Update(db gorm.DB, workPlace WorkPlaceInfo) (*WorkPlaceInfo, error) {
	err := db.Debug().Model(&WorkPlaceInfo{}).Where("id = ?", workPlace.ID).Updates(workPlace).Error
	if err != nil {
		return nil, err
	}

	return w, nil
}
