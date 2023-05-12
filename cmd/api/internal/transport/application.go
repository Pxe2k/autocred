package transport

import (
	"autocredit/cmd/api/auth"
	"autocredit/cmd/api/helpers/responses"
	"autocredit/cmd/api/internal/service"
	"autocredit/cmd/api/internal/storage"
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

var ctx = context.Background()

var Redis = redis.NewClient(&redis.Options{
	Addr:     os.Getenv("REDIS_ADDRESS") + ":" + os.Getenv("REDIS_PORT"),
	Password: "",
})

func (server *Server) createApplication(w http.ResponseWriter, r *http.Request) {
	tokenID, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}
	if tokenID == 0 {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	applicationCreated, err := service.CreateApplicationService(server.DB, body, uint(tokenID))

	responses.JSON(w, http.StatusCreated, applicationCreated)
}

func (server *Server) createBCCApplication(w http.ResponseWriter, r *http.Request) {
	tokenID, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}
	if tokenID == 0 {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	responseData, err := service.CreateBCCApplication(body)

	responses.JSON(w, http.StatusCreated, responseData)
}

func (server *Server) allApplications(w http.ResponseWriter, r *http.Request) {
	tokenID, err := auth.ExtractTokenID(r)
	if tokenID == 0 {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}

	application := storage.Application{}
	applications, err := application.All(server.DB, uint(tokenID))
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	responses.JSON(w, http.StatusOK, applications)
}

func (server *Server) getApplication(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	tokenID, err := auth.ExtractTokenID(r)
	if tokenID == 0 {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}

	applicationGotten, err := service.GetApplication(server.DB, uint(id))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	responses.JSON(w, http.StatusOK, applicationGotten)
}

func (server *Server) getBCCResponse(w http.ResponseWriter, r *http.Request) {
	tokenString := auth.ExtractToken(r)
	if tokenString == "" {
		responses.ERROR(w, http.StatusUnprocessableEntity, errors.New("empty token"))
		return
	}

	fmt.Println(tokenString)

	val, err := Redis.Get(ctx, "bcc").Result()
	if err == redis.Nil {
		err = errors.New("key does not exist")
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	} else if err != nil {
		fmt.Println(err)
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	if val != tokenString {
		responses.ERROR(w, http.StatusUnprocessableEntity, errors.New("wrong token"))
		return
	}

	m := make(map[string]string)
	m["message"] = "Accepted"

	responses.JSON(w, http.StatusAccepted, m)
}

func (server *Server) getBankToken(w http.ResponseWriter, r *http.Request) {
	var username, password string

	authHeader := r.Header.Get("Authorization")
	if authHeader != "" {
		// Check if the Authorization header starts with "Basic "
		if strings.HasPrefix(authHeader, "Basic ") {
			// Extract the base64-encoded username and password
			encodedCreds := authHeader[6:]
			creds, err := base64.StdEncoding.DecodeString(encodedCreds)
			if err == nil {
				// Split the decoded credentials into username and password
				credentials := strings.SplitN(string(creds), ":", 2)
				if len(credentials) == 2 {
					username = credentials[0]
					password = credentials[1]
				}
			}
		}
	}

	var loginStatus bool

	if username == "bcc" {
		if password == "xLTx6J9ddfl9F5sTU#lG8y30o" {
			loginStatus = true
		}
	}

	if loginStatus == false {
		err := errors.New("login error")
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	token, err := auth.CreateBankToken(username)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = setToken(username, token)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	m := make(map[string]string)
	m["message"] = token

	responses.JSON(w, http.StatusAccepted, m)
}

func setToken(bank, token string) error {
	err := Redis.Set(ctx, bank, token, 6000000000000).Err()
	if err != nil {
		return err
	}

	return nil
}
