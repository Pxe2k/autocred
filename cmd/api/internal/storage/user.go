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
	Email        string        `gorm:"size:100;unique"`
	FullName     string        `gorm:"size:100;" json:"fullName"`
	Phone        string        `gorm:"size:100;not null;unique"`
	Address      string        `gorm:"size:100;"`
	Password     string        `gorm:"size:100;"`
	RoleID       *uint         `gorm:"default:2;" json:"roleID,omitempty"`
	Role         Role          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"role"`
	Applications []Application `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"applications"`
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
