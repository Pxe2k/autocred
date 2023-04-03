package transport

import (
	"autocredit/cmd/api/helpers/responses"
	"autocredit/cmd/api/internal/storage"
	"encoding/json"
	"io"
	"net/http"
)

func (server *Server) createCarBrand(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	carBrand := storage.CarBrand{}
	err = json.Unmarshal(body, &carBrand)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	carBrandCreated, err := carBrand.Save(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	responses.JSON(w, http.StatusCreated, carBrandCreated)
}

func (server *Server) allCarBrands(w http.ResponseWriter, r *http.Request) {
	carBrand := storage.CarBrand{}
	carBrands, err := carBrand.All(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	responses.JSON(w, http.StatusOK, carBrands)
}
