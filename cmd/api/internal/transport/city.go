package transport

import (
	"autocredit/cmd/api/helpers/responses"
	"autocredit/cmd/api/internal/storage"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (server *Server) allCity(w http.ResponseWriter, r *http.Request) {
	city := storage.City{}
	cities, err := city.All(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	responses.JSON(w, http.StatusOK, cities)
}

func (server *Server) findCityById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	city := storage.City{}
	cities, err := city.Find(server.DB, uint(id))
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	responses.JSON(w, http.StatusOK, cities)
}
