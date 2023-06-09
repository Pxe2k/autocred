package responses

import (
	"autocredit/cmd/api/internal/storage"
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type LoginResponse struct {
	Code string `json:"code"`
}

type UserResponseData struct {
	gorm.Model
	Email          string                `json:"email"`
	FirstName      string                `json:"firstName"`
	MiddleName     string                `json:"middleName"`
	LastName       *string               `json:"lastName,omitempty"`
	IIN            *string               `json:"iin,omitempty"`
	Document       string                `json:"document"`
	DocumentNumber string                `json:"documentNumber"`
	JobTitle       string                `json:"jobTitle"`
	OrderNumber    string                `json:"orderNumber"`
	Phone          string                `json:"phone"`
	WorkPhone      string                `json:"workPhone"`
	AutoDealerID   uint                  `json:"autoDealerID,omitempty"`
	AutoDealer     *storage.AutoDealer   `json:"autodealer,omitempty"`
	RoleID         *uint                 `json:"roleID,omitempty"`
	Role           storage.Role          `json:"role"`
	Applications   []storage.Application `json:"applications"`
}

type IndividualClientResponseData struct {
	ID               uint                                 `json:"ID"`
	TypeOfClient     string                               `gorm:"size:100" json:"typeOfClient"` // Тип клиента
	FirstName        string                               `gorm:"size:100" json:"firstName"`
	MiddleName       string                               `gorm:"size:100" json:"middleName"`
	LastName         string                               `gorm:"size:100" json:"lastName,omitempty"`
	Sex              string                               `gorm:"size:100" json:"sex"`       // Пол
	BirthDate        string                               `gorm:"size:100" json:"birthDate"` // ДР
	Country          string                               `gorm:"size:100" json:"country"`
	Phone            string                               `gorm:"size:100;unique" json:"phone"` // Телефон
	SecondPhone      string                               `gorm:"size:100" json:"secondPhone"`
	Email            string                               `gorm:"size:100" json:"email"`     // Email
	Education        string                               `gorm:"size:100" json:"education"` // Образование
	Status           bool                                 `json:"status"`
	Comment          string                               `gorm:"size:100" json:"comment"`
	Image            string                               `json:"image,omitempty"` // Аватарка
	UserID           uint                                 `json:"userId,omitempty"`
	User             *storage.User                        `json:"user,omitempty"`
	Documents        *[]storage.Media                     `json:"documents,omitempty"`        // Документы
	MaritalStatus    *storage.MaritalStatus               `json:"maritalStatus,omitempty"`    // Семейное положение
	CurrentLoans     *[]storage.CurrentLoans              `json:"currentLoans,omitempty"`     // Действующие кредиты и займы
	Contacts         *[]storage.ClientContact             `json:"contacts,omitempty"`         // Доп. контакты
	BonusInfo        *storage.BonusInfo                   `json:"bonusInfo"`                  // Дополнительная информация
	BeneficialOwners *[]storage.BeneficialOwnerIndividual `json:"beneficialOwners,omitempty"` // Бенефициарные владельцы
	Pledges          *[]storage.Pledge                    `json:"pledges,omitempty"`          // Залоги
	WorkPlaceInfo    struct {
		OrganizationName  string `json:"organizationName"` // Название организанции
		WorkPlace         string `json:"workPlaceType"`    // Тип места работы
		WorkingActivityID uint   `json:"activityTypeID"`   // Тип деятельности
		JobTitle          string `json:"jobTitle"`         // Должность
		MonthlyIncome     int    `json:"monthlyIncome"`    // Доход
		Address           string `json:"address"`          // Адрес
		Experience        string `json:"experience"`       // Стаж работы в организации (мес)
		EmploymentRate    string `json:"employmentRate"`   // Степень занятости
		EmploymentDate    string `json:"employmentDate"`   // Дата трудоустройства
		DateNextSalary    string `json:"dateNextSalary"`   // Дата следующей з/п
		OrganizationPhone string `json:"organizationPhone"`
	} `json:"workPlaceInfo"` // Информация о месте работы   // Отношения с банками
	RegistrationAddress struct {
		Address  string `json:"address"` // Адрес
		Address1 uint   `json:"address1"`
		Address2 uint   `json:"address2"`
		Address3 uint   `json:"address3"`
		Address4 uint   `json:"address4"`
		Address5 uint   `json:"address5"`
		Address6 uint   `json:"address6"`
		Kato     string `json:"kato"` // Код КАТО
	} `json:"registrationAddress"` // Адрес прописки
	ResidentialAddress struct {
		Address  string `json:"address"` // Адрес
		Address1 uint   `json:"address1"`
		Address2 uint   `json:"address2"`
		Address3 uint   `json:"address3"`
		Address4 uint   `json:"address4"`
		Address5 uint   `json:"address5"`
		Address6 uint   `json:"address6"`
		Kato     string `json:"kato"` // Код КАТО
	} `json:"residentialAddress"` // Адрес проживания
	Document struct {
		Type              string `json:"type"`              // Тип документа
		IIN               string `json:"IIN"`               // ИИН
		Number            string `json:"number"`            // Номер
		IssuingAuthority  string `json:"issuingAuthority"`  // Орган выдачи
		PlaceOfBirth      string `json:"placeOfBirth"`      // Место рождения
		DocumentIssueDate string `json:"documentIssueDate"` // Дата выдачи
		DocumentEndDate   string `json:"documentEndDate"`   // Дата истечения
	} `json:"document"`
	Applications *[]storage.Application `json:"applications"`
	CreatedAt    time.Time              `json:"createdAt"`
}

type BusinessClientResponseData struct {
	ID                  uint                                   `json:"ID"`
	TypeOfClient        string                                 `json:"typeOfClient"` // Тип клиента
	Image               string                                 `json:"image"`
	BIN                 string                                 `gorm:"size:100" json:"BIN"`         // БИН
	CompanyName         string                                 `gorm:"size:100" json:"companyName"` // Название организации
	CompanyPhone        string                                 `gorm:"size:100" json:"companyPhone"`
	MonthlyIncome       uint                                   `json:"monthlyIncome"`                    // Ежемесячный доход компании
	CompanyLifespan     string                                 `gorm:"size:100" json:"companyLifespan"`  // Срок существования компании
	KindActivity        string                                 `gorm:"size:100" json:"kindActivity"`     // Вид деятельности
	ActivityType        string                                 `gorm:"size:100" json:"activityType"`     // Тип деятельности
	RegistrationDate    string                                 `gorm:"size:100" json:"registrationDate"` // Тип деятельности
	Status              bool                                   `json:"status"`
	UserID              uint                                   `json:"userID"`
	User                storage.User                           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user"`
	RegistrationAddress *[]storage.RegistrationAddressBusiness `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"registrationAddress,omitempty"` // Адрес регистрации юридического лица
	BeneficialOwner     *[]storage.BeneficialOwnerBusiness     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"beneficialOwner,omitempty"`
	CreatedAt           time.Time                              `json:"createdAt"`
}

type SubmitResponse struct {
	Token  *string `json:"token,omitempty"`
	RoleID *uint32 `json:"roleID,omitempty"`
}

type BCCTokenResponseData struct {
	TokenType   string `json:"token_type"`
	AccessToken string `json:"access_token"`
	ExpiresIn   uint   `json:"expires_in"`
	ConsentedOn uint   `json:"consented_on"`
	Scope       string `json:"bcc.application.private"`
}

type ApplicationResponseData struct {
	Status              bool                `json:"status"`
	Error               error               `json:"error"`
	BCCResponseData     BCCResponseData     `json:"BCCResponseData"`
	EUResponseData      EUResponseData      `json:"EUResponseData"`
	ShinhanResponseData ShinhanResponseData `json:"shinhanResponseData"`
}

type BCCResponseData struct {
	Status    string `json:"status"`
	Message   string `json:"message"`
	RequestId string `json:"requestId"`
}

type EUResponseData struct {
	OrderID string `json:"orderId"`
	Success bool   `json:"success"`
	Status  string `json:"status"`
	Msg     string `json:"msg"`
}

type ShinhanResponseData struct {
	LeadID        int     `json:"leadID"`
	ApplicationID int     `json:"applicationId"`
	ClientID      int     `json:"clientId"`
	CollateralID  []int   `json:"collateral_id"`
	Status        string  `json:"status"`
	CarPrice      int     `json:"car price"`
	Durations     int     `json:"durations"`
	Insurance     bool    `json:"insurance"`
	DownPayment   float64 `json:"downpayment"`
}

type SMSResponse struct {
	Code int `json:"code"`
	Data struct {
		CampaignId string `json:"campaignId"`
		MessageId  string `json:"messageId"`
		Status     int    `json:"status"`
	} `json:"data"`
	Message string `json:"message"`
}

type ShinhanStatusResponseData struct {
	Status string `json:"status"`
}

type EUBankStatusResponseData struct {
	IIN         string `json:"IIN"`
	ClientType  string `json:"ClientType"`
	LastName    string `json:"LastName"`
	FirstName   string `json:"FirstName"`
	MiddleName  string `json:"MiddleName"`
	FIO         string `json:"FIO"`
	Status      string `json:"status"`
	Description string `json:"Description"`
	LoanNumber  string `json:"loanNumber"`
	Product     string `json:"product"`
	Term        int    `json:"term"`
}

type BankDocumentsCreated struct {
	ID    uint   `json:"ID"`
	Title string `json:"title"`
	File  string `json:"file"`
}

type BCCUpdateResponse struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

type ApplicationsResponseData struct {
	AllApplications      int                   `json:"allApplications"`
	SuccessApplications  int                   `json:"successApplications"`
	DeclinedApplications int                   `json:"declinedApplications"`
	Applications         []storage.Application `json:"applications"`
}

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

func ERROR(w http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		JSON(w, statusCode, struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
		return
	}
	JSON(w, http.StatusBadRequest, nil)
}
