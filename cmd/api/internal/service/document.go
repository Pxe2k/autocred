package service

import (
	"autocredit/cmd/api/helpers"
	"autocredit/cmd/api/internal/storage"
	"errors"
	"gorm.io/gorm"
	"mime/multipart"
)

func UploadFileService(db *gorm.DB, uid uint32, title, folderName string, clientID uint, handler *multipart.FileHeader, file multipart.File) (storage.Media, error) {
	defer file.Close()
	if uid == 0 {
		return storage.Media{}, errors.New("token null")
	}
	media := storage.Media{}
	filePath, err := helpers.CreateFolder(uid, title, folderName, handler, file)
	if err != nil {
		return storage.Media{}, err
	}
	media.File = filePath
	media.Title = title
	media.IndividualClientID = clientID
	if title == "" {
		media.Title = handler.Filename
	}
	//media.Format = strings.ToLower(filepath.Ext(handler.Filename))[1:]
	//validate := helpers.MediaValidate(handler.Filename, handler.Size)

	//if validate == false {
	//	return storage.Media{}, errors.New("wrong file format")
	//}
	mediaCreated, err := media.Save(db)

	return *mediaCreated, nil
}
