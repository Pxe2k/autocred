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

func (server *Server) createClient(w http.ResponseWriter, r *http.Request) {
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

	clientCreated, err := service.CreateClientService(server.DB, body, uint(tokenID))
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	responses.JSON(w, http.StatusCreated, clientCreated)
}

func (server *Server) allClients(w http.ResponseWriter, r *http.Request) {
	tokenID, err := auth.ExtractTokenID(r)
	if tokenID == 0 {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}

	userID := r.URL.Query().Get("user_id")
	fullName := r.URL.Query().Get("full_name")
	sex := r.URL.Query().Get("sex")
	birthDate := r.URL.Query().Get("birth_date")

	client := storage.Client{}
	clients, err := client.All(server.DB, fullName, sex, birthDate, userID)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	responses.JSON(w, http.StatusOK, clients)
}

func (server *Server) getClient(w http.ResponseWriter, r *http.Request) {
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

	client := storage.Client{}
	clientGotten, err := client.Get(server.DB, uint(id))
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	if clientGotten.UserID != uint(tokenID) {
		responseData := responses.UnauthorizedUserResponseData{}
		responseData.TypeOfClient = clientGotten.TypeOfClient
		responseData.Document = clientGotten.Document
		responseData.Phone = clientGotten.Phone
		responseData.ID = clientGotten.ID
		responseData.BirthDate = clientGotten.BirthDate
		responseData.CreatedAt = clientGotten.CreatedAt
		responseData.FirstName = clientGotten.FirstName
		responseData.MiddleName = clientGotten.MiddleName
		responseData.LastName = clientGotten.LastName
		responseData.ResidentialAddress = clientGotten.ResidentialAddress

		responses.JSON(w, http.StatusOK, responseData)
		return
	}

	responses.JSON(w, http.StatusOK, clientGotten)
}

func (server *Server) uploadAvatar(w http.ResponseWriter, r *http.Request) {
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

	tokenID, _ := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("unauthorized"))
		return
	}
	if tokenID == 0 {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("token is missing"))
		return
	}

	clientAvatar, err := service.UploadAvatarForClient(server.DB, uint32(clientID), file, handler)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, clientAvatar)
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

func (server *Server) updateClient(w http.ResponseWriter, r *http.Request) {
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

	updatedClient, err := service.UpdateClientInfo(server.DB, body, uint(clientID), uint(tokenID))
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	responses.JSON(w, http.StatusAccepted, updatedClient)
}
