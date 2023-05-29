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

func CreateClientService(db *gorm.DB, body []byte, uid uint) (*storage.Client, error) {
	client := storage.Client{}
	err := json.Unmarshal(body, &client)
	if err != nil {
		return &storage.Client{}, err
	}

	client.UserID = uid
	if client.Country == "Kazakhstan" {
		client.Residency = true
	} else {
		client.Residency = false
	}

	createdClient, err := client.Save(db)
	if err != nil {
		return &storage.Client{}, err
	}

	return createdClient, nil
}

func GetClientService(db *gorm.DB, id, tokenID uint) (responses.UserResponseData, error) {
	client := storage.Client{}
	responseData := responses.UserResponseData{}
	clientGotten, err := client.Get(db, id)
	if err != nil {
		return responses.UserResponseData{}, err
	}

	responseData.ID = clientGotten.ID
	responseData.TypeOfClient = clientGotten.TypeOfClient
	responseData.FirstName = clientGotten.FirstName
	responseData.MiddleName = clientGotten.MiddleName
	responseData.LastName = clientGotten.LastName
	responseData.BirthDate = clientGotten.BirthDate
	responseData.Phone = clientGotten.Phone
	responseData.Document = clientGotten.Document
	responseData.ResidentialAddress = clientGotten.ResidentialAddress
	responseData.CreatedAt = clientGotten.CreatedAt

	if clientGotten.UserID != tokenID {
		responseData.Status = false
		return responseData, nil
	}

	responseData.Status = true
	responseData.Sex = clientGotten.Sex
	responseData.Country = clientGotten.Country
	responseData.Residency = clientGotten.Residency
	responseData.Bin = clientGotten.Bin
	responseData.SecondPhone = clientGotten.SecondPhone
	responseData.Email = clientGotten.Email
	responseData.Education = clientGotten.Education
	responseData.Comment = clientGotten.Comment
	responseData.Image = clientGotten.Image
	responseData.UserID = clientGotten.UserID
	responseData.User = clientGotten.User
	responseData.MaritalStatus = clientGotten.MaritalStatus
	responseData.WorkPlaceInfo = clientGotten.WorkPlaceInfo
	responseData.RegistrationAddress = clientGotten.RegistrationAddress
	responseData.ResidentialAddress = clientGotten.ResidentialAddress
	responseData.Contacts = clientGotten.Contacts
	responseData.BonusInfo = clientGotten.BonusInfo
	responseData.PersonalProperty = clientGotten.PersonalProperty
	responseData.CurrentLoans = clientGotten.CurrentLoans
	responseData.BeneficialOwners = clientGotten.BeneficialOwners
	responseData.Pledges = clientGotten.Pledges
	responseData.Documents = clientGotten.Documents

	return responseData, nil
}

func UploadAvatarForClient(db *gorm.DB, uid uint32, file multipart.File, handler *multipart.FileHeader) (*storage.Client, error) {
	client := storage.Client{}
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

	return code, nil
}

func SubmitClientOTP(db *gorm.DB, body []byte, id uint) (string, error) {
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

	client := storage.Client{}
	client.UserID = id
	client.Phone = requestData.Phone

	err = client.UpdateUserID(db, client)
	if err != nil {
		return "error", err
	}

	return "success", nil
}
