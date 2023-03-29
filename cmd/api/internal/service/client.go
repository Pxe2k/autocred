package service

import (
	"autocredit/cmd/api/helpers"
	"autocredit/cmd/api/internal/storage"
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"io"
	"mime/multipart"
	"os"
	"strconv"
	"strings"
)

func CreateClientService(db *gorm.DB, body []byte, uid uint) (*storage.Client, error) {
	client := storage.Client{}
	err := json.Unmarshal(body, &client)
	if err != nil {
		return &storage.Client{}, err
	}

	client.UserID = uid

	createdClient, err := client.Save(db)
	if err != nil {
		return &storage.Client{}, err
	}

	return createdClient, nil
}

func UploadAvatarForClient(db *gorm.DB, uid uint32, file multipart.File, handler *multipart.FileHeader) (*storage.Client, error) {
	client := storage.Client{}
	clientGotten, err := client.Get(db, uint(uid))
	if err != nil {
		return nil, nil
	}
	if clientGotten.Image != "" {
		err := os.Remove(clientGotten.Image)
		if err != nil {
			fmt.Println(err)
		}
	}

	validateStatus := helpers.ImagesValidate(handler.Filename, handler.Size)
	if validateStatus != true {
		validateErr := errors.New("image not validated")
		return nil, validateErr
	}

	fid := helpers.RandFileId()
	fidString := strconv.Itoa(fid)
	fileName := strings.Split(handler.Filename, ".")
	fileFmt := fileName[len(fileName)-1]
	storageFileName := "upload-" + "client-avatar" + fidString + "." + fileFmt
	deleteSpaces := strings.ReplaceAll(storageFileName, " ", "")
	filePath := "storage/" + deleteSpaces

	tempFile, err := os.Create(filePath)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer tempFile.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	tempFile.Write(fileBytes)

	clientGotten.Image = filePath

	updatedClient, err := clientGotten.UpdateAvatar(db, uint(uid))
	if err != nil {
		return nil, err
	}

	return updatedClient, nil
}
