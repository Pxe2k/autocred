package helpers

import (
	"fmt"
	"time"
)

func RandEmailCode() string {
	return fmt.Sprint(time.Now().Nanosecond())[:6]
}
