package transport

import (
	"autocredit/cmd/api/auth"
	"autocredit/cmd/api/helpers/responses"
	"autocredit/cmd/api/internal/service"
	"autocredit/cmd/api/internal/storage"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"strconv"
)

func (server *Server) createBank(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	title := r.FormValue("title")

	file, handler, err := r.FormFile("image")
	if err == http.ErrMissingFile {
		bank := storage.Bank{}
		bank.Title = title
		bankCreated, err1 := bank.Save(server.DB)
		if err1 != nil {
			responses.ERROR(w, http.StatusInternalServerError, err)
			return
		}
		responses.JSON(w, http.StatusOK, bankCreated)
		return
	}
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
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

	bankCreated, err := service.CreateBankService(server.DB, title, file, handler)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, bankCreated)
}

func (server *Server) updateBank(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bankID, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	title := r.FormValue("title")

	file, handler, err := r.FormFile("image")
	if err == http.ErrMissingFile {
		bank := storage.Bank{}
		bank.Title = title
		bankUpdated, err1 := bank.Update(server.DB, uint(bankID))
		if err1 != nil {
			responses.ERROR(w, http.StatusInternalServerError, err)
			return
		}
		responses.JSON(w, http.StatusOK, bankUpdated)
		return
	}
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
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

	bankUpdated, err := service.UpdateBankService(server.DB, uint32(bankID), title, file, handler)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, bankUpdated)
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

//func (server *Server) updateBank(w http.ResponseWriter, r *http.Request) {
//	vars := mux.Vars(r)
//	bankID, err := strconv.ParseUint(vars["id"], 10, 32)
//	if err != nil {
//		responses.ERROR(w, http.StatusBadRequest, err)
//		return
//	}
//
//	tokenID, err := auth.ExtractTokenID(r)
//	if err != nil {
//		responses.ERROR(w, http.StatusUnauthorized, errors.New("unauthorized"))
//		return
//	}
//	if tokenID == 0 {
//		responses.ERROR(w, http.StatusUnauthorized, errors.New("token is missing"))
//		return
//	}
//
//	body, err := io.ReadAll(r.Body)
//	if err != nil {
//		responses.ERROR(w, http.StatusUnprocessableEntity, err)
//		return
//	}
//
//	bank := storage.Bank{}
//	err = json.Unmarshal(body, &bank)
//	if err != nil {
//		responses.ERROR(w, http.StatusUnprocessableEntity, err)
//		return
//	}
//
//	bankUpdate, err := bank.Update(server.DB, int(bankID))
//	if err != nil {
//		responses.ERROR(w, http.StatusUnprocessableEntity, err)
//		return
//	}
//
//	responses.JSON(w, http.StatusCreated, bankUpdate)
//}

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

func (server *Server) updateProduct(w http.ResponseWriter, r *http.Request) {
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

	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	bankProduct := storage.BankProduct{}
	err = json.Unmarshal(body, &bankProduct)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	bankProductUpdate, err := bankProduct.Update(server.DB, int(bankProductID))
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	responses.JSON(w, http.StatusCreated, bankProductUpdate)
}

func (server *Server) deleteBank(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bankID, err := strconv.ParseUint(vars["id"], 10, 32)
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

	bank := storage.Bank{}

	bankDeleted, err := bank.SoftDelete(server.DB, uint(bankID))
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	responses.JSON(w, http.StatusCreated, bankDeleted)
}

func (server *Server) deleteBankProduct(w http.ResponseWriter, r *http.Request) {
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

	bankProduct := storage.BankProduct{}

	bankProductDeleted, err := bankProduct.SoftDelete(server.DB, uint(bankProductID))
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	responses.JSON(w, http.StatusCreated, bankProductDeleted)
}

func (server *Server) getBankProduct(w http.ResponseWriter, r *http.Request) {
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

	bankProduct := storage.BankProduct{}

	bankProductGotten, err := bankProduct.Get(server.DB, uint(bankProductID))
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	responses.JSON(w, http.StatusCreated, bankProductGotten)
}
