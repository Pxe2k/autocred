package storage

import "gorm.io/gorm"

type BeneficialOwnerBusiness struct {
	gorm.Model
	TypeOfClient          string                `gorm:"size:100" json:"typeOfClient"`       // Тип клиента
	FirstName             string                `gorm:"size:100" json:"firstName"`          //
	MiddleName            string                `gorm:"size:100" json:"middleName"`         //
	LastName              *string               `gorm:"size:100" json:"lastName,omitempty"` //
	Country               string                `gorm:"size:100" json:"country"`            //
	Sex                   string                `gorm:"size:100" json:"sex"`                // Пол
	BirthDate             string                `gorm:"size:100" json:"birthDate"`          // ДР
	Email                 string                `gorm:"size:100" json:"email"`              // Email
	Phone                 string                `gorm:"size:100" json:"phone"`              // Телефон
	SecondPhone           string                `gorm:"size:100" json:"secondPhone"`        //
	Education             string                `gorm:"size:100" json:"education"`          //
	MaritalStatus         MaritalStatusBusiness `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"maritalStatus"`
	WorkPlaceInfoBusiness WorkPlaceInfoBusiness `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"workPlaceInfoBusiness"`
	DocumentBusiness      DocumentBusiness      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"documentBusiness"`
	ResidentialAddress    ResidentialAddress    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"residentialAddress"`
	RegistrationAddress   RegistrationAddress   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"registrationAddress"`
	BusinessContact       BusinessContact       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"businessContact"`
	BonusInfoBusiness     BonusInfoBusiness     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"bonusInfoBusiness"`
	CurrentLoanBusiness   CurrentLoanBusiness   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"currentLoanBusiness"`
	Comment               string                `gorm:"size:1000" json:"comment"` // Комментарий
	BusinessClientID      uint
}
