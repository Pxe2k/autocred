package service

import (
	"autocredit/cmd/api/helpers"
	"autocredit/cmd/api/helpers/requests"
	"autocredit/cmd/api/helpers/responses"
	"autocredit/cmd/api/internal/storage"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"strconv"
	"strings"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func CreateIndividualClientService(db *gorm.DB, body []byte, uid uint) (*storage.IndividualClient, error) {
	client := storage.IndividualClient{}
	err := json.Unmarshal(body, &client)
	if err != nil {
		return &storage.IndividualClient{}, err
	}

	client.UserID = uid

	createdClient, err := client.Save(db)
	if err != nil {
		return &storage.IndividualClient{}, err
	}

	return createdClient, nil
}

func GetIndividualClientService(db *gorm.DB, id, tokenID, roleID uint) (responses.IndividualClientResponseData, error) {
	client := storage.IndividualClient{}
	responseData := responses.IndividualClientResponseData{}
	clientGotten, err := client.Get(db, id)
	if err != nil {
		return responses.IndividualClientResponseData{}, err
	}
	if clientGotten == nil {
		// Handle the case where clientGotten is nil (if applicable)
		return responses.IndividualClientResponseData{}, errors.New("client not found")
	}

	responseData.ID = clientGotten.ID
	responseData.TypeOfClient = clientGotten.TypeOfClient
	responseData.FirstName = clientGotten.FirstName
	responseData.MiddleName = clientGotten.MiddleName
	responseData.LastName = clientGotten.LastName
	responseData.BirthDate = clientGotten.BirthDate
	responseData.Phone = clientGotten.Phone
	responseData.Document.Number = clientGotten.Document.Number
	responseData.Document.IIN = clientGotten.Document.IIN
	responseData.Document.Type = clientGotten.Document.Type
	responseData.Document.IssuingAuthority = clientGotten.Document.IssuingAuthority
	responseData.Document.PlaceOfBirth = clientGotten.Document.PlaceOfBirth
	responseData.Document.DocumentIssueDate = clientGotten.Document.DocumentIssueDate
	responseData.Document.DocumentEndDate = clientGotten.Document.DocumentIssueDate
	responseData.RegistrationAddress.Address = clientGotten.RegistrationAddress.Address
	responseData.RegistrationAddress.Address1 = clientGotten.RegistrationAddress.Address1
	responseData.RegistrationAddress.Address2 = clientGotten.RegistrationAddress.Address2
	responseData.RegistrationAddress.Address3 = clientGotten.RegistrationAddress.Address3
	responseData.RegistrationAddress.Address4 = clientGotten.RegistrationAddress.Address4
	responseData.RegistrationAddress.Address5 = clientGotten.RegistrationAddress.Address5
	responseData.RegistrationAddress.Address6 = clientGotten.RegistrationAddress.Address6
	responseData.RegistrationAddress.Kato = clientGotten.RegistrationAddress.Kato
	responseData.CreatedAt = clientGotten.CreatedAt

	if roleID != 1 {
		if clientGotten.UserID != tokenID {
			responseData.Status = false
			return responseData, nil
		}
	}

	responseData.Status = true
	responseData.Sex = clientGotten.Sex
	responseData.Country = clientGotten.Country
	responseData.SecondPhone = clientGotten.SecondPhone
	responseData.Education = clientGotten.Education
	responseData.Comment = clientGotten.Comment
	responseData.Image = clientGotten.Image
	responseData.UserID = clientGotten.UserID
	responseData.User = clientGotten.User
	responseData.MaritalStatus = clientGotten.MaritalStatus
	responseData.WorkPlaceInfo.Experience = clientGotten.WorkPlaceInfo.Experience
	responseData.WorkPlaceInfo.WorkPlace = clientGotten.WorkPlaceInfo.WorkPlace
	responseData.WorkPlaceInfo.Address = clientGotten.WorkPlaceInfo.Address
	responseData.WorkPlaceInfo.OrganizationPhone = clientGotten.WorkPlaceInfo.OrganizationPhone
	responseData.WorkPlaceInfo.OrganizationName = clientGotten.WorkPlaceInfo.OrganizationName
	responseData.WorkPlaceInfo.MonthlyIncome = clientGotten.WorkPlaceInfo.MonthlyIncome
	responseData.WorkPlaceInfo.JobTitle = clientGotten.WorkPlaceInfo.JobTitle
	responseData.WorkPlaceInfo.DateNextSalary = clientGotten.WorkPlaceInfo.DateNextSalary
	responseData.WorkPlaceInfo.EmploymentDate = clientGotten.WorkPlaceInfo.EmploymentDate
	responseData.ResidentialAddress.Address = clientGotten.ResidentialAddress.Address
	responseData.ResidentialAddress.Address1 = clientGotten.ResidentialAddress.Address1
	responseData.ResidentialAddress.Address2 = clientGotten.ResidentialAddress.Address2
	responseData.ResidentialAddress.Address3 = clientGotten.ResidentialAddress.Address3
	responseData.ResidentialAddress.Address4 = clientGotten.ResidentialAddress.Address4
	responseData.ResidentialAddress.Address5 = clientGotten.ResidentialAddress.Address5
	responseData.ResidentialAddress.Address6 = clientGotten.ResidentialAddress.Address6
	responseData.ResidentialAddress.Kato = clientGotten.ResidentialAddress.Kato
	responseData.Contacts = clientGotten.Contacts
	responseData.BonusInfo = clientGotten.BonusInfo
	responseData.CurrentLoans = clientGotten.CurrentLoans
	responseData.BeneficialOwners = clientGotten.BeneficialOwners
	responseData.Pledges = clientGotten.Pledges
	responseData.Documents = clientGotten.Documents
	responseData.Applications = clientGotten.Applications

	return responseData, nil
}

func UploadAvatarForIndividualClient(db *gorm.DB, uid uint32, file multipart.File, handler *multipart.FileHeader) (*storage.IndividualClient, error) {
	client := storage.IndividualClient{}
	clientGotten, err := client.Get(db, uint(uid))
	if err != nil {
		return nil, nil
	}
	if clientGotten.Image != "" {
		err := os.Remove(clientGotten.Image)
		if err != nil {
			fmt.Println(err)
		}
	}

	validateStatus := helpers.ImagesValidate(handler.Filename, handler.Size)
	if validateStatus != true {
		validateErr := errors.New("image not validated")
		return nil, validateErr
	}

	fid := helpers.RandFileId()
	fidString := strconv.Itoa(fid)
	fileName := strings.Split(handler.Filename, ".")
	fileFmt := fileName[len(fileName)-1]
	storageFileName := "upload-" + "client-avatar" + fidString + "." + fileFmt
	deleteSpaces := strings.ReplaceAll(storageFileName, " ", "")
	filePath := "storage/" + deleteSpaces

	tempFile, err := os.Create(filePath)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer tempFile.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	tempFile.Write(fileBytes)

	clientGotten.Image = filePath

	updatedClient, err := clientGotten.UpdateAvatar(db, uint(uid))
	if err != nil {
		return nil, err
	}

	return updatedClient, nil
}

func GenerateClientOTP(body []byte) (string, error) {
	requestData := requests.OTPRequestData{}
	err := json.Unmarshal(body, &requestData)
	if err != nil {
		return "error", err
	}

	code, err := helpers.GenerateCode(requestData.Phone)
	if err != nil {
		return "error", err
	}

	serverENV := os.Getenv("SERVER")
	if serverENV == "PROD" {
		err = helpers.SendMessage(code, requestData.Phone)
		if err != nil {
			return "error", err
		}
	}

	return code, nil
}

func SubmitIndividualClientOTP(db *gorm.DB, body []byte, id uint) (string, error) {
	requestData := requests.OTPRequestData{}
	err := json.Unmarshal(body, &requestData)
	if err != nil {
		return "error", err
	}

	val, err := helpers.Redis.Get(helpers.Ctx, requestData.Phone).Result()
	if err == redis.Nil {
		fmt.Println("key does not exist")
		return "key does not exist", err
	} else if err != nil {
		fmt.Println(err)
		return "error ", err
	}

	if val != requestData.Code {
		return "wrong code", errors.New("code != value")
	}

	client := storage.IndividualClient{}
	client.UserID = id
	client.Phone = requestData.Phone

	err = client.UpdateUserID(db, client)
	if err != nil {
		return "error", err
	}

	return "success", nil
}

func CreateBusinessClientService(db *gorm.DB, body []byte, uid uint) (*storage.BusinessClient, error) {
	client := storage.BusinessClient{}
	err := json.Unmarshal(body, &client)
	if err != nil {
		return &storage.BusinessClient{}, err
	}

	client.UserID = uid

	createdClient, err := client.Save(db)
	if err != nil {
		return &storage.BusinessClient{}, err
	}

	return createdClient, nil
}

func GetBusinessClientService(db *gorm.DB, id, tokenID uint) (responses.BusinessClientResponseData, error) {
	client := storage.BusinessClient{}
	responseData := responses.BusinessClientResponseData{}
	clientGotten, err := client.Get(db, id)
	if err != nil {
		return responses.BusinessClientResponseData{}, err
	}

	responseData.ID = clientGotten.ID
	responseData.TypeOfClient = clientGotten.TypeOfClient
	responseData.CompanyName = clientGotten.CompanyName
	responseData.CompanyPhone = clientGotten.CompanyPhone
	responseData.CompanyLifespan = clientGotten.CompanyLifespan
	responseData.ActivityType = clientGotten.ActivityType
	responseData.KindActivity = clientGotten.KindActivity
	responseData.RegistrationDate = clientGotten.RegistrationDate
	responseData.CreatedAt = clientGotten.CreatedAt

	if clientGotten.UserID != tokenID {
		responseData.Status = false
		return responseData, nil
	}

	responseData.Status = true
	responseData.Image = clientGotten.Image
	responseData.UserID = clientGotten.UserID
	responseData.User = clientGotten.User
	responseData.RegistrationAddress = clientGotten.RegistrationAddress
	responseData.BeneficialOwner = clientGotten.BeneficialOwner

	return responseData, nil
}

func UploadAvatarForBusinessClient(db *gorm.DB, uid uint32, file multipart.File, handler *multipart.FileHeader) (*storage.BusinessClient, error) {
	client := storage.BusinessClient{}
	clientGotten, err := client.Get(db, uint(uid))
	if err != nil {
		return nil, nil
	}
	if clientGotten.Image != "" {
		err := os.Remove(clientGotten.Image)
		if err != nil {
			fmt.Println(err)
		}
	}

	validateStatus := helpers.ImagesValidate(handler.Filename, handler.Size)
	if validateStatus != true {
		validateErr := errors.New("image not validated")
		return nil, validateErr
	}

	fid := helpers.RandFileId()
	fidString := strconv.Itoa(fid)
	fileName := strings.Split(handler.Filename, ".")
	fileFmt := fileName[len(fileName)-1]
	storageFileName := "upload-" + "client-avatar" + fidString + "." + fileFmt
	deleteSpaces := strings.ReplaceAll(storageFileName, " ", "")
	filePath := "storage/" + deleteSpaces

	tempFile, err := os.Create(filePath)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer tempFile.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	tempFile.Write(fileBytes)

	clientGotten.Image = filePath

	updatedClient, err := clientGotten.UpdateAvatar(db, uint(uid))
	if err != nil {
		return nil, err
	}

	return updatedClient, nil
}

func SubmitBusinessClientOTP(db *gorm.DB, body []byte, id uint) (string, error) {
	requestData := requests.OTPRequestData{}
	err := json.Unmarshal(body, &requestData)
	if err != nil {
		return "error", err
	}

	val, err := helpers.Redis.Get(helpers.Ctx, requestData.Phone).Result()
	if err == redis.Nil {
		fmt.Println("key does not exist")
		return "key does not exist", err
	} else if err != nil {
		fmt.Println(err)
		return "error ", err
	}

	if val != requestData.Code {
		return "wrong code", errors.New("code != value")
	}

	client := storage.BusinessClient{}
	client.UserID = id
	client.CompanyPhone = requestData.Phone

	err = client.UpdateUserID(db, client)
	if err != nil {
		return "error", err
	}

	return "success", nil
}

func UpdateIndividualClient(db *gorm.DB, body []byte, clientID uint) (string, error) {
	individualClient := storage.IndividualClient{}
	err := json.Unmarshal(body, &individualClient)
	if err != nil {
		return "error", err
	}

	err = individualClient.Update(db, clientID)
	if err != nil {
		return "error", err
	}

	if individualClient.Document != nil {
		document := storage.Document{}
		err = document.Update(db, individualClient.Document, clientID)
		if err != nil {
			return "error", err
		}
	}

	if individualClient.RegistrationAddress != nil {
		registrationAddress := storage.RegistrationAddress{}
		err = registrationAddress.Update(db, individualClient.RegistrationAddress, clientID)
		if err != nil {
			return "error", err
		}
	}

	if individualClient.ResidentialAddress != nil {
		residentialAddress := storage.ResidentialAddress{}
		err = residentialAddress.Update(db, individualClient.ResidentialAddress, clientID)
		if err != nil {
			return "error", err
		}
	}

	if individualClient.WorkPlaceInfo != nil {
		workPlaceInfo := storage.WorkPlaceInfo{}
		err = workPlaceInfo.Update(db, individualClient.WorkPlaceInfo, clientID)
		if err != nil {
			return "error", err
		}
	}

	return "success", nil
}
