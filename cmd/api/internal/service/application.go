package service

import (
	"autocredit/cmd/api/helpers"
	"autocredit/cmd/api/helpers/requests"
	"autocredit/cmd/api/helpers/responses"
	"autocredit/cmd/api/internal/storage"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"gorm.io/gorm"
)

func CreateApplicationService(db *gorm.DB, body []byte, uid uint) (*storage.Application, error) {
	application := storage.Application{}
	responseData := responses.ApplicationResponseData{}
	err := json.Unmarshal(body, &application)
	if err != nil {
		return nil, err
	}

	application.UserID = uid

	applicationCreated, err := application.Save(db)
	if err != nil {
		return nil, err
	}

	responseData.Status = true

	return applicationCreated, nil
}

func SendApplications(db *gorm.DB, id uint, body []byte) (*storage.BankResponse, error) {
	application := storage.Application{}
	applicationGotten, err := application.Get(db, id)
	if err != nil {
		return nil, err
	}

	individualClient := storage.IndividualClient{}
	individualClientGotten, err := individualClient.Get(db, applicationGotten.IndividualClientID)
	if err != nil {
		return nil, err
	}

	var bankResponses []storage.BankResponse

	otpRequestData := requests.OTPShinhanRequestData{}

	if len(body) != 0 {
		err = json.Unmarshal(body, &otpRequestData)
		if err != nil {
			return nil, err
		}
	}

	// TODO if status ok create bankResponse
	for i := range applicationGotten.BankApplications {
		if application.BankApplications[i].BankID == 1 {
			bccResponseData, err1 := createBCCApplication(individualClientGotten, applicationGotten, application.BankApplications[i])
			if err1 != nil {
				fmt.Println("error:", err1)
			}
			if bccResponseData.Status == "OK" {
				bankResponses = append(bankResponses, storage.BankResponse{Status: "В ожидании", Description: bccResponseData.Message, ApplicationID: bccResponseData.RequestId, BankApplicationID: application.BankApplications[i].ID})
			} else {
				bankResponses = append(bankResponses, storage.BankResponse{Status: "В ожидании", Description: bccResponseData.Message, ApplicationID: bccResponseData.RequestId, BankApplicationID: application.BankApplications[i].ID})
			}
		} else if application.BankApplications[i].BankID == 2 {
			euBankResponseData, err2 := createEUApplication(individualClientGotten, applicationGotten, application.BankApplications[i])
			if err2 != nil {
				fmt.Println("error:", err2)
			}
			if euBankResponseData.Success == true {
				bankResponses = append(bankResponses, storage.BankResponse{Status: "В ожидании", Description: euBankResponseData.Msg, ApplicationID: euBankResponseData.OrderID, BankApplicationID: application.BankApplications[i].ID})
			} else {
				bankResponses = append(bankResponses, storage.BankResponse{Status: "Отказано", Description: euBankResponseData.Msg, ApplicationID: euBankResponseData.OrderID, BankApplicationID: application.BankApplications[i].ID})
			}
		} else if application.BankApplications[i].BankID == 3 {
			shinhanResponseData, err3 := createShinhanApplication(individualClientGotten, applicationGotten, application.BankApplications[i], otpRequestData.OTP)
			if err3 != nil {
				fmt.Println("error:", err3)
			}
			stringShinhanRequestID := strconv.Itoa(shinhanResponseData.ApplicationID)
			application.BankApplications[i].BankResponse.ApplicationID = stringShinhanRequestID
			bankResponses = append(bankResponses, storage.BankResponse{Status: "В ожидании", Description: "", ApplicationID: stringShinhanRequestID, BankApplicationID: application.BankApplications[i].ID})
		}
	}

	bankResponse := storage.BankResponse{}
	bankResponsesCreated, err := bankResponse.Save(db, bankResponses)
	if err != nil {
		return nil, err
	}

	return bankResponsesCreated, nil
}

func createBCCApplication(individualClient *storage.IndividualClient, application *storage.Application, bankApplication storage.BankApplication) (responses.BCCResponseData, error) {
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
	fmt.Println("bcc route", url)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		return responses.BCCResponseData{}, err
	}

	// Add header parameters to the request
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+authToken)

	//fmt.Println("token", authToken)

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("status code", resp.StatusCode)
		return responses.BCCResponseData{}, err
	}

	defer resp.Body.Close()

	fmt.Println("status code", resp.StatusCode)

	serverResponse, err := io.ReadAll(resp.Body)
	if err != nil {
		return responses.BCCResponseData{}, err
	}

	var responseData responses.BCCResponseData

	err = json.Unmarshal(serverResponse, &responseData)
	if err != nil {
		return responses.BCCResponseData{}, err
	}
	var result map[string]interface{}
	err = json.Unmarshal(serverResponse, &result)

	fmt.Println("result", result)

	return responseData, nil
}

func fillingBCCRequestData(client *storage.IndividualClient, applicationData *storage.Application, bankApplicationData storage.BankApplication) (requests.BCCApplicationRequestData, error) {
	var requestData requests.BCCApplicationRequestData

	issueYear, err := strconv.Atoi(applicationData.YearIssue)
	if err != nil {
		return requests.BCCApplicationRequestData{}, err
	}

	// TODO Поправить данные
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
	//for _, document := range applicationData.BankProcessingDocuments {
	//	if document.BankID == 1 {
	requestData.Document.File, err = helpers.EncodeFileToBase64("storage/bcc-data-processing" + strconv.Itoa(int(client.ID)) + "_" + helpers.CurrentDateString() + ".pdf")
	if err != nil {
		return requests.BCCApplicationRequestData{}, err
	}
	//	}
	//}
	requestData.Document.Extension = "pdf"
	requestData.Document.Code = "SOG"

	fmt.Println("processing data doc", requestData.Document.File)

	return requestData, nil
}

// TODO createApplication -> filling
func createEUApplication(individualClient *storage.IndividualClient, application *storage.Application, bankApplication storage.BankApplication) (responses.EUResponseData, error) {
	requestData, err := fillingEUBankRequestData(individualClient, application, bankApplication)
	if err != nil {
		return responses.EUResponseData{}, err
	}

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

	var result map[string]interface{}
	err = json.Unmarshal(serverResponse, &result)
	if err != nil {
		return responses.EUResponseData{}, err
	}
	fmt.Println("result", result)

	return responseData, nil
}

func fillingEUBankRequestData(client *storage.IndividualClient, applicationData *storage.Application, bankApplicationData storage.BankApplication) (requests.EUApplicationRequestData, error) {
	var requestData requests.EUApplicationRequestData

	fmt.Println("at start phone number", client.WorkPlaceInfo.OrganizationPhone)

	issueYear, err := strconv.Atoi(applicationData.YearIssue)
	if err != nil {
		return requests.EUApplicationRequestData{}, err
	}

	if applicationData.Condition == false {
		requestData.Car.Condition = "B2C"
	} else {
		requestData.Car.Condition = "NEW"
	}
	requestData.Car.Brand = applicationData.CarBrand
	requestData.Car.Model = applicationData.CarModel
	requestData.Car.Year = uint(issueYear)
	requestData.Car.Insurance = false
	requestData.Car.Price = uint(applicationData.CarPrice)
	requestData.City = "Алматы"
	requestData.Income = true
	requestData.PartyID = "11201740"
	requestData.DownPayment = uint(applicationData.InitFee)
	requestData.Duration = uint(bankApplicationData.TrenchesNumber)
	requestData.Iin = client.Document.IIN

	if client != nil && client.Phone != "" {
		// Remove any leading "+" symbol
		if strings.HasPrefix(client.Phone, "+") {
			client.Phone = client.Phone[2:]
		}
		// Remove leading "8" if present
		if strings.HasPrefix(client.Phone, "8") {
			client.Phone = client.Phone[1:]
		}
		requestData.Phone = client.Phone
	}

	if client != nil && client.WorkPlaceInfo != nil && client.WorkPlaceInfo.OrganizationPhone != "" {
		// Remove any leading "+" symbol
		if strings.HasPrefix(client.WorkPlaceInfo.OrganizationPhone, "+") {
			client.WorkPlaceInfo.OrganizationPhone = client.WorkPlaceInfo.OrganizationPhone[2:]
		}
		// Remove leading "8" if present
		if strings.HasPrefix(client.WorkPlaceInfo.OrganizationPhone, "8") {
			client.WorkPlaceInfo.OrganizationPhone = client.WorkPlaceInfo.OrganizationPhone[1:]
		}
		requestData.JobPhone = client.WorkPlaceInfo.OrganizationPhone
	}

	if client != nil && client.Contacts != nil && len(*client.Contacts) > 0 {
		for i := range *client.Contacts {
			contact := &(*client.Contacts)[i] // Get a pointer to the original element
			if contact != nil && contact.Phone != "" {
				// Remove any leading "+" symbol
				if strings.HasPrefix(contact.Phone, "+") {
					contact.Phone = contact.Phone[2:]
				}
				// Remove leading "8" if present
				if strings.HasPrefix(contact.Phone, "8") {
					contact.Phone = contact.Phone[1:]
				}
				requestData.ContactPersonContact = contact.Phone
			}
			requestData.ContactPersonName = contact.FullName
		}
	}

	requestData.IncomeMain = client.BonusInfo.AmountIncome

	switch client.MaritalStatus.Status {
	case "Холост/Не замужен":
		requestData.MaritalStatus = "1"
	case "Женат/Замужем":
		requestData.MaritalStatus = "2"
	case "Разведен/Разведена":
		requestData.MaritalStatus = "0"
	case "Гражданский брак":
		requestData.MaritalStatus = "4"
	case "Вдовец/вдова":
		requestData.MaritalStatus = "3"
	default:
		requestData.MaritalStatus = "0"
	}
	requestData.IncomeAddConfirmed = strconv.Itoa(0)

	requestData.Gsvp.Base64Content, err = helpers.EncodeFileToBase64("storage/eu-data-processing" + strconv.Itoa(int(client.ID)) + "_" + helpers.CurrentDateString() + ".pdf")
	if err != nil {
		return requests.EUApplicationRequestData{}, err
	}
	for _, document := range *client.Documents {
		if document.Title == "idFront" {
			requestData.Idcd.Base64Content, err = helpers.EncodeFileToBase64(document.File)
			if err != nil {
				return requests.EUApplicationRequestData{}, err
			}
		}
	}
	requestData.Photo.Base64Content, err = helpers.EncodeFileToBase64(client.Image)
	if err != nil {
		return requests.EUApplicationRequestData{}, err
	}

	requestData.OrderID = helpers.RandBankApplicationID(16)
	requestData.Gsvp.Name = "GSPV"
	requestData.Gsvp.Extension = "pdf"
	requestData.Idcd.Name = "IDCD"
	requestData.Idcd.Extension = "jpg"
	requestData.Photo.Name = "PHTO"
	requestData.Photo.Extension = "jpg"

	return requestData, nil
}

func createShinhanApplication(individualClient *storage.IndividualClient, application *storage.Application, bankApplication storage.BankApplication, otp string) (responses.ShinhanResponseData, error) {
	requestData, err := fillingShinhanBankRequestData(individualClient, application, bankApplication, otp)
	if err != nil {
		return responses.ShinhanResponseData{}, err
	}

	requestBody, err := json.Marshal(requestData)
	if err != nil {
		return responses.ShinhanResponseData{}, err
	}

	url := os.Getenv("SHINHAN_APPLICATION")
	fmt.Println("url: ", url)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println(req)
		return responses.ShinhanResponseData{}, err
	}

	// Add header parameters to the request
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Basic "+os.Getenv("SHINHAN_TOKEN"))

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("status code", resp.StatusCode)
		return responses.ShinhanResponseData{}, err
	}

	defer resp.Body.Close()

	fmt.Println("StatusCode: ", resp.StatusCode)

	serverResponse, err := io.ReadAll(resp.Body)
	if err != nil {
		return responses.ShinhanResponseData{}, err
	}

	var result map[string]interface{}
	err = json.Unmarshal(serverResponse, &result)
	if err != nil {
		fmt.Println(string(serverResponse))
		return responses.ShinhanResponseData{}, err
	}
	fmt.Println("result", result)

	var responseData responses.ShinhanResponseData
	err = json.Unmarshal(serverResponse, &responseData)
	if err != nil {
		fmt.Println(string(serverResponse))
		return responses.ShinhanResponseData{}, err
	}

	return responseData, nil
}

func fillingShinhanBankRequestData(client *storage.IndividualClient, applicationData *storage.Application, bankApplicationData storage.BankApplication, otp string) (requests.ShinhanApplicationRequestData, error) {
	var requestData requests.ShinhanApplicationRequestData
	var err error

	requestData.CalculationType = "A"
	requestData.Car.Brand = applicationData.CarBrand
	requestData.Car.Model = applicationData.CarModel
	requestData.Car.Year = applicationData.YearIssue
	requestData.Car.Country = "KOREAN"
	requestData.Car.Price = strconv.Itoa(applicationData.CarPrice)
	requestData.Car.FuelType = "GAZOLINE"
	requestData.Car.Colour = "белый"
	requestData.Car.Type = "SALOON"
	requestData.Car.Condition = "USED"
	requestData.Cas = false
	requestData.City = "Алматы " + client.User.AutoDealer.Address
	requestData.Customer.ActualAddress.District = client.ResidentialAddress.Address
	requestData.Customer.ActualAddress.Flat = client.ResidentialAddress.Address
	requestData.Customer.ActualAddress.House = client.ResidentialAddress.Address
	requestData.Customer.ActualAddress.Region = client.ResidentialAddress.Address
	requestData.Customer.ActualAddress.Settlement = client.ResidentialAddress.Address
	requestData.Customer.ActualAddress.Street = client.ResidentialAddress.Address
	requestData.Customer.BirthDate = "2002-03-22"
	requestData.Customer.BirthPlace = client.Document.PlaceOfBirth
	for _, contact := range *client.Contacts {
		requestData.Customer.ContactPersonPhone = contact.Phone[1:]
		requestData.Customer.ContactPersonFullName = contact.FullName
	}
	requestData.Customer.Document.CountryOfResidence = "KZ"
	requestData.Customer.Document.IssuedDate = "2018-03-22"
	requestData.Customer.Document.ExpirationDate = "2028-03-22"
	requestData.Customer.Document.Issuer = client.Document.IssuingAuthority
	requestData.Customer.Document.Number = client.Document.Number
	for _, media := range *client.Documents {
		if media.Title == "idBack" {
			requestData.Customer.Document.PhotoBack, err = helpers.EncodeFileToBase64(media.File)
			if err != nil {
				fmt.Println(1)
				return requests.ShinhanApplicationRequestData{}, err
			}
		}
		if media.Title == "idFront" {
			requestData.Customer.Document.PhotoFront, err = helpers.EncodeFileToBase64("storage/shinhan-data-processing" + strconv.Itoa(int(client.ID)) + "_" + helpers.CurrentDateString() + ".pdf")
			if err != nil {
				return requests.ShinhanApplicationRequestData{}, err
			}
		}
	}
	requestData.Customer.Document.Type = "ID_CARD"
	requestData.Customer.EmployerAddress.District = client.WorkPlaceInfo.Address
	requestData.Customer.EmployerAddress.Flat = client.WorkPlaceInfo.Address
	requestData.Customer.EmployerAddress.House = client.WorkPlaceInfo.Address
	requestData.Customer.EmployerAddress.Region = client.WorkPlaceInfo.Address
	requestData.Customer.EmployerAddress.Settlement = client.WorkPlaceInfo.Address
	requestData.Customer.EmployerAddress.Street = client.WorkPlaceInfo.Address
	requestData.Customer.EmployerName = client.WorkPlaceInfo.OrganizationName
	requestData.Customer.EmployerPhone = client.WorkPlaceInfo.OrganizationPhone
	requestData.Customer.EmploymentType = "PRIVATE_COMPANY"
	requestData.Customer.Firstname = client.FirstName
	requestData.Customer.Lastname = client.LastName
	requestData.Customer.Patronymic = client.MiddleName
	requestData.Customer.Gender = client.Sex
	requestData.Customer.Iin = client.Document.IIN
	requestData.Customer.Income = true
	requestData.Customer.MaritalStatus = "SINGLE"
	requestData.Customer.MobilePhone = "7751022255"
	requestData.Customer.NumberOfDependents = client.MaritalStatus.MinorChildren
	requestData.Customer.OfficialIncome = strconv.Itoa(client.BonusInfo.AmountIncome)
	requestData.Customer.Photo, err = helpers.EncodeFileToBase64("shinhan.png")
	if err != nil {
		return requests.ShinhanApplicationRequestData{}, err
	}
	requestData.Customer.RegistrationAddress.District = client.RegistrationAddress.Address
	requestData.Customer.RegistrationAddress.Flat = client.RegistrationAddress.Address
	requestData.Customer.RegistrationAddress.House = client.RegistrationAddress.Address
	requestData.Customer.RegistrationAddress.Region = client.RegistrationAddress.Address
	requestData.Customer.RegistrationAddress.Settlement = client.RegistrationAddress.Address
	requestData.Customer.RegistrationAddress.Street = client.RegistrationAddress.Address
	requestData.Customer.ResidencyStatus = "RESIDENT"
	requestData.Discount = false
	requestData.Downpayment = strconv.Itoa(applicationData.InitFee)
	requestData.Duration = strconv.Itoa(bankApplicationData.TrenchesNumber)
	requestData.GosProgram = false
	requestData.Grace = false
	requestData.InstalmentDate = "1991-03-22"
	requestData.Insurance = false
	requestData.PartnerId = "1778"
	requestData.Verification.Code = otp
	requestData.Verification.Date = time.Now().Format("2006-01-02 15:04:05")

	return requestData, nil
}

func getBCCToken() (string, error) {
	var respData responses.BCCTokenResponseData

	url := os.Getenv("BCC_TOKEN")

	payload := strings.NewReader("grant_type=client_credentials&scope=bcc.application.private")

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return "error", err
	}

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

func GetApplication(db *gorm.DB, id uint) (storage.Application, error) {
	application := storage.Application{}
	applicationGotten, err := application.Get(db, id)
	if err != nil {
		return storage.Application{}, err
	}

	return *applicationGotten, nil
}

func AllApplication(db *gorm.DB, uid uint) (*[]storage.Application, error) {
	application := storage.Application{}
	applications, err := application.All(db, uid)
	if err != nil {
		return nil, err
	}

	for i := range *applications {
		for j := range (*applications)[i].BankApplications {
			bankApplication := &(*applications)[i].BankApplications[j]
			if bankApplication.BankID == 2 {
				if bankApplication.BankResponse.ApplicationID != "" {
					statusResponse, err := getEUStatus(bankApplication.BankResponse.ApplicationID)
					if err != nil {
						fmt.Println("error: ", err)
					} else {
						bankApplication.BankResponse.Status = statusResponse.Status
						bankApplication.BankResponse.Description = statusResponse.Description
					}
				}
			} else if bankApplication.BankID == 3 {
				if bankApplication.BankResponse.ApplicationID != "" {
					status, err := getShinhanStatus(bankApplication.BankResponse.ApplicationID)
					if err != nil {
						fmt.Println("error: ", err)
					} else {
						bankApplication.BankResponse.Status = status
					}
				}
			}
		}
	}

	return applications, nil
}

func getShinhanStatus(shinhanApplicationID string) (string, error) {
	url := "https://is.shinhanfinance.kz/api/v1/orbis/application_status/" + shinhanApplicationID + "/"

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return "error: ", err
	}

	req.Header.Set("Authorization", "Basic "+os.Getenv("SHINHAN_TOKEN"))

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return "error: ", err
	}
	defer resp.Body.Close()

	if err != nil {
		return "error: ", err
	}

	serverResponse, err := io.ReadAll(resp.Body)
	if err != nil {
		return "error: ", err
	}

	var result map[string]interface{}
	err = json.Unmarshal(serverResponse, &result)
	if err != nil {
		fmt.Println(string(serverResponse))
		return "error", err
	}
	fmt.Println("result", result)

	responseData := responses.ShinhanStatusResponseData{}
	err = json.Unmarshal(serverResponse, &responseData)
	if err != nil {
		return "error: ", err
	}

	return responseData.Status, nil
}

func getEUStatus(euApplicationID string) (*responses.EUBankStatusResponseData, error) {
	url := "https://test-auto.eubank.kz/orbis/partner/" + euApplicationID

	fmt.Println(url)

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	serverResponse, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	responseData := responses.EUBankStatusResponseData{}
	err = json.Unmarshal(serverResponse, &responseData)
	if err != nil {
		return nil, err
	}

	return &responseData, nil
}
