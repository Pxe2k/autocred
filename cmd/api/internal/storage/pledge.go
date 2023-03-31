package storage

import "gorm.io/gorm"

type Pledge struct {
	gorm.Model
	Type      string `gorm:"size:100;" json:"type"`
	Title     string `gorm:"size:100;" json:"title"`
	CarBrand  string `gorm:"size:100" json:"carBrand"`
	CarModel  string `gorm:"size:100" json:"carModel"`
	YearIssue string `gorm:"size:100" json:"yearIssue"`
	Condition string `gorm:"size:100" json:"condition"`
	Mileage   string `gorm:"size:100" json:"mileage"`
	VINCode   string `gorm:"size:100" json:"vinCode"`
	ClientID  uint   `json:"clientID"`
	Client    Client `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"client"`
}

func (p *Pledge) Save(db *gorm.DB) (*Pledge, error) {
	err := db.Debug().Create(&p).Error
	if err != nil {
		return &Pledge{}, err
	}

	return p, nil
}

func (p *Pledge) All(db *gorm.DB, id uint) (*[]Pledge, error) {
	var pledges []Pledge
	err := db.Debug().Model(&Pledge{}).Limit(100).Where("client_id = ?", id).Find(&pledges).Error
	if err != nil {
		return nil, err
	}

	return &pledges, nil
}

func (p *Pledge) Get(db *gorm.DB, id uint) (*Pledge, error) {
	err := db.Debug().Model(&Pledge{}).Limit(100).Where("id = ?", id).Take(&p).Error
	if err != nil {
		return nil, err
	}

	return p, nil
}
