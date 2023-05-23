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

type UnauthorizedUserResponseData struct {
	ID                 uint                        `json:"ID"`
	TypeOfClient       string                      `json:"typeOfClient"` // Тип клиента
	FirstName          string                      `json:"firstName"`
	MiddleName         string                      `json:"middleName"`
	LastName           *string                     `json:"lastName,omitempty"`
	BirthDate          string                      `json:"birthDate"`
	Phone              string                      `json:"phone"` // Телефон
	Document           *storage.Document           `json:"document,omitempty"`
	ResidentialAddress *storage.ResidentialAddress `json:"residentialAddress,omitempty"` // Адрес проживания
	CreatedAt          time.Time                   `json:"createdAt"`
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
