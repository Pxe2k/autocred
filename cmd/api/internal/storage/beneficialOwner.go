package storage

import "gorm.io/gorm"

type BeneficialOwner struct {
	gorm.Model

	ClientID uint
}
