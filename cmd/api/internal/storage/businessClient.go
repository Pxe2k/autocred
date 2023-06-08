package storage

import (
	"gorm.io/gorm"
)

type BusinessClient struct {
	gorm.Model
	TypeOfClient        string                      `gorm:"size:100" json:"typeOfClient"` // Тип клиента
	Image               string                      `gorm:"size:100" json:"image"`
	BIN                 string                      `gorm:"size:100;unique" json:"BIN"`  // БИН
	CompanyName         string                      `gorm:"size:100" json:"companyName"` // Название организации
	CompanyPhone        string                      `gorm:"size:100" json:"companyPhone"`
	MonthlyIncome       uint                        `json:"monthlyIncome"`                    // Ежемесячный доход компании
	CompanyLifespan     string                      `gorm:"size:100" json:"companyLifespan"`  // Срок существования компании
	KindActivity        string                      `gorm:"size:100" json:"kindActivity"`     // Вид деятельности
	ActivityType        string                      `gorm:"size:100" json:"activityType"`     // Тип деятельности
	RegistrationDate    string                      `gorm:"size:100" json:"registrationDate"` // Тип деятельности
	UserID              uint                        `json:"userID"`
	User                User                        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user"`
	RegistrationAddress RegistrationAddressBusiness `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"registrationAddress"` // Адрес регистрации юридического лица
	BeneficialOwner     BeneficialOwnerBusiness     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"beneficialOwner"`
	Pledges             *[]Pledge                   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"pledges,omitempty"` // Залоги
	Documents           *[]Media                    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"documents"`
}

func (bc *BusinessClient) Save(db *gorm.DB) (*BusinessClient, error) {
	err := db.Debug().Create(&bc).Error
	if err != nil {
		return nil, err
	}

	return bc, nil
}

func (bc *BusinessClient) All(db *gorm.DB, fullName, sex, birthDate, userID, sortUser string) (*[]BusinessClient, error) {
	var individualClients []BusinessClient

	query := db.Debug().Model(&BusinessClient{})

	if fullName != "" {
		query = db.Raw("SELECT clients.* FROM clients JOIN (SELECT id, concat_ws(' ', last_name, first_name, middle_name) as fullName FROM clients) clients2 ON clients2.fullName ILIKE ? AND clients2.id = clients.id", "%"+fullName+"%")
	}
	if userID != "" {
		query = query.Where("user_id = ?", userID)
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

func (bc *BusinessClient) Get(db *gorm.DB, id uint) (*BusinessClient, error) {
	err := db.Debug().Model(&BusinessClient{}).Where("id = ?", id).
		Take(&bc).Error
	if err != nil {
		return nil, err
	}

	return bc, nil
}

func (bc *BusinessClient) UpdateAvatar(db *gorm.DB, id uint) (*BusinessClient, error) {
	err := db.Debug().Model(&BusinessClient{}).Where("id = ?", id).Take(&BusinessClient{}).UpdateColumns(
		map[string]interface{}{
			"image": bc.Image,
		},
	).Error
	if err != nil {
		return &BusinessClient{}, err
	}
	return bc, nil
}

func (bc *BusinessClient) UpdateUserID(db *gorm.DB, client BusinessClient) error {
	err := db.Debug().Model(&BusinessClient{}).Where("phone = ?", client.BIN).Update("user_id", client.UserID).Save(&client).Error
	if err != nil {
		return err
	}

	return nil
}
