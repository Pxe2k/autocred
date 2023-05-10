package responses

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type LoginResponse struct {
	Code string `json:"code"`
}

type SubmitResponse struct {
	Token  string `json:"token"`
	RoleID uint32 `json:"roleID"`
}

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

func ERROR(w http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		JSON(w, statusCode, struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
		return
	}
	JSON(w, http.StatusBadRequest, nil)
}
