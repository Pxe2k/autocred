package storage

import (
	"errors"
	"github.com/badoux/checkmail"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"strings"
)

type User struct {
	gorm.Model
	Email                string        `gorm:"size:100;unique"`
	FullName             string        `gorm:"size:100;unique"`
	Phone                string        `gorm:"size:100;not null;unique"`
	Password             string        `gorm:"size:100;"`
	Creditor             bool          `json:"creditor"`
	BankID               *uint         `json:"bankID,omitempty"`
	Bank                 *Bank         `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"bank,omitempty"`
	UserApplications     []Application `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"userApplications"`
	CreditedApplications []Application `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"creditedApplications"`
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func (u *User) VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *User) BeforeSave() error {
	hashedPassword, err := Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.Password == "" {
			return errors.New("required password")
		}

		return nil
	case "login":
		if u.Password == "" {
			return errors.New("required password")
		}
		if u.Phone == "" {
			return errors.New("required Phone")
		}
		return nil

	default:
		if u.Password == "" {
			return errors.New("required password")
		}
		if u.Email == "" {
			return errors.New("required phone")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("invalid phone")
		}
		return nil
	}
}

func (u *User) SaveUser(db *gorm.DB) (*User, error) {
	err := u.BeforeSave()
	if err != nil {
		return &User{}, err
	}

	err = db.Debug().Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) Get(db *gorm.DB, uid uint) (*User, error) {
	err := db.Debug().Model(User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	errors.Is(err, gorm.ErrRecordNotFound)

	return u, nil
}
