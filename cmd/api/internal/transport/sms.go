package transport

import (
	"autocredit/cmd/api/helpers/responses"
	"encoding/json"
	"io"
	"net/http"
)

func (server *Server) getBalance(w http.ResponseWriter, r *http.Request) {
	url := "https://api.mobizon.kz/service/user/getownbalance?apiKey=kz0255ed64ae757890471a84d91f26408f62685a7d584d8b26c2ff1e98f9d0205a39e1"

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
	}

	//// Add header parameters to the request
	//req.Header.Set("Content-Type", "application/json")
	//req.Header.Set("x-eub-token", os.Getenv("EU_TOKEN"))

	resp, err := client.Do(req)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
	}
	defer resp.Body.Close()

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
	}

	serverResponse, err := io.ReadAll(resp.Body)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
	}

	var result map[string]interface{}
	err = json.Unmarshal(serverResponse, &result)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)

	}

	responses.JSON(w, http.StatusOK, result)
}
