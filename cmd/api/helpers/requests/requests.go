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

type EUApplicationRequestData struct {
	Car struct {
		Brand       string `json:"brand"`
		Colour      string `json:"colour"`
		Condition   string `json:"condition"`
		Country     string `json:"country"`
		FuelType    string `json:"fuelType"`
		Insurance   bool   `json:"insurance"`
		Model       string `json:"model"`
		Price       string `json:"price"`
		Type        string `json:"type"`
		Year        string `json:"year"`
		ChannelType string `json:"channelType"`
	} `json:"car"`
	PartyID              string `json:"partyId"`
	JobPhone             string `json:"JobPhone"`
	City                 string `json:"city"`
	CarCityLocation      string `json:"carCityLocation"`
	ClientName           string `json:"clientName"`
	DeliveryAddress      string `json:"deliveryAddress"`
	DownpaySum           string `json:"downpaySum"`
	Duration             string `json:"duration"`
	DownPayment          uint   `json:"downpayment"`
	Email                string `json:"email"`
	Income               bool   `json:"income"`
	IsPartnerOwner       bool   `json:"isPartnerOwner"`
	Phone                string `json:"phone"`
	StoreAddress         string `json:"storeAddress"`
	IncomeMain           int    `json:"incomeMain"`
	MaritalStatus        string `json:"MaritalStatus"`
	ContactPersonName    string `json:"ContactPersonName"`
	ContactPersonContact string `json:"ContactPersonContact"`
	Iin                  string `json:"iin"`
	OrderId              string `json:"orderId"`
	Gsvp                 struct {
		Name          string `json:"name"`
		Ext           string `json:"ext"`
		Base64Content string `json:"base64content"`
	} `json:"gsvp"`
	IdcdBack struct {
		Name          string `json:"name"`
		Ext           string `json:"ext"`
		Base64Content string `json:"base64content"`
	} `json:"idcdBack"`
}

type BankTokenRequestData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type OTPRequestData struct {
	Phone string `json:"phone"`
	Code  string `json:"code"`
}
