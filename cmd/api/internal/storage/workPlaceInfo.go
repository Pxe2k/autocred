package storage

import "gorm.io/gorm"

type WorkPlaceInfo struct {
	gorm.Model
	Address          string `gorm:"size:100;" json:"address"`          // Адрес
	MonthlyIncome    int    `json:"monthlyIncome"`                     // Доход
	OrganizationName string `gorm:"size:100;" json:"organizationName"` // Название организанции
	ActivityType     string `gorm:"size:100" json:"activityType"`      // Тип деятельности
	JobTitle         string `gorm:"size:100;" json:"jobTitle"`         // Должность
	ClientID         uint
}
