package storage

import "gorm.io/gorm"

type PersonalProperty struct {
	gorm.Model
	ResidentialProperty *[]ResidentialProperty `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"residentialProperty,omitempty"`
	VehicleProperty     *[]VehicleProperty     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"vehicleProperty,omitempty"`
	ClientID            uint
}
