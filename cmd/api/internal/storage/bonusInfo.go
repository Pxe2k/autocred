package storage

import "gorm.io/gorm"

type BonusInfo struct {
	gorm.Model

	ClientID uint
}
