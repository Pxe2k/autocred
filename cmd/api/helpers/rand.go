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

func RandEmailCode() string {
	return fmt.Sprint(time.Now().Nanosecond())[:6]
}
