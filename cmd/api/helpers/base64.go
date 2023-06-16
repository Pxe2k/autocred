package helpers

import (
	"bufio"
	"encoding/base64"
	"io"
	"os"
)

func EncodeFileToBase64(filePath string) (string, error) {
	f, err := os.Open(filePath)
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
