package transport

import (
	"autocredit/cmd/api/auth"
	"autocredit/cmd/api/helpers/requests"
	"autocredit/cmd/api/helpers/responses"
	"autocredit/cmd/api/internal/service"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func (server *Server) createUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	userCreated, err := service.CreateUserService(server.DB, body)
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
