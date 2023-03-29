package helpers

import "strings"

func ImagesValidate(name string, size int64) bool {
	validate := false
	if size > 10485760 {
		return false
	}
	file := strings.Split(name, ".")
	fileExt := file[len(file)-1]
	switch fileExt {
	case "jpeg":
		validate = true
	case "jpg":
		validate = true
	case "png":
		validate = true
	}
	return validate
}
