package storage

import (
	"errors"
	"strings"

	"github.com/badoux/checkmail"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email          string        `gorm:"size:100,unique"`
	FirstName      string        `gorm:"size:100" json:"firstName"`
	MiddleName     string        `gorm:"size:100" json:"middleName"`
	LastName       *string       `gorm:"size:100" json:"lastName,omitempty"`
	IIN            string        `gorm:"size:100,unique" json:"iin"`
	Document       string        `gorm:"size:100" json:"document"`
	DocumentNumber string        `gorm:"size:100" json:"documentNumber"`
	JobTitle       string        `gorm:"size:100" json:"jobTitle"`
	OrderNumber    string        `gorm:"size:100" json:"orderNumber"`
	Phone          string        `gorm:"size:100,unique" json:"phone"`
	WorkPhone      string        `gorm:"size:100" json:"workPhone"`
	Password       string        `gorm:"size:100"`
	AutoDealerID   uint          `json:"autoDealerID,omitempty"`
	AutoDealer     *AutoDealer   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"autodealer,omitempty"`
	RoleID         *uint         `json:"roleID,omitempty"`
	Role           Role          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"role"`
	Applications   []Application `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"applications"`
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

func (u *User) Save(db *gorm.DB) (*User, error) {
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

func (u *User) All(db *gorm.DB, autodealerID uint) (*[]User, error) {
	var users []User
	err := db.Debug().Model(&User{}).Where("auto_dealer_id = ?", autodealerID).Preload("AutoDealer").Preload("Role").Limit(100).Find(&users).Error
	if err != nil {
		return nil, err
	}

	return &users, nil
}

func (u *User) Update(db *gorm.DB, id int) (*User, error) {
	err := db.Debug().Model(&User{}).Where("id = ?", id).Updates(&u).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (u *User) SoftDelete(db *gorm.DB, id uint) (int64, error) {
	err := db.Debug().Model(&User{}).Where("id = ?", id).Take(&User{}).Delete(&User{})
	if err != nil {
		return 0, err.Error
	}

	if err.Error != nil {
		return 0, err.Error
	}

	return err.RowsAffected, nil
}

func (u *User) AllSoftDeleted(db *gorm.DB, autodealerID uint) (*[]User, error) {
	var users []User
	err := db.Debug().Model(&User{}).
		Unscoped(). // Include soft deleted records
		Where("auto_dealer_id = ?", autodealerID).
		Where("deleted_at IS NOT NULL"). // Filter soft deleted records
		Preload("AutoDealer").
		Preload("Role").
		Limit(100).
		Find(&users).Error
	if err != nil {
		return nil, err
	}

	return &users, nil
}

func (u *User) Recover(db *gorm.DB, id uint) error {
	return db.Unscoped().Model(&User{}).Where("id = ?", id).Update("deleted_at", nil).Error
}
