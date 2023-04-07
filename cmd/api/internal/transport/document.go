package transport

import (
	"autocredit/cmd/api/auth"
	"autocredit/cmd/api/helpers/responses"
	"autocredit/cmd/api/internal/service"
	"fmt"
	"net/http"
	"strconv"
)

func (server *Server) uploadFile(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		fmt.Println(err)
	}
	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		fmt.Println(err)
	}
	title := r.FormValue("title")
	clientIDString := r.FormValue("clientID")
	clientID, err := strconv.Atoi(clientIDString)
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}

	mediaCreated, err := service.UploadFileService(server.DB, uid, title, "user", uint(clientID), handler, file)
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, err)
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, mediaCreated.ID))
	responses.JSON(w, http.StatusCreated, mediaCreated)
}
