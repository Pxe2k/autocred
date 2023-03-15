package service

import (
	"autocredit/cmd/api/auth"
	"autocredit/cmd/api/helpers"
	"autocredit/cmd/api/helpers/requests"
	"autocredit/cmd/api/internal/storage"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"os"
)

var ctx = context.Background()

var Redis = redis.NewClient(&redis.Options{
	Addr:     os.Getenv("REDIS_ADDRESS") + ":" + os.Getenv("REDIS_PORT"),
	Password: "",
})

func CreateUserService(db *gorm.DB, body []byte) (*storage.User, error) {
	requestData := requests.UserRequestData{}

	err := json.Unmarshal(body, &requestData)
	if err != nil {
		return &storage.User{}, err
	}

	user := storage.User{}
	user.Email = requestData.Email
	user.Phone = requestData.Phone
	user.Password = requestData.Password
	user.Creditor = requestData.Creditor

	if user.Email == "" {
		return &storage.User{}, err
	}

	userCreated, err := user.SaveUser(db)
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
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		fmt.Println(err)
		return "error", err
	}

	authCode, err := generateCode(phone)
	if err != nil {
		return "error", err
	}

	return authCode, nil
}

func generateCode(phone string) (string, error) {
	authCode := helpers.RandEmailCode()

	err := Redis.Set(ctx, phone, authCode, 6000000000000).Err()
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

	val, err := Redis.Get(ctx, requestData.Phone).Result()
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

	return auth.CreateToken(uint32(user.ID))
}
