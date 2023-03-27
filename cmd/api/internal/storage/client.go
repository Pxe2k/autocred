package storage

import (
	"gorm.io/gorm"
)

type Client struct {
	gorm.Model
	IsBusiness          bool                 `json:"isBusiness"`                    // Физ/не физ
	FullName            string               `gorm:"size:100;" json:"fullName"`     // ФИО
	TypeOfClient        string               `gorm:"size:100;" json:"typeOfClient"` // Тип клиента
	Sex                 string               `gorm:"size:100;" json:"sex"`          // Пол
	BirthDate           string               `gorm:"size:100;" json:"birthDate"`    // ДР
	Country             string               `gorm:"size:100;" json:"country"`      // Страна
	Residency           string               `gorm:"size:100;" json:"residency"`    // Резиденство
	Education           string               `gorm:"size:100"  json:"education"`    // Образование
	UserID              uint                 `json:"userId"`
	MaritalStatus       *MaritalStatus       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"maritalStatus,omitempty"`       // Семейное положение
	Document            *Document            `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"document,omitempty"`            // Документы
	WorkPlaceInfo       *WorkPlaceInfo       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"workPlaceInfo,omitempty"`       // Информация о месте работы
	RelationWithBank    *RelationWithBank    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"relationWithBank,omitempty"`    // Отношения с банками
	RegistrationAddress *RegistrationAddress `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"registrationAddress,omitempty"` // Адрес прописки
	ResidentialAddress  *ResidentialAddress  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"residentialAddress,omitempty"`  // Адрес проживания
	Contacts            *[]ClientContact     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"contacts,omitempty"`            // Доп. контакты
	BeneficialOwners    *[]BeneficialOwner   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"beneficialOwners,omitempty"`    // Бенефициарные владельцы
	ClientComment       *ClientComment       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"clientComment,omitempty"`
	Pledge              *[]Pledge            `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"pledge,omitempty"` // Залог
}

func (c *Client) Save(db *gorm.DB) (*Client, error) {
	err := db.Debug().Create(&c).Error
	if err != nil {
		return &Client{}, err
	}
	return c, nil
}

func (c *Client) All(db *gorm.DB) (*[]Client, error) {
	var clients []Client

	err := db.Debug().Model(&Client{}).Limit(100).Find(&clients).Error
	if err != nil {
		return nil, err
	}

	return &clients, nil
}

func (c *Client) Get(db *gorm.DB, id uint) (*Client, error) {
	err := db.Debug().Model(&Client{}).Where("id = ?", id).
		Preload("Document").
		Preload("WorkPlaceInfo").
		Preload("MaritalStatus").
		Preload("RelationWithBank").
		Preload("RegistrationAddress").
		Preload("ResidentialAddress").
		Preload("Contacts").
		Preload("BeneficialOwners").
		Preload("ClientComment").
		Take(&c).Error
	if err != nil {
		return nil, err
	}

	return c, nil
}
