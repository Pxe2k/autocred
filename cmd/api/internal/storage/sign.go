package storage

import "gorm.io/gorm"

type Sign struct {
	gorm.Model
	Sign          bool        `json:"sign"`
	UserID        uint        `json:"userID"`
	User          User        `json:"user"`
	ApplicationID uint        `json:"creditorID"`
	Application   Application `json:"creditor"`
}
