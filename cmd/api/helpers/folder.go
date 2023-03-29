package helpers

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"strconv"
	"strings"
)

func CreateFolder(uid uint32, title, folderName string, handler *multipart.FileHeader, file multipart.File) (string, error) {
	stringUid := strconv.FormatUint(uint64(uid), 10)
	storagePath, err := createStoragePath(folderName, stringUid)
	if err != nil {
		fmt.Println(err)
	}

	if storageInfo, err := os.Stat(storagePath); os.IsNotExist(err) {
		fmt.Println(err)
		if storageInfo == nil {
			if err = os.MkdirAll(storagePath, 0777); err != nil {
				fmt.Println(err)
			}
		}
	}

	fid := RandFileId()
	fidString := strconv.Itoa(fid)
	fileName := strings.Split(handler.Filename, ".")
	fileFmt := fileName[len(fileName)-1]
	storageFileName := "upload-" + title + fidString + "." + fileFmt
	deleteSpaces := strings.ReplaceAll(storageFileName, " ", "")

	filePath := storagePath + deleteSpaces

	err = os.Chmod(filePath, os.ModeAppend)
	if err != nil {
		fmt.Println(err)
	}

	tempFile, err := os.Create(filePath)
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	tempFile.Write(fileBytes)

	return filePath, nil
}

func createStoragePath(folderName, stringUid string) (string, error) {
	var err error
	if folderName == "user" {
		storagePath := "storage/client" + stringUid + "/"
		return storagePath, nil
	} else {
		return "media-storage path not found", err
	}
}
