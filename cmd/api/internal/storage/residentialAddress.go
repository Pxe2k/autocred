package storage

import "gorm.io/gorm"

type ResidentialAddress struct {
	gorm.Model
	Address                   string `gorm:"size:100;" json:"address"` // Адрес
	Address1                  uint   `json:"address1"`
	Address2                  uint   `json:"address2"`
	Address3                  uint   `json:"address3"`
	Address4                  uint   `json:"address4"`
	Address5                  uint   `json:"address5"`
	Address6                  uint   `json:"address6"`
	Kato                      string `gorm:"size:100;" json:"kato"` // Код КАТО
	IndividualClientID        *uint  `json:"individualClientID,omitempty"`
	BeneficialOwnerBusinessID *uint  `json:"beneficialOwnerBusinessID,omitempty"`
}
