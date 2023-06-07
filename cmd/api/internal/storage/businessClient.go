package storage

import "gorm.io/gorm"

type BusinessClient struct {
	gorm.Model
	TypeOfClient        string                      `gorm:"size:100" json:"typeOfClient"` // Тип клиента
	BIN                 string                      `gorm:"size:100" json:"BIN"`          // БИН
	CompanyName         string                      `gorm:"size:100" json:"companyName"`  // Название организации
	CompanyPhone        string                      `gorm:"size:100" json:"companyPhone"`
	MonthlyIncome       uint                        `json:"monthlyIncome"`                                                             // Ежемесячный доход компании
	CompanyLifespan     string                      `gorm:"size:100" json:"companyLifespan"`                                           // Срок существования компании
	KindActivity        string                      `gorm:"size:100" json:"kindActivity"`                                              // Вид деятельности
	ActivityType        string                      `gorm:"size:100" json:"activityType"`                                              // Тип деятельности
	RegistrationDate    string                      `gorm:"size:100" json:"registrationDate"`                                          // Тип деятельности
	RegistrationAddress RegistrationAddressBusiness `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"registrationAddress"` // Адрес регистрации юридического лица
	BeneficialOwner     BeneficialOwnerBusiness     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"beneficialOwner"`
}
