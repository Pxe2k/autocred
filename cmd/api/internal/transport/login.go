package transport

import (
	"autocredit/cmd/api/auth"
	"autocredit/cmd/api/helpers/requests"
	"autocredit/cmd/api/helpers/responses"
	"autocredit/cmd/api/internal/service"
	"autocredit/cmd/api/internal/storage"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func (server *Server) createUser(w http.ResponseWriter, r *http.Request) {
	tokenID, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("unauthorized"))
		return
	}
	if tokenID == 0 {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("token is missing"))
		return
	}
	roleID, err := auth.ExtractRoleID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	var autoDealerID uint

	if roleID == 2 {
		user := storage.User{}
		userGotten, err1 := user.Get(server.DB, uint(tokenID))
		if err1 != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		autoDealerID = userGotten.AutoDealerID
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	userCreated, err := service.CreateUserService(server.DB, body, autoDealerID)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, userCreated.ID))

	responses.JSON(w, http.StatusOK, userCreated)
}

func (server *Server) login(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	requestData := requests.UserRequestData{}
	err = json.Unmarshal(body, &requestData)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	code, err := service.SignIn(requestData.Phone, requestData.Password, server.DB)
	if err != nil {
		passErr := errors.New("Incorrect Details")
		fmt.Println(err)
		responses.ERROR(w, http.StatusUnauthorized, passErr)
		return
	}

	responses.JSON(w, http.StatusOK, responses.LoginResponse{Code: code})
}

func (server *Server) ecp(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	responseObject := requests.ResponseObject{}

	err = json.Unmarshal(body, &responseObject)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	base64ResponseObject, err := base64.StdEncoding.DecodeString(responseObject.Data)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	stringBase64ResponseObject := string(base64ResponseObject)
	index := -1
	for i := 0; i < len(stringBase64ResponseObject)-2; i++ {
		if stringBase64ResponseObject[i:i+3] == "IIN" {
			index = i
			break
		}
	}

	if index != -1 && index+15 <= len(stringBase64ResponseObject) {
		substring := stringBase64ResponseObject[index+3 : index+15]
		decodedBytes, err := base64.StdEncoding.DecodeString(substring)
		if err != nil {
			fmt.Println("Error decoding base64:", err)
			return
		}
		fmt.Println(string(decodedBytes))
	} else {
		fmt.Println("Substring not found or out of range.")
	}

	responses.JSON(w, http.StatusOK, responseObject)
}

func (server *Server) submit(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	token, err := service.CreateToken(server.DB, body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	responses.JSON(w, http.StatusOK, responses.SubmitResponse{Token: &token})
}

func (server *Server) getRoleID(w http.ResponseWriter, r *http.Request) {
	roleID, err := auth.ExtractRoleID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	responses.JSON(w, http.StatusOK, responses.SubmitResponse{RoleID: &roleID})
}
