package storage

import "gorm.io/gorm"

type ResidentialAddress struct {
	gorm.Model

	ClientID uint
}
