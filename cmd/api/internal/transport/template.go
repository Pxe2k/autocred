package transport

import (
	"autocredit/cmd/api/helpers/responses"
	"autocredit/cmd/api/internal/service"
	"io"
	"net/http"
)

func (server *Server) GenerateTemplate(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	generatedPDF, err := service.GeneratePdf(server.DB, body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	responses.JSON(w, http.StatusCreated, generatedPDF)
}
