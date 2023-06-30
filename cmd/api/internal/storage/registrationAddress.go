package storage

import "gorm.io/gorm"

type RegistrationAddress struct {
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

func (ra RegistrationAddress) Update(db *gorm.DB, address *RegistrationAddress, clientID uint) error {
	err := db.Debug().Model(&WorkPlaceInfo{}).Where("individual_client_id = ?", clientID).Updates(address).Error
	if err != nil {
		return err
	}

	return nil
}
