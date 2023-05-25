package transport

import (
	"autocredit/cmd/api/auth"
	"autocredit/cmd/api/helpers/responses"
	"autocredit/cmd/api/internal/storage"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (server *Server) createAutoDealer(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	autoDealer := storage.AutoDealer{}
	err = json.Unmarshal(body, &autoDealer)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	autoDealerCreated, err := autoDealer.Save(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	responses.JSON(w, http.StatusCreated, autoDealerCreated)
}

func (server *Server) updateAutoDealer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	autoDealerID, err := strconv.ParseUint(vars["id"], 10, 32)
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

	autoDealer := storage.AutoDealer{}
	err = json.Unmarshal(body, &autoDealer)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	autoDealerUpdate, err := autoDealer.Update(server.DB, int(autoDealerID))
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	responses.JSON(w, http.StatusCreated, autoDealerUpdate)
}

func (server *Server) deleteAutoDealer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bankProductID, err := strconv.ParseUint(vars["id"], 10, 32)
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

	autoDealer := storage.AutoDealer{}

	autoDealerDeleted, err := autoDealer.SoftDelete(server.DB, uint(bankProductID))
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	responses.JSON(w, http.StatusCreated, autoDealerDeleted)
}

func (server *Server) getAutoDealer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bankProductID, err := strconv.ParseUint(vars["id"], 10, 32)
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

	autoDealer := storage.AutoDealer{}

	autoDealerGotten, err := autoDealer.Get(server.DB, uint(bankProductID))
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	responses.JSON(w, http.StatusCreated, autoDealerGotten)
}

func (server *Server) allAutoDealers(w http.ResponseWriter, r *http.Request) {
	autoDealer := storage.AutoDealer{}
	autoDealers, err := autoDealer.All(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	responses.JSON(w, http.StatusOK, autoDealers)
}
