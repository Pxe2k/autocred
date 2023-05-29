package responses

import (
	"autocredit/cmd/api/internal/storage"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type LoginResponse struct {
	Code string `json:"code"`
}

type UserResponseData struct {
	ID                  uint                         `json:"ID"`
	IsBusiness          bool                         `json:"isBusiness"`   // Физ/не физ
	TypeOfClient        string                       `json:"typeOfClient"` // Тип клиента
	FirstName           string                       `json:"firstName"`
	MiddleName          string                       `json:"middleName"`
	LastName            string                       `json:"lastName,omitempty"`
	Sex                 string                       `json:"sex,omitempty"` // Пол
	BirthDate           string                       `json:"birthDate"`     // ДР
	Country             string                       `json:"country,omitempty"`
	Residency           bool                         `json:"residency,omitempty"` // Резиденство
	Bin                 string                       `json:"bin,omitempty"`       // ИИН
	Phone               string                       `json:"phone"`               // Телефон
	SecondPhone         string                       `json:"secondPhone,omitempty"`
	Email               string                       `json:"email,omitempty"`     // Email
	Education           string                       `json:"education,omitempty"` // Образование
	Status              bool                         `json:"status"`
	Comment             string                       `json:"comment,omitempty"`
	Image               string                       `json:"image,omitempty"` // Аватарка
	UserID              uint                         `json:"userId,omitempty"`
	User                *storage.User                `json:"user,omitempty"`
	Document            *storage.Document            `json:"document,omitempty"`            // Документы
	MaritalStatus       *storage.MaritalStatus       `json:"maritalStatus,omitempty"`       // Семейное положение
	WorkPlaceInfo       *storage.WorkPlaceInfo       `json:"workPlaceInfo,omitempty"`       // Информация о месте работы   // Отношения с банками
	RegistrationAddress *storage.RegistrationAddress `json:"registrationAddress,omitempty"` // Адрес прописки
	ResidentialAddress  *storage.ResidentialAddress  `json:"residentialAddress,omitempty"`  // Адрес проживания
	Contacts            *[]storage.ClientContact     `json:"contacts,omitempty"`            // Доп. контакты
	BonusInfo           *storage.BonusInfo           `json:"bonusInfo,omitempty"`           // Дополнительная информация
	PersonalProperty    *[]storage.PersonalProperty  `json:"personalProperty,omitempty"`    // Личное имущество
	CurrentLoans        *[]storage.CurrentLoans      `json:"currentLoans,omitempty"`        // Действующие кредиты и займы
	BeneficialOwners    *[]storage.BeneficialOwner   `json:"beneficialOwners,omitempty"`    // Бенефициарные владельцы
	Pledges             *[]storage.Pledge            `json:"pledges,omitempty"`             // Залоги
	Documents           *[]storage.Media             `json:"documents,omitempty"`
	CreatedAt           time.Time                    `json:"createdAt"`
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
