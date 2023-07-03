package transport

import (
	"autocredit/cmd/api/auth"
	"autocredit/cmd/api/helpers/responses"
	"autocredit/cmd/api/internal/service"
	"autocredit/cmd/api/internal/storage"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (server *Server) createIndividualClient(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	tokenID, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}

	clientCreated, err := service.CreateIndividualClientService(server.DB, body, uint(tokenID))
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	responses.JSON(w, http.StatusCreated, clientCreated)
}

func (server *Server) createBusinessClient(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	tokenID, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}

	clientCreated, err := service.CreateBusinessClientService(server.DB, body, uint(tokenID))
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	responses.JSON(w, http.StatusCreated, clientCreated)
}

func (server *Server) allIndividualClient(w http.ResponseWriter, r *http.Request) {
	tokenID, err := auth.ExtractTokenID(r)
	if tokenID == 0 {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}

	sortUserID := r.URL.Query().Get("sort_user_id")
	fullName := r.URL.Query().Get("full_name")
	sex := r.URL.Query().Get("sex")
	birthDate := r.URL.Query().Get("birth_date")

	client := storage.IndividualClient{}
	clients, err := client.All(server.DB, fullName, sex, birthDate, sortUserID, uint(tokenID))
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	for i := range *clients {
		if (*clients)[i].UserID != uint(tokenID) {
			(*clients)[i].Phone = ""
		}
	}

	responses.JSON(w, http.StatusOK, clients)
}

func (server *Server) allBusinessClient(w http.ResponseWriter, r *http.Request) {
	tokenID, err := auth.ExtractTokenID(r)
	if tokenID == 0 {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}

	sortUserID := r.URL.Query().Get("sort_user_id")
	fullName := r.URL.Query().Get("full_name")
	sex := r.URL.Query().Get("sex")
	birthDate := r.URL.Query().Get("birth_date")

	client := storage.BusinessClient{}
	clients, err := client.All(server.DB, fullName, sex, birthDate, sortUserID, uint(tokenID))
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	responses.JSON(w, http.StatusOK, clients)
}

func (server *Server) getIndividualClient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	tokenID, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}
	if tokenID == 0 {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}

	roleID, err := auth.ExtractRoleID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	clientGotten, err := service.GetIndividualClientService(server.DB, uint(id), uint(tokenID), uint(roleID))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	responses.JSON(w, http.StatusOK, clientGotten)
}

func (server *Server) getBusinessClient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	tokenID, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}
	if tokenID == 0 {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}

	clientGotten, err := service.GetBusinessClientService(server.DB, uint(id), uint(tokenID))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	responses.JSON(w, http.StatusOK, clientGotten)
}

func (server *Server) uploadIndividualClientAvatar(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	clientID, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	err = r.ParseMultipartForm(10 << 20)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	file, handler, err := r.FormFile("image")
	if err != nil {
		retrErr := errors.New("error while uploading file")
		responses.ERROR(w, http.StatusInternalServerError, retrErr)
		return
	}

	tokenID, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("unauthorized"))
		return
	}
	if tokenID == 0 {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("token is missing"))
		return
	}

	clientAvatar, err := service.UploadAvatarForIndividualClient(server.DB, uint32(clientID), file, handler)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, clientAvatar)
}

func (server *Server) uploadBusinessClientAvatar(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	clientID, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	err = r.ParseMultipartForm(10 << 20)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	file, handler, err := r.FormFile("image")
	if err != nil {
		retrErr := errors.New("error while uploading file")
		responses.ERROR(w, http.StatusInternalServerError, retrErr)
		return
	}

	tokenID, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("unauthorized"))
		return
	}
	if tokenID == 0 {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("token is missing"))
		return
	}

	clientAvatar, err := service.UploadAvatarForBusinessClient(server.DB, uint32(clientID), file, handler)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, clientAvatar)
}

func (server *Server) generateClientOTP(w http.ResponseWriter, r *http.Request) {
	tokenID, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("unauthorized"))
		return
	}
	if tokenID == 0 {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("token is missing"))
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	code, err := service.GenerateClientOTP(body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	responseData := make(map[string]string)
	responseData["code"] = code

	responses.JSON(w, http.StatusAccepted, responseData)
}

func (server *Server) submitIndividualClientOTP(w http.ResponseWriter, r *http.Request) {
	tokenID, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("unauthorized"))
		return
	}
	if tokenID == 0 {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("token is missing"))
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	status, err := service.SubmitIndividualClientOTP(server.DB, body, uint(tokenID))
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	responseData := make(map[string]string)
	responseData["message"] = status

	responses.JSON(w, http.StatusAccepted, responseData)
}

func (server *Server) submitBusinessClientOTP(w http.ResponseWriter, r *http.Request) {
	tokenID, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("unauthorized"))
		return
	}
	if tokenID == 0 {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("token is missing"))
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	status, err := service.SubmitBusinessClientOTP(server.DB, body, uint(tokenID))
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	responseData := make(map[string]string)
	responseData["message"] = status

	responses.JSON(w, http.StatusAccepted, responseData)
}

func (server *Server) issuingAuthorityAll(w http.ResponseWriter, r *http.Request) {
	issuingAuthority := storage.IssuingAuthority{}
	issuingAuthorities, err := issuingAuthority.All(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	responses.JSON(w, http.StatusOK, issuingAuthorities)
}

func (server *Server) updateIndividualClient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	clientID, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	tokenID, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("unauthorized"))
		return
	}
	if tokenID == 0 {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("token is missing"))
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	updatedIndividualClient, err := service.UpdateIndividualClient(server.DB, body, uint(clientID))
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	responses.JSON(w, http.StatusOK, updatedIndividualClient)
}
