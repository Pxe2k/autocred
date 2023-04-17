package storage

import "gorm.io/gorm"

type CurrentLoans struct {
	gorm.Model

	ClientID uint
}
