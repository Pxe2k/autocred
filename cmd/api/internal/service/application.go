package service

import (
	"autocredit/cmd/api/helpers"
	"autocredit/cmd/api/helpers/requests"
	"autocredit/cmd/api/helpers/responses"
	"autocredit/cmd/api/internal/storage"
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

func CreateApplicationService(db *gorm.DB, body []byte, uid uint) (*storage.Application, error) {
	application := storage.Application{}
	err := json.Unmarshal(body, &application)
	if err != nil {
		return nil, err
	}

	individualClient := storage.IndividualClient{}

	individualClientGotten, err := individualClient.Get(db, 10)
	if err != nil {
		return nil, err
	}

	for _, bankApplication := range application.BankApplications {
		if bankApplication.Bank == "BCC" {
			bccResponseData, err := createBCCApplication(individualClientGotten, application, bankApplication)
			if err != nil {
				fmt.Println("error: ", err)
			}
			fmt.Println("bcc resp data: ", bccResponseData)
		}
		if bankApplication.Bank == "EU" {
			fmt.Println("Kaspi")
		}
		if bankApplication.Bank == "Kaspi" {
			fmt.Println("Kaspi")
		}
	}

	application.UserID = uid

	return &application, nil
}

func createBCCApplication(individualClient *storage.IndividualClient, application storage.Application, bankApplication storage.BankApplication) (responses.BCCResponseData, error) {
	authToken, err := getBCCToken()
	if err != nil {
		return responses.BCCResponseData{}, err
	}

	requestData, err := fillingBCCRequestData(individualClient, application, bankApplication)
	if err != nil {
		return responses.BCCResponseData{}, err
	}

	requestBody, err := json.Marshal(requestData)
	if err != nil {
		return responses.BCCResponseData{}, err
	}

	url := os.Getenv("BCC_APPLICATION")
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		return responses.BCCResponseData{}, err
	}

	// Add header parameters to the request
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+authToken)
	// Add more headers as needed

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return responses.BCCResponseData{}, err
	}

	defer resp.Body.Close()

	serverResponse, err := io.ReadAll(resp.Body)
	if err != nil {
		return responses.BCCResponseData{}, err
	}

	var responseData responses.BCCResponseData

	err = json.Unmarshal(serverResponse, &responseData)
	if err != nil {
		return responses.BCCResponseData{}, err
	}

	return responseData, nil
}

func fillingBCCRequestData(client *storage.IndividualClient, applicationData storage.Application, bankApplicationData storage.BankApplication) (requests.BCCApplicationRequestData, error) {
	var requestData requests.BCCApplicationRequestData

	issueYear, err := strconv.Atoi(applicationData.YearIssue)
	if err != nil {
		return requests.BCCApplicationRequestData{}, err
	}

	requestData.PartnerID = "185124"
	requestData.PartnerName = "TOO BRROKER"
	requestData.PartnerBin = "170540017799"
	requestData.DealerID = "4011"
	requestData.PartnerCity = "Алматы"
	requestData.CostObject = applicationData.CarPrice
	requestData.DownPaymentAmt = applicationData.InitFee
	requestData.LoanAmt = bankApplicationData.LoanAmount
	requestData.LoanDuration = bankApplicationData.TrenchesNumber
	requestData.SimpleFinAnalysis = 0
	requestData.Brand = applicationData.CarBrand
	requestData.Model = applicationData.CarModel
	requestData.IssueYear = issueYear
	requestData.Iin = client.Document.IIN
	requestData.IDocType = "УЛ"
	if applicationData.Condition == false {
		requestData.ProductCode = "0.201.1.0514"
	} else {
		requestData.ProductCode = "0.201.1.0513"
	}
	requestData.MobilePhoneNo = client.Phone
	requestData.WorkName = client.WorkPlaceInfo.OrganizationName
	requestData.WorkAddress = client.WorkPlaceInfo.Address
	switch client.WorkPlaceInfo.WorkingActivityID {
	case 1:
		requestData.WorkStatus = "Пенсионер"
	case 2:
		requestData.WorkStatus = "Работающий пенсионер"
	case 3:
		requestData.WorkStatus = "Военнослужащий"
	default:
		requestData.WorkStatus = "Обычный клиент"
	}
	requestData.OrganizationPhoneNo = client.WorkPlaceInfo.OrganizationPhone
	requestData.BasicIncome = client.BonusInfo.AmountIncome
	requestData.AdditionalIncome = 0
	requestData.UserCode = client.MiddleName + " " + client.FirstName + " " + client.LastName
	for _, contact := range *client.Contacts {
		requestData.ContactPerson = append(requestData.ContactPerson, requests.ContactPerson{
			FullName: contact.FullName,
			PhoneNo:  contact.Phone,
		})
	}
	requestData.Document.File, err = encodeFileToBase64("storage/bcc-data-processing_" + helpers.CurrentDateString() + ".pdf")
	requestData.Document.Extension = "pdf"
	requestData.Document.Code = "SOG"

	return requestData, nil
}

func CreateEUApplication(body []byte) (responses.EUResponseData, error) {
	var requestData requests.EUApplicationRequestData
	err := json.Unmarshal(body, &requestData)
	if err != nil {
		return responses.EUResponseData{}, err
	}

	requestData.Gsvp.Base64Content, err = encodeFileToBase64("templates/resultMedia/outputPDF/autocredit.pdf")
	if err != nil {
		return responses.EUResponseData{}, err
	}
	requestData.Idcd.Base64Content, err = encodeFileToBase64("eu-bank.jpg")
	if err != nil {
		return responses.EUResponseData{}, err
	}
	requestData.Photo.Base64Content, err = encodeFileToBase64("eu-bank.jpg")
	if err != nil {
		return responses.EUResponseData{}, err
	}
	requestData.OrderID = helpers.RandBankApplicationID(16)
	requestData.Gsvp.Name = "GSPV"
	requestData.Gsvp.Extension = "pdf"
	requestData.Idcd.Name = "IDCD"
	requestData.Idcd.Extension = "jpg"
	requestData.Photo.Name = "PHTO"
	requestData.Photo.Extension = "jpg"

	requestBody, err := json.Marshal(requestData)
	if err != nil {
		return responses.EUResponseData{}, err
	}

	url := os.Getenv("EU_APPLICATION")
	fmt.Println("url: ", url)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		return responses.EUResponseData{}, err
	}

	// Add header parameters to the request
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-eub-token", os.Getenv("EU_TOKEN"))

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return responses.EUResponseData{}, err
	}

	fmt.Println("StatusCode ", resp.StatusCode)

	defer resp.Body.Close()

	serverResponse, err := io.ReadAll(resp.Body)
	if err != nil {
		return responses.EUResponseData{}, err
	}

	var responseData responses.EUResponseData

	err = json.Unmarshal(serverResponse, &responseData)
	if err != nil {
		return responses.EUResponseData{}, err
	}

	return responseData, nil
}

func CreateShinhanApplication(body []byte) (responses.ShinhanResponseData, error) {
	var requestData requests.ShinhanApplicationRequestData
	err := json.Unmarshal(body, &requestData)
	if err != nil {
		return responses.ShinhanResponseData{}, err
	}

	requestData.Customer.Document.PhotoBack, err = encodeFileToBase64("templates/resultMedia/outputPDF/autocredit.pdf")
	if err != nil {
		return responses.ShinhanResponseData{}, err
	}
	requestData.Customer.Document.PhotoFront, err = encodeFileToBase64("templates/resultMedia/outputPDF/autocredit.pdf")
	if err != nil {
		return responses.ShinhanResponseData{}, err
	}
	requestData.Customer.Photo, err = encodeFileToBase64("eu-bank.jpg")
	if err != nil {
		return responses.ShinhanResponseData{}, err
	}
	requestData.CalculationType = "A"
	requestData.Cas = false
	requestData.Discount = false

	requestBody, err := json.Marshal(requestData)
	if err != nil {
		return responses.ShinhanResponseData{}, err
	}

	url := os.Getenv("SHINHAN_APPLICATION")
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		return responses.ShinhanResponseData{}, err
	}

	// Add header parameters to the request
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Basic "+os.Getenv("SHINHAN_TOKEN"))

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return responses.ShinhanResponseData{}, err
	}

	defer resp.Body.Close()

	fmt.Println("StatusCode: ", resp.StatusCode)

	serverResponse, err := io.ReadAll(resp.Body)
	if err != nil {
		return responses.ShinhanResponseData{}, err
	}

	var responseData responses.ShinhanResponseData

	err = json.Unmarshal(serverResponse, &responseData)
	if err != nil {
		return responses.ShinhanResponseData{}, err
	}

	return responseData, nil
}

func getBCCToken() (string, error) {
	var respData responses.BCCTokenResponseData

	url := os.Getenv("BCC_TOKEN")

	payload := strings.NewReader("grant_type=client_credentials&scope=bcc.application.private")

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return "error", err
	}

	fmt.Println(os.Getenv("BCC_CRED"))

	req.Header.Add("authorization", "Basic "+os.Getenv("BCC_CRED"))
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("accept", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "error", err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "error", err
	}

	err = json.Unmarshal(body, &respData)
	if err != nil {
		return "error", err
	}

	return respData.AccessToken, nil
}

// TODO Перенести в Helpers
func encodeFileToBase64(filePath string) (string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return "error", err
	}

	reader := bufio.NewReader(f)
	content, err := io.ReadAll(reader)
	if err != nil {
		return "error", err
	}

	encoded := base64.StdEncoding.EncodeToString(content)

	return encoded, nil
}

func GetApplication(db *gorm.DB, id uint) (storage.Application, error) {
	application := storage.Application{}
	applicationGotten, err := application.Get(db, id)
	if err != nil {
		return storage.Application{}, err
	}

	return *applicationGotten, nil
}
