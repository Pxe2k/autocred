package service

import (
	"autocredit/cmd/api/helpers/requests"
	"autocredit/cmd/api/helpers/responses"
	"autocredit/cmd/api/internal/storage"
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"gorm.io/gorm"
)

func CreateApplicationService(db *gorm.DB, body []byte, uid uint) (*storage.Application, error) {
	application := storage.Application{}
	err := json.Unmarshal(body, &application)
	if err != nil {
		return &storage.Application{}, err
	}

	application.UserID = uid

	createdApplication, err := application.Save(db)
	if err != nil {
		return &storage.Application{}, err
	}

	return createdApplication, nil
}

func CreateBCCApplication(body []byte) (responses.BCCApplicationData, error) {
	authToken, err := getBCCToken()
	if err != nil {
		return responses.BCCApplicationData{}, err
	}

	var requestData requests.BCCApplicationRequestData
	err = json.Unmarshal(body, &requestData)
	if err != nil {
		return responses.BCCApplicationData{}, err
	}

	requestData.Document.File, err = encodePDFtoBase64(requestData.Document.File)
	if err != nil {
		return responses.BCCApplicationData{}, err
	}

	fmt.Println("requestData: ", requestData)

	requestBody, err := json.Marshal(requestData)
	if err != nil {
		return responses.BCCApplicationData{}, err
	}

	url := os.Getenv("BCC_APPLICATION")
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		return responses.BCCApplicationData{}, err
	}

	// Add header parameters to the request
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+authToken)
	// Add more headers as needed

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return responses.BCCApplicationData{}, err
	}

	defer resp.Body.Close()

	serverResponse, err := io.ReadAll(resp.Body)
	if err != nil {
		return responses.BCCApplicationData{}, err
	}

	var responseData responses.BCCApplicationData

	err = json.Unmarshal(serverResponse, &responseData)
	if err != nil {
		return responses.BCCApplicationData{}, err
	}

	return responseData, nil
}

func getBCCToken() (string, error) {
	var respData responses.BCCTokenResponseData

	url := os.Getenv("BCC_TOKEN")

	payload := strings.NewReader("grant_type=client_credentials&scope=bcc.application.private")

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return "error", err
	}

	fmt.Println(os.Getenv("BCC_CRED"))

	req.Header.Add("authorization", "Basic "+os.Getenv("BCC_CRED"))
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("accept", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "error", err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "error", err
	}

	err = json.Unmarshal(body, &respData)
	if err != nil {
		return "error", err
	}

	return respData.AccessToken, nil
}

func encodePDFtoBase64(filePath string) (string, error) {
	f, err := os.Open("templates/resultMedia/outputPDF/autocredit.pdf")
	if err != nil {
		return "error", err
	}

	reader := bufio.NewReader(f)
	content, err := io.ReadAll(reader)
	if err != nil {
		return "error", err
	}

	encoded := base64.StdEncoding.EncodeToString(content)

	return encoded, nil
}

func GetApplication(db *gorm.DB, id uint) (storage.Application, error) {
	application := storage.Application{}
	applicationGotten, err := application.Get(db, id)
	if err != nil {
		return storage.Application{}, err
	}

	return *applicationGotten, nil
}
