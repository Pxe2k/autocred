package storage

import "gorm.io/gorm"

type RegistrationAddress struct {
	gorm.Model
	City     string `gorm:"size:100;" json:"city"`     // Населённый пункт
	Address  string `gorm:"size:100;" json:"address"`  // Адрес
	Region   string `gorm:"size:100;" json:"region"`   // Область
	Area     string `gorm:"size:100;" json:"area"`     // Сельский округ
	District string `gorm:"size:100;" json:"district"` // Район
	Street   string `gorm:"size:100;" json:"street"`   // Улица/Микрорайон
	Flat     string `gorm:"size:100;" json:"flat"`     // Номер квартиры
	House    string `gorm:"size:100;" json:"house"`    // Номер дома
	Kato     string `gorm:"size:100;" json:"kato"`     // Код КАТО
	ClientID uint
}
