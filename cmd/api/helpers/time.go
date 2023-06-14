package helpers

import (
	"fmt"
	"time"
)

func CurrentDateString() string {
	currentTime := time.Now()
	formattedDate := currentTime.Format("02-01-06")
	dateString := fmt.Sprintf("%s", formattedDate)

	return dateString
}
