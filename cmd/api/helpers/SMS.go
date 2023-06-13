package helpers

import (
	"bytes"
	"net/http"
	"os"
)

func SendMessage(OTP, phone string) error {
	SMSKey := os.Getenv("SMS_KEY")
	url := os.Getenv("SMS_LINK")

	requestData := "recipient=" + phone + "&text="

	req, err := http.NewRequest("POST", url, bytes.NewBuffer())
	if err != nil {
		return err
	}
}
