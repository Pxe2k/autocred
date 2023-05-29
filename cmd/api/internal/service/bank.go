package service

import (
	"autocredit/cmd/api/helpers"
	"autocredit/cmd/api/internal/storage"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"io"
	"mime/multipart"
	"os"
	"strconv"
	"strings"
)

func CreateBankService(db *gorm.DB, title string, file multipart.File, handler *multipart.FileHeader) (*storage.Bank, error) {
	bank := storage.Bank{}

	validateStatus := helpers.ImagesValidate(handler.Filename, handler.Size)
	if validateStatus != true {
		validateErr := errors.New("image not validated")
		return nil, validateErr
	}

	fid := helpers.RandFileId()
	fidString := strconv.Itoa(fid)
	fileName := strings.Split(handler.Filename, ".")
	fileFmt := fileName[len(fileName)-1]
	storageFileName := "upload-" + "bank" + fidString + "." + fileFmt
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

	bank.Image = &filePath
	bank.Title = title

	bankCreated, err := bank.Save(db)
	if err != nil {
		return nil, err
	}

	return bankCreated, nil
}

func UpdateBankService(db *gorm.DB, id uint32, title string, file multipart.File, handler *multipart.FileHeader) (*storage.Bank, error) {
	bank := storage.Bank{}

	validateStatus := helpers.ImagesValidate(handler.Filename, handler.Size)
	if validateStatus != true {
		validateErr := errors.New("image not validated")
		return nil, validateErr
	}

	fid := helpers.RandFileId()
	fidString := strconv.Itoa(fid)
	fileName := strings.Split(handler.Filename, ".")
	fileFmt := fileName[len(fileName)-1]
	storageFileName := "upload-" + "bank" + fidString + "." + fileFmt
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

	bank.Image = &filePath
	bank.Title = title

	bankUpdated, err := bank.Update(db, uint(id))
	if err != nil {
		return nil, err
	}

	return bankUpdated, nil
}
