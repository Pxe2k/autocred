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
		Condition string `json:"condition"`
		Brand     string `json:"brand"`
		Model     string `json:"model"`
		Insurance bool   `json:"insurance"`
		Price     uint   `json:"price"`
		Year      uint   `json:"year"`
	} `json:"car"`
	City                 string `json:"city"`
	Income               bool   `json:"income"`
	PartyID              string `json:"partyId"`
	DownPayment          uint   `json:"downpayment"`
	Duration             uint   `json:"duration"`
	OrderID              string `json:"orderId"`
	Iin                  string `json:"iin"`
	Phone                string `json:"phone"`
	JobPhone             string `json:"JobPhone"`
	IncomeMain           int    `json:"incomeMain"`
	MaritalStatus        string `json:"MaritalStatus"`
	ContactPersonName    string `json:"ContactPersonName"`
	ContactPersonContact string `json:"ContactPersonContact"`
	IncomeAddConfirmed   string `json:"IncomeAddConfirmed"`
	Gsvp                 struct {
		Name          string `json:"name"`
		Extension     string `json:"ext"`
		Base64Content string `json:"base64content"`
	} `json:"gsvp"`
	Idcd struct {
		Name          string `json:"name"`
		Extension     string `json:"ext"`
		Base64Content string `json:"base64content"`
	} `json:"idcd"`
	Photo struct {
		Name          string `json:"name"`
		Extension     string `json:"ext"`
		Base64Content string `json:"base64content"`
	} `json:"phto"`
}

type ShinhanApplicationRequestData struct {
	CalculationType string `json:"calculationType"`
	Car             struct {
		Brand     string `json:"brand"`
		Colour    string `json:"colour"`
		Condition string `json:"condition"`
		Country   string `json:"country"`
		FuelType  string `json:"fuelType"`
		Model     string `json:"model"`
		Price     string `json:"price"`
		Type      string `json:"type"`
		Year      string `json:"year"`
	} `json:"car"`
	Cas      bool   `json:"cas"`
	City     string `json:"city"`
	Customer struct {
		ActualAddress struct {
			District   string `json:"district"`
			Flat       string `json:"flat"`
			House      string `json:"house"`
			Region     string `json:"region"`
			Settlement string `json:"settlement"`
			Street     string `json:"street"`
		} `json:"actualAddress"`
		BirthDate             string `json:"birthDate"`
		BirthPlace            string `json:"birthPlace"`
		ContactPersonFullName string `json:"contactPersonFullName"`
		ContactPersonPhone    string `json:"contactPersonPhone"`
		Document              struct {
			CountryOfResidence string `json:"countryOfResidence"`
			ExpirationDate     string `json:"expirationDate"`
			IssuedDate         string `json:"issuedDate"`
			Issuer             string `json:"issuer"`
			Number             string `json:"number"`
			PhotoBack          string `json:"photoBack"`
			PhotoFront         string `json:"photoFront"`
			Type               string `json:"type"`
		} `json:"document"`
		EmployerAddress struct {
			District   string `json:"district"`
			Flat       string `json:"flat"`
			House      string `json:"house"`
			Region     string `json:"region"`
			Settlement string `json:"settlement"`
			Street     string `json:"street"`
		} `json:"employerAddress"`
		EmployerName        string `json:"employerName"`
		EmployerPhone       string `json:"employerPhone"`
		EmploymentType      string `json:"employmentType"`
		Firstname           string `json:"firstname"`
		Gender              string `json:"gender"`
		Iin                 string `json:"iin"`
		Income              bool   `json:"income"`
		Lastname            string `json:"lastname"`
		MaritalStatus       string `json:"maritalStatus"`
		MobilePhone         string `json:"mobilePhone"`
		NumberOfDependents  string `json:"numberOfDependents"`
		OfficialIncome      string `json:"officialIncome"`
		Patronymic          string `json:"patronymic"`
		Photo               string `json:"photo"`
		RegistrationAddress struct {
			District   string `json:"district"`
			Flat       string `json:"flat"`
			House      string `json:"house"`
			Region     string `json:"region"`
			Settlement string `json:"settlement"`
			Street     string `json:"street"`
		} `json:"registrationAddress"`
		ResidencyStatus string `json:"residencyStatus"`
	} `json:"customer"`
	Discount       bool   `json:"discount"`
	Downpayment    string `json:"downpayment"`
	Duration       string `json:"duration"`
	GosProgram     bool   `json:"gosProgram"`
	Grace          bool   `json:"grace"`
	InstalmentDate string `json:"instalmentDate"`
	Insurance      bool   `json:"insurance"`
	PartnerId      string `json:"partnerId"`
	Verification   struct {
		Code string `json:"code"`
		Date string `json:"date"`
	} `json:"verification"`
}

type BankTokenRequestData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type OTPRequestData struct {
	Phone string `json:"phone"`
	Code  string `json:"code"`
}

type BCCTemplateData struct {
	FIO         string
	Phone       string
	OTP         string
	CurrentDate string
	Place       string // Место приема заявок
}

type ContactPerson struct {
	FullName string `json:"fullName"`
	PhoneNo  string `json:"phoneNo"`
}

type GenerateDocumentRequestData struct {
	Banks []struct {
		Title string `json:"title"`
	} `json:"banks"`
}
