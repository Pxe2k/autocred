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

type BCCApplicationRequestData struct {
	PartnerID           string `json:"partnerId"`
	PartnerName         string `json:"partnerName"`
	PartnerBin          string `json:"partnerBin"`
	DealerID            string `json:"dealer_id"`
	PartnerCity         string `json:"partnerCity"`
	CostObject          int    `json:"costObject"`
	DownPaymentAmt      int    `json:"downPaymentAmt"`
	LoanAmt             int    `json:"loanAmt"`
	LoanDuration        int    `json:"loanDuration"`
	SimpleFinAnalysis   int    `json:"simpleFinAnalysis"`
	Brand               string `json:"brand"`
	Model               string `json:"model"`
	IssueYear           int    `json:"issueYear"`
	Iin                 string `json:"iin"`
	IDocType            string `json:"iDocType"`
	ProductCode         string `json:"productCode"`
	MobilePhoneNo       string `json:"mobilePhoneNo"`
	WorkName            string `json:"workName"`
	WorkAddress         string `json:"workAddress"`
	WorkStatus          string `json:"workStatus"`
	OrganizationPhoneNo string `json:"organizationPhoneNo"`
	BasicIncome         int    `json:"basicIncome"`
	AdditionalIncome    int    `json:"additionalIncome"`
	UserCode            string `json:"user_code"`
	StatementType       string `json:"statementType"`
	ContactPerson       []struct {
		FullName string `json:"fullName"`
		PhoneNo  string `json:"phoneNo"`
	} `json:"contactPerson"`
	Document struct {
		File      string `json:"file"`
		Extension string `json:"extension"`
		Code      string `json:"code"`
	} `json:"document"`
	MsgID string `json:"msg_id"`
}

type IntegrationRequestData struct {
	ClientID uint `json:"ClientID"`
}
