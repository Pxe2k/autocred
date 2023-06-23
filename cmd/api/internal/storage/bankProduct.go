package storage

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BankProduct struct {
	gorm.Model
	Title                          string         `gorm:"size:100;" json:"title"`
	WithPercentage                 bool           `json:"withPercentage"`
	PercentageWithProvenIncome     float64        `json:"percentageWithProvenIncome"`
	PercentageWithoutProvenIncome  float64        `json:"percentageWithoutProvenIncome"`
	MaxAmountWithVerifiedIncome    uint           `json:"maxAmountWithVerifiedIncome"`
	MaxAmountWithoutVerifiedIncome uint           `json:"maxAmountWithoutVerifiedIncome"`
	MaxAmount                      uint           `json:"maxAmount"`
	Rate12                         float64        `json:"rate12"`
	Rate24                         float64        `json:"rate24"`
	Rate36                         float64        `json:"rate36"`
	Rate48                         float64        `json:"rate48"`
	Rate60                         float64        `json:"rate60"`
	Rate72                         float64        `json:"rate72"`
	Rate84                         float64        `json:"rate84"`
	BankID                         uint           `json:"bankID"`
	Rate                           datatypes.JSON `json:"rate"`
	Comment                        string         `json:"comment"`
}

func (b *BankProduct) Save(db *gorm.DB) (*BankProduct, error) {
	err := db.Debug().Create(&b).Error
	if err != nil {
		return &BankProduct{}, err
	}

	return b, nil
}

func (b *BankProduct) Update(db *gorm.DB, id int) (*BankProduct, error) {
	err := db.Debug().Model(&BankProduct{}).Where("id = ?", id).Updates(&b).Error
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (b *BankProduct) SoftDelete(db *gorm.DB, id uint) (int64, error) {
	err := db.Debug().Model(&BankProduct{}).Where("id = ?", id).Take(&BankProduct{}).Select(clause.Associations).Delete(&BankProduct{})
	if err != nil {
		return 0, err.Error
	}

	if err.Error != nil {
		return 0, err.Error
	}

	return err.RowsAffected, nil
}

func (b *BankProduct) Get(db *gorm.DB, id uint) (*BankProduct, error) {
	err := db.Debug().Model(&BankProduct{}).Where("id = ?", id).First(&b).Error
	if err != nil {
		return nil, err
	}
	return b, nil
}
