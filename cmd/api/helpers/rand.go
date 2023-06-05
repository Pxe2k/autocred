package helpers

import (
	"fmt"
	"math/rand"
	"time"
)

func RandFileId() int {
	min := 10000000
	max := 99999999
	fid := rand.Intn(max-min) + min

	return fid
}

func RandBankApplicationID(length int) string {
	characters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	code := make([]rune, length)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < length; i++ {
		code[i] = characters[rand.Intn(len(characters))]
	}

	return string(code)
}

func RandEmailCode() string {
	return fmt.Sprint(time.Now().Nanosecond())[:6]
}
