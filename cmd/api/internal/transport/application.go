package transport

import (
	"autocredit/cmd/api/auth"
	"autocredit/cmd/api/helpers/responses"
	"autocredit/cmd/api/internal/service"
	"autocredit/cmd/api/internal/storage"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"strconv"
)

func (server *Server) createApplication(w http.ResponseWriter, r *http.Request) {
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
	if tokenID == 0 {
		responses.ERROR(w, http.StatusUnauthorized, err)
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

	applicationGotten, err := service.GetApplication(server.DB, uint(id), uint(tokenID))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	responses.JSON(w, http.StatusOK, applicationGotten)
}
