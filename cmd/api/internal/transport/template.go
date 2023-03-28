package transport

import (
	"autocredit/cmd/api/auth"
	"autocredit/cmd/api/helpers/responses"
	"autocredit/cmd/api/internal/service"
	"errors"
	"io"
	"net/http"
)

func (server *Server) GenerateTemplate(w http.ResponseWriter, r *http.Request) {
	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("unauthorized"))
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	generatedPDF, err := service.GeneratePdf(server.DB, uid, body)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	/*w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(generatedPDF)*/
	responses.JSON(w, http.StatusCreated, generatedPDF)
}
