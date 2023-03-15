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

	client := storage.Client{}
	clients, err := client.All(server.DB)
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
	if tokenID == 0 {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}

	client := storage.Client{}
	clientGotten, err := client.Get(server.DB, uint(id))
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	responses.JSON(w, http.StatusOK, clientGotten)
}
