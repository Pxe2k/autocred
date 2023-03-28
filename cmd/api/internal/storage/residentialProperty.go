package storage

import "gorm.io/gorm"

type ResidentialProperty struct {
	gorm.Model
	Type               string `gorm:"size:100;" json:"type"`             // Вид имущества
	City               string `gorm:"size:100;" json:"city"`             // Город
	Square             string `gorm:"size:100;" json:"square"`           // Площадь
	ConstructionYear   string `gorm:"size:100;" json:"constructionYear"` // Год постройки
	Price              int    `json:"price"`                             // Стоимость
	PurchaseMethod     string `json:"purchaseMethod"`                    // Способ приобретения
	Description        string `json:"description"`                       // Описание
	Document           string `gorm:"size:100;" json:"document"`         // Документ
	PersonalPropertyID uint
}
