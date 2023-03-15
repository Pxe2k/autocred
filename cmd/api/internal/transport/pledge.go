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

func (server *Server) createPledge(w http.ResponseWriter, r *http.Request) {
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

	pledgeCreated, err := service.CreatePledgeService(server.DB, body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	responses.JSON(w, http.StatusCreated, pledgeCreated)
}

func (server *Server) allPledges(w http.ResponseWriter, r *http.Request) {
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

	pledge := storage.Pledge{}
	pledgesGotten, err := pledge.All(server.DB, uint(id))
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	responses.JSON(w, http.StatusOK, pledgesGotten)
}

func (server *Server) getPledge(w http.ResponseWriter, r *http.Request) {
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

	pledge := storage.Pledge{}
	pledgesGotten, err := pledge.Get(server.DB, uint(id))
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	responses.JSON(w, http.StatusOK, pledgesGotten)
}
