package service

import (
	"autocredit/cmd/api/auth"
	"autocredit/cmd/api/helpers"
	"autocredit/cmd/api/helpers/requests"
	"autocredit/cmd/api/internal/storage"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func CreateUserService(db *gorm.DB, body []byte, autoDealerID uint) (*storage.User, error) {
	user := storage.User{}
	requestData := requests.UserRequestData{}
	err := json.Unmarshal(body, &requestData)
	if err != nil {
		return &storage.User{}, err
	}

	user.Email = requestData.Email
	user.FirstName = requestData.FirstName
	user.MiddleName = requestData.MiddleName
	user.LastName = requestData.LastName
	if requestData.ResponseObject != "" {
		iin, err1 := takeDataFromResponseObject(requestData.ResponseObject)
		if err1 != nil {
			return nil, err1
		}
		user.IIN = &iin
	} else {
		user.IIN = nil
	}
	user.Document = requestData.Document
	user.DocumentNumber = requestData.DocumentNumber
	user.JobTitle = requestData.JobTitle
	user.OrderNumber = requestData.OrderNumber
	user.Phone = requestData.Phone
	user.WorkPhone = requestData.WorkPhone
	user.Password = requestData.Password
	user.RoleID = requestData.RoleID

	if autoDealerID != 0 {
		user.AutoDealerID = autoDealerID
	} else {
		user.AutoDealerID = requestData.AutoDealerID
	}

	if user.Email == "" {
		return &storage.User{}, err
	}

	userCreated, err := user.Save(db)
	if err != nil {
		return nil, err
	}

	return userCreated, nil
}

func SignIn(phone, password string, db *gorm.DB) (string, error) {
	user := storage.User{}
	user.Phone = phone
	user.Password = password
	err := user.Validate("login")
	if err != nil {
		fmt.Println(err)
		return "Error ", err
	}

	err = db.Debug().Model(storage.User{}).Where("phone = ?", phone).Take(&user).Error
	if err != nil {
		fmt.Println(err)
		return "error", err
	}

	err = user.VerifyPassword(user.Password, password)
	if err != nil {
		return "error", err
	}

	authCode, err := helpers.GenerateCode(phone)
	if err != nil {
		return "error", err
	}

	err = helpers.SendMessage(authCode, phone)
	if err != nil {
		return "error", err
	}

	return authCode, nil
}

func CreateToken(db *gorm.DB, body []byte) (string, error) {
	requestData := requests.SubmitCode{}
	err := json.Unmarshal(body, &requestData)
	if err != nil {
		fmt.Println(err)
		return "error ", err
	}

	user := storage.User{}
	err = db.Debug().Model(storage.User{}).Where("phone = ?", requestData.Phone).Take(&user).Error
	if err != nil {
		fmt.Println(err)
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

	return auth.CreateToken(uint32(user.ID), user.RoleID, user.AutoDealerID)
}

func ECPDecode(db *gorm.DB, body []byte) (string, error) {
	responseObject := requests.ResponseObject{}

	err := json.Unmarshal(body, &responseObject)
	if err != nil {
		return "error", err
	}

	iin, err := takeDataFromResponseObject(responseObject.Data)
	if err != nil {
		return "error", err
	}

	user := storage.User{}
	err = db.Debug().Model(storage.User{}).Where("iin = ?", iin).Take(&user).Error
	if err != nil {
		return "error", err
	}

	return auth.CreateToken(uint32(user.ID), user.RoleID, user.AutoDealerID)
}

func takeDataFromResponseObject(data string) (string, error) {
	base64ResponseObject, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "error", err
	}

	position := strings.Index(string(base64ResponseObject), "IIN")
	iin := strings.TrimSpace(string(base64ResponseObject[position+3 : position+15]))

	return iin, err
}
