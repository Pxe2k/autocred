package transport

import (
	"autocredit/cmd/api/auth"
	"autocredit/cmd/api/helpers"
	"autocredit/cmd/api/helpers/responses"
	"autocredit/cmd/api/internal/service"
	"autocredit/cmd/api/internal/storage"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/redis/go-redis/v9"

	"github.com/gorilla/mux"
)

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
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	responses.JSON(w, http.StatusCreated, applicationCreated)
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

	val, err := helpers.Redis.Get(helpers.Ctx, "bcc").Result()
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

	err = helpers.SetToken(username, token)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	m := make(map[string]string)
	m["message"] = token

	responses.JSON(w, http.StatusAccepted, m)
}
