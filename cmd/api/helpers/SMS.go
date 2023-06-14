package helpers

import (
	"autocredit/cmd/api/helpers/responses"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func SendMessage(OTP, phone string) error {
	url := os.Getenv("SMS_LINK")

	phoneNumber := strings.TrimPrefix(phone, "+")
	requestData := "recipient=" + phoneNumber[1:] + "&text=Your SMS Code - " + OTP

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(requestData)))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("cache-control", "no-cache")

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	serverResponse, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var responseData responses.SMSResponse
	fmt.Println("Server Response: ", responseData)

	err = json.Unmarshal(serverResponse, &responseData)
	if err != nil {
		return err
	}

	return nil
}
