package service

import (
	"autocredit/cmd/api/helpers/responses"
	"autocredit/cmd/api/internal/storage"
	"gorm.io/gorm"
)

func GetUserService(db *gorm.DB, id uint) (responses.UserResponseData, error) {
	user := storage.User{}
	userGotten, err := user.Get(db, id)
	if err != nil {
		return responses.UserResponseData{}, err
	}
	responseData := responses.UserResponseData{}
	responseData.ID = userGotten.ID
	responseData.Email = userGotten.Email
	responseData.FirstName = userGotten.FirstName
	responseData.MiddleName = userGotten.MiddleName
	responseData.LastName = userGotten.LastName
	responseData.Document = userGotten.Document
	responseData.DocumentNumber = userGotten.DocumentNumber
	responseData.IIN = userGotten.IIN
	responseData.JobTitle = userGotten.JobTitle
	responseData.OrderNumber = userGotten.OrderNumber
	responseData.Phone = userGotten.Phone
	responseData.WorkPhone = userGotten.WorkPhone
	responseData.AutoDealerID = userGotten.AutoDealerID
	responseData.AutoDealer = userGotten.AutoDealer
	responseData.RoleID = userGotten.RoleID
	responseData.Role = userGotten.Role
	responseData.Applications = userGotten.Applications

	return responseData, nil
}
