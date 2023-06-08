package storage

import (
	"fmt"
	"gorm.io/gorm"
)

type IndividualClient struct {
	gorm.Model
	TypeOfClient        string                       `gorm:"size:100" json:"typeOfClient"` // Тип клиента
	FirstName           string                       `gorm:"size:100" json:"firstName"`
	MiddleName          string                       `gorm:"size:100" json:"middleName"`
	LastName            *string                      `gorm:"size:100" json:"lastName,omitempty"`
	Sex                 string                       `gorm:"size:100" json:"sex"`       // Пол
	BirthDate           string                       `gorm:"size:100" json:"birthDate"` // ДР
	Country             string                       `gorm:"size:100" json:"country"`
	Phone               string                       `gorm:"size:100;unique" json:"phone"` // Телефон
	SecondPhone         string                       `gorm:"size:100" json:"secondPhone"`
	Email               string                       `gorm:"size:100" json:"email"`     // Email
	Education           string                       `gorm:"size:100" json:"education"` // Образование
	Image               string                       `gorm:"size:100" json:"image"`     // Аватарка
	Comment             string                       `gorm:"size:100" json:"comment"`
	UserID              uint                         `json:"userId"`
	User                *User                        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user,omitempty"`
	Document            *Document                    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"document,omitempty"`            // Документы
	MaritalStatus       *MaritalStatus               `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"maritalStatus,omitempty"`       // Семейное положение
	WorkPlaceInfo       *WorkPlaceInfo               `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"workPlaceInfo,omitempty"`       // Информация о месте работы   // Отношения с банками
	CurrentLoans        *[]CurrentLoans              `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"currentLoans,omitempty"`        // Действующие кредиты и займы
	RegistrationAddress *RegistrationAddress         `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"registrationAddress,omitempty"` // Адрес прописки
	ResidentialAddress  *ResidentialAddress          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"residentialAddress,omitempty"`  // Адрес проживания
	Contacts            *[]ClientContact             `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"contacts,omitempty"`            // Доп. контакты
	BonusInfo           *BonusInfo                   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"bonusInfo"`                     // Дополнительная информация
	BeneficialOwners    *[]BeneficialOwnerIndividual `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"beneficialOwners,omitempty"`    // Бенефициарные владельцы
	Pledges             *[]Pledge                    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"pledges,omitempty"`             // Залоги
	Documents           *[]Media                     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"documents"`
}

func (ic *IndividualClient) Save(db *gorm.DB) (*IndividualClient, error) {
	err := db.Debug().Create(&ic).Error
	if err != nil {
		return nil, err
	}

	return ic, nil
}

func (ic *IndividualClient) All(db *gorm.DB, fullName, sex, birthDate, sortUser string, userID uint) (*[]IndividualClient, error) {
	var individualClients []IndividualClient

	query := db.Debug().Model(&IndividualClient{})

	if fullName != "" {
		fmt.Println("fullname", fullName)
		query = db.Raw("SELECT individual_clients.* FROM individual_clients JOIN (SELECT id, concat_ws(' ', last_name, first_name, middle_name) as fullName FROM individual_clients) individual_clients2 ON individual_clients2.fullName ILIKE ? AND individual_clients2.id = individual_clients.id", "%"+fullName+"%")
	}
	if sex != "" {
		query = query.Order("sex " + sex)
	}
	if birthDate != "" {
		query = query.Order("birth_date " + birthDate)
	}
	if sortUser != "" {
		query = query.Order("user_id " + sortUser)
	}

	query.Preload("User").Preload("User.AutoDealer").Where("user_id = ?", userID).Find(&individualClients)

	err := query.Error
	if err != nil {
		return nil, err
	}

	return &individualClients, nil
}

func (ic *IndividualClient) Get(db *gorm.DB, id uint) (*IndividualClient, error) {
	err := db.Debug().Model(&IndividualClient{}).Where("id = ?", id).
		Preload("User").
		Preload("User.AutoDealer").
		Preload("Document").
		Preload("WorkPlaceInfo").
		Preload("MaritalStatus").
		Preload("RegistrationAddress").
		Preload("ResidentialAddress").
		Preload("Contacts").
		Preload("BeneficialOwners").
		Preload("Documents").
		Preload("Pledges").
		Preload("Pledges.CarModel").
		Preload("Pledges.CarBrand").
		Take(&ic).Error
	if err != nil {
		return nil, err
	}

	return ic, nil
}

func (ic *IndividualClient) UpdateAvatar(db *gorm.DB, id uint) (*IndividualClient, error) {
	err := db.Debug().Model(&IndividualClient{}).Where("id = ?", id).Take(&IndividualClient{}).UpdateColumns(
		map[string]interface{}{
			"image": ic.Image,
		},
	).Error
	if err != nil {
		return &IndividualClient{}, err
	}
	return ic, nil
}

func (ic *IndividualClient) UpdateUserID(db *gorm.DB, client IndividualClient) error {
	err := db.Debug().Model(&IndividualClient{}).Where("phone = ?", client.Phone).Update("user_id", client.UserID).Save(&client).Error
	if err != nil {
		return err
	}

	return nil
}
