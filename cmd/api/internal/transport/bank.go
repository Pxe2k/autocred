package transport

import (
	"autocredit/cmd/api/auth"
	"autocredit/cmd/api/helpers/responses"
	"autocredit/cmd/api/internal/service"
	"autocredit/cmd/api/internal/storage"
	"encoding/json"
	"io"
	"net/http"
)

func (server *Server) createBank(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	bank := storage.Bank{}
	err = json.Unmarshal(body, &bank)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	bankCreated, err := bank.Save(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	responses.JSON(w, http.StatusCreated, bankCreated)
}

func (server *Server) allBank(w http.ResponseWriter, r *http.Request) {
	bank := storage.Bank{}
	banks, err := bank.All(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	responses.JSON(w, http.StatusOK, banks)
}

func (server *Server) signApplication(w http.ResponseWriter, r *http.Request) {
	tokenID, err := auth.ExtractTokenID(r)
	if tokenID == 0 {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	signedResponse, err := service.SignApplication(server.DB, body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	responses.JSON(w, http.StatusAccepted, signedResponse)
}

func (server *Server) createProduct(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	product := storage.BankProduct{}
	err = json.Unmarshal(body, &product)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	productCreated, err := product.Save(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	responses.JSON(w, http.StatusCreated, productCreated)
}
