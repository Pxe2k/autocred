package storage

import (
	"fmt"
	"gorm.io/gorm"
)

type Client struct {
	gorm.Model
	IsBusiness          bool                 `json:"isBusiness"`                    // Физ/не физ
	TypeOfClient        string               `gorm:"size:100;" json:"typeOfClient"` // Тип клиента
	FirstName           string               `gorm:"size:100;" json:"firstName"`
	MiddleName          string               `gorm:"size:100;" json:"middleName"`
	LastName            *string              `gorm:"size:100;" json:"lastName,omitempty"`
	Sex                 string               `gorm:"size:100;" json:"sex"`       // Пол
	BirthDate           string               `gorm:"size:100;" json:"birthDate"` // ДР
	Country             string               `gorm:"size:100;" json:"country"`
	Residency           string               `gorm:"size:100;" json:"residency"` // Резиденство
	Bin                 string               `gorm:"size:100;" json:"bin"`       // ИИН
	Phone               string               `gorm:"size:100;" json:"phone"`     // Телефон
	SecondPhone         string               `gorm:"size:100;" json:"secondPhone"`
	Email               string               `gorm:"size:100;" json:"email"`     // Email
	Education           string               `gorm:"size:100;" json:"education"` // Образование
	Image               string               `gorm:"size:100;" json:"image"`     // Аватарка
	UserID              uint                 `json:"userId"`
	User                *User                `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user,omitempty"`
	Document            *Document            `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"document,omitempty"`            // Документы
	MaritalStatus       *MaritalStatus       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"maritalStatus,omitempty"`       // Семейное положение
	WorkPlaceInfo       *WorkPlaceInfo       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"workPlaceInfo,omitempty"`       // Информация о месте работы
	RelationWithBank    *RelationWithBank    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"relationWithBank,omitempty"`    // Отношения с банками
	RegistrationAddress *RegistrationAddress `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"registrationAddress,omitempty"` // Адрес прописки
	ResidentialAddress  *ResidentialAddress  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"residentialAddress,omitempty"`  // Адрес проживания
	Contacts            *[]ClientContact     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"contacts,omitempty"`            // Доп. контакты
	BonusInfo           *BonusInfo           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"bonusInfo"`                     // Дополнительная информация
	PersonalProperty    *[]PersonalProperty  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"personalProperty,omitempty"`    // Личное имущество
	CurrentLoans        *[]CurrentLoans      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"currentLoans,omitempty"`        // Действующие кредиты и займы
	BeneficialOwners    *[]BeneficialOwner   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"beneficialOwners,omitempty"`    // Бенефициарные владельцы
	ClientComment       *ClientComment       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"clientComment,omitempty"`
	Pledges             *[]Pledge            `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"pledges,omitempty"` // Залоги
	Documents           *[]Media             `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"documents"`
}

func (c *Client) Save(db *gorm.DB) (*Client, error) {
	err := db.Debug().Create(&c).Error
	if err != nil {
		return &Client{}, err
	}

	return c, nil
}

func (c *Client) All(db *gorm.DB, firstName, middleName, lastName, userID string) (*[]Client, error) {
	var clients []Client

	query := db.Debug().Model(&Client{}).Preload("User")

	if userID != "" {
		query = query.Where("user_id = ?", userID)
	}
	if firstName != "" {
		fmt.Println("test")
		query = query.Where("first_name LIKE ?", "%"+firstName+"%")
	}
	if middleName != "" {
		query = query.Where("middle_name LIKE ?", "%"+middleName+"%")
	}
	if lastName != "" {
		query = query.Where("last_name LIKE ?", "%"+lastName+"%")
	}

	query.Find(&clients)

	err := query.Error
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
		Preload("Documents").
		Preload("Pledges").
		Preload("Pledges.CarModel").
		Preload("Pledges.CarBrand").
		Take(&c).Error
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Client) UpdateAvatar(db *gorm.DB, id uint) (*Client, error) {
	err := db.Debug().Model(&Client{}).Where("id = ?", id).Take(&Client{}).UpdateColumns(
		map[string]interface{}{
			"image": c.Image,
		},
	).Error
	if err != nil {
		return &Client{}, err
	}
	return c, nil
}

func (c *Client) Update(db gorm.DB, client Client) (*Client, error) {
	err := db.Debug().Session(&gorm.Session{FullSaveAssociations: true}).Save(&client).Error
	if err != nil {
		return nil, err
	}

	return &client, nil
}
