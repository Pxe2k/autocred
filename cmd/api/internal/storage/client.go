package storage

import (
	"gorm.io/gorm"
)

type Client struct {
	gorm.Model
	IsBusiness                         bool                `json:"isBusiness"`                    // Физ/не физ
	TypeOfClient                       string              `gorm:"size:100;" json:"typeOfClient"` // Тип клиента
	FirstName                          string              `gorm:"size:100;" json:"firstName"`
	MiddleName                         string              `gorm:"size:100;" json:"middleName"`
	LastName                           *string             `gorm:"size:100;" json:"lastName,omitempty"`
	Sex                                string              `gorm:"size:100;" json:"sex"`       // Пол
	BirthDate                          string              `gorm:"size:100;" json:"birthDate"` // ДР
	Country                            string              `gorm:"size:100;" json:"country"`
	Residency                          string              `gorm:"size:100;" json:"residency"` // Резиденство
	Bin                                string              `gorm:"size:100;" json:"bin"`       // ИИН
	Phone                              string              `gorm:"size:100;" json:"phone"`     // Телефон
	SecondPhone                        string              `gorm:"size:100;" json:"secondPhone"`
	Email                              string              `gorm:"size:100;" json:"email"`              // Email
	Status                             string              `gorm:"size:100;" json:"status"`             // Семейное положение
	FamilyPartnerName                  string              `gorm:"size:100;" json:"familyPartnerName"`  // ФИО партнера
	FamilyPartnerPhone                 string              `gorm:"size:100;" json:"familyPartnerPhone"` // Телефон мужа/жены
	MinorChildren                      string              `gorm:"size:100;" json:"minorChildren"`      // Кол-во несовершеннолетних детей
	Education                          string              `gorm:"size:100;" json:"education"`          // Образование
	DocumentType                       string              `gorm:"size:100;" json:"documentType"`       // Тип документа
	IIN                                string              `gorm:"size:100;" json:"IIN"`                // ИИН
	Number                             string              `gorm:"size:100;" json:"number"`             // Номер
	DocumentIssueDate                  string              `gorm:"size:100;" json:"documentIssueDate"`  // Дата выдачи
	DocumentEndDate                    string              `gorm:"size:100;" json:"documentEndDate"`    // Дата истечения
	IssuingAuthority                   string              `gorm:"size:100;" json:"issuingAuthority"`   // Орган выдачи
	PlaceOfBirth                       string              `gorm:"size:100;" json:"placeOfBirth"`       // Место рождения
	OrganizationName                   string              `gorm:"size:100;" json:"organizationName"`   // Название организанции
	WorkPlace                          string              `gorm:"size:100;" json:"workPlaceType"`      // Тип места работы
	WorkingActivityID                  uint                `json:"activityTypeID"`                      // Тип деятельности
	JobTitleID                         uint                `json:"jobTitleID"`                          // Должность
	MonthlyIncome                      int                 `json:"monthlyIncome"`                       // Доход
	WorkPlaceAddress                   string              `gorm:"size:100;" json:"workPlaceAddress"`   // Адрес
	Experience                         string              `gorm:"size:100;" json:"experience"`         // Стаж работы в организации (мес)
	EmploymentRate                     string              `gorm:"size:100;" json:"employmentRate"`     // Степень занятости
	EmploymentDate                     string              `gorm:"size:100;" json:"employmentDate"`     // Дата трудоустройства
	DateNextSalary                     string              `gorm:"size:100;" json:"dateNextSalary"`     // Дата следующей з/п
	OrganizationPhone                  string              `gorm:"size:100;" json:"organizationPhone"`
	WorkingActivity                    *WorkingActivity    `json:"workingActivity,omitempty"`
	JobTitle                           *JobTitle           `json:"jobTitle,omitempty"`
	AmountOfCredits                    string              `gorm:"size:100;" json:"amountOfCredits"`     // Количество текущих кредитов
	AmountOfPayments                   int                 `json:"amountOfPayments"`                     // Сумма ежемесячных платежей
	OverdraftAmount                    int                 `json:"overdraftAmount"`                      // Сумма ОД по текущим кредитам
	Delay                              string              `gorm:"size:100;" json:"delay"`               // Текущая просрочка свыше 7 дней
	RegistrationAddress                string              `gorm:"size:100;" json:"registrationAddress"` // Адрес
	RegistrationAddress1               uint                `json:"registrationAddress1"`
	RegistrationAddress2               uint                `json:"registrationAddress2"`
	RegistrationAddress3               uint                `json:"registrationAddress3"`
	RegistrationAddress4               uint                `json:"registrationAddress4"`
	RegistrationAddress5               uint                `json:"registrationAddress5"`
	RegistrationAddress6               uint                `json:"registrationAddress6"`
	RegistrationKato                   string              `gorm:"registrationAddress:100;" json:"registrationKato"` // Код КАТО
	ResidentialAddress                 string              `gorm:"size:100;" json:"residentialAddress"`              // Адрес
	ResidentialAddress1                uint                `json:"residentialAddress1"`
	ResidentialAddress2                uint                `json:"residentialAddress2"`
	ResidentialAddress3                uint                `json:"residentialAddress3"`
	ResidentialAddress4                uint                `json:"residentialAddress4"`
	ResidentialAddress5                uint                `json:"residentialAddress5"`
	ResidentialAddress6                uint                `json:"residentialAddress6"`
	ResidentialKato                    string              `gorm:"size:100;" json:"residentialKato"`     // Код КАТО
	AmountIncome                       int                 `json:"amountIncome"`                         // Сумма доходов
	AmountExpenses                     int                 `json:"amountExpenses"`                       // Сумма расходов
	NetIncome                          int                 `json:"netIncome"`                            // Чистый доход
	RecentBankContacts                 bool                `json:"recentBankContacts"`                   // Обращались ли вы в банки за последние 3 мес
	BankTitle                          *string             `gorm:"size:100;" json:"bankTitle,omitempty"` // Название банков
	Returnee                           bool                `json:"returnee"`                             // Статус Оралман
	ThirdPersons                       bool                `json:"thirdPerson"`                          // Проведение операции клиентом под руководством третьего лица и/или лиц присутствующих при операции
	NonProfitOrganizations             bool                `json:"nonProfitOrganization"`                // Некомерческие и благотворительные организации, религиозные объедения
	BadHabits                          bool                `json:"badHabits"`                            // По внешнему виду (признаки лица без определенного места жительства, признаки наркомании и (или) алкоголизма) лица осуществляющего
	IPDLStatus                         bool                `json:"IPDLStatus"`                           // Флаг: Статус ИПДЛ
	CriminalRecord                     bool                `json:"criminalRecord"`                       // Судимость
	Lawsuits                           bool                `json:"lawsuits"`                             // Судебные иски
	OverdueDebts                       bool                `json:"overdueDebts"`                         // Имеются ли просроченные долги
	IncapacitationDecisions            bool                `json:"incapacitationDecisions"`              // Существуют ли или существовали в прошлом решения суда об ограничении дееспособности клиента или об установлении над клиентом опекунства
	UnfulfilledJudgment                bool                `json:"unfulfilledJudgment"`                  // Существует невыполненное судебное решение
	RegisteredPsychiatrist             bool                `json:"registeredPsychiatrist"`               // Состоит на учете у психиатра или нарколога
	RestrictionsRightCreditTransaction bool                `json:"restrictionsRightCreditTransaction"`   // Установлены ли какие-либо ограничения права заключать кредитную сделку (в т.ч. брачным договором)
	Pregnancy                          bool                `json:"pregnancy"`                            // Беременность
	TaxArrears                         bool                `json:"taxArrears"`                           // Имеется ли задолженность по налогам и сборам
	Document                           string              `gorm:"size:100;" json:"document"`            // Документ
	CurrentLoanBankTitle               string              `gorm:"size:100;" json:"currentLoanBankTitle"`
	AgreementNumber                    string              `gorm:"size:100;" json:"agreementNumber"`
	Sum                                int                 `json:"sum"`
	BalanceOwed                        int                 `json:"balanceOwed"`
	MonthlyPayment                     int                 `json:"monthlyPayment"`
	InterestRate                       string              `json:"interestRate"`
	PresenceOfDelays                   bool                `json:"presenceOfDelays"`
	FrequencyOfPayments                string              `gorm:"size:100;" json:"frequencyOfPayments"`
	DateLastPayment                    string              `gorm:"size:100;" json:"dateLastPayment"`
	WorkInfo                           string              `json:"workInfo"`
	Comment                            string              `json:"comment"`
	RiskStatus                         string              `json:"riskStatus"`
	Image                              string              `gorm:"size:100;" json:"image"` // Аватарка
	UserID                             uint                `json:"userId"`
	User                               *User               `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user,omitempty"`
	Contacts                           *[]ClientContact    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"contacts,omitempty"`         // Доп. контакты
	PersonalProperties                 *[]PersonalProperty `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"personalProperty,omitempty"` // Личное имущество
	BeneficialOwners                   *[]BeneficialOwner  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"beneficialOwners,omitempty"` // Бенефициарные владельцы
	Pledges                            *[]Pledge           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"pledges,omitempty"`          // Залоги
	Documents                          *[]Media            `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"documents"`
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

	err := db.Debug().Model(&Client{}).Preload("User").Limit(100).Find(&clients).Error
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
	err := db.Debug().Model(&Client{}).Where("id = ?", client.ID).Session(&gorm.Session{FullSaveAssociations: true}).Updates(client).Error
	if err != nil {
		return nil, err
	}

	return c, nil
}
