package storage

import "gorm.io/gorm"

type BeneficialOwner struct {
	gorm.Model
	FullName           string `gorm:"size:100;" json:"fullName"`
	Bin                string `gorm:"size:100;" json:"bin"`
	Sex                string `gorm:"size:100;" json:"sex"`
	BirthDate          string `gorm:"size:100;" json:"birthDate"`
	TypeOfOwnership    string `gorm:"size:100;" json:"typeOfOwnership"`
	IndividualClientID uint
}
