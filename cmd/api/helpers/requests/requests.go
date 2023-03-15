package requests

type UserRequestData struct {
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Creditor bool   `json:"creditor"`
	Bank     string `json:"bank"`
	Password string `json:"password"`
	BankID   *uint  `json:"bankID"`
}

type SubmitCode struct {
	Phone string `json:"phone"`
	Code  string `json:"code"`
}
