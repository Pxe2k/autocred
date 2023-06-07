package storage

import "gorm.io/gorm"

type BonusInfoBusiness struct {
	gorm.Model
	AmountIncome              int    `json:"amountIncome"`                // Сумма доходов
	AmountExpenses            int    `json:"amountExpenses"`              // Сумма расходов
	FinAnalysis               bool   `json:"finAnalysis"`                 // Фин Анализ
	ThirdPersons              bool   `json:"thirdPerson"`                 // Проведение операции клиентом под руководством третьего лица и/или лиц присутствующих при операции
	NonProfitOrganizations    bool   `json:"nonProfitOrganization"`       // Некомерческие и благотворительные организации, религиозные объедения
	BadHabits                 bool   `json:"badHabits"`                   // По внешнему виду (признаки лица без определенного места жительства, признаки наркомании и (или) алкоголизма) лица осуществляющего
	Executive                 bool   `json:"executive"`                   // Должностно лицо
	Lawsuits                  bool   `json:"lawsuits"`                    // Судебные иски
	OverdueDebts              bool   `json:"overdueDebts"`                // Имеются ли просроченные долги
	UnfulfilledJudgment       bool   `json:"unfulfilledJudgment"`         // Существует невыполненное судебное решение
	TaxArrears                bool   `json:"taxArrears"`                  // Имеется ли задолженность по налогам и сборам
	VAT                       bool   `json:"VAT"`                         // НДС
	VATSeries                 int    `json:"VATSeries"`                   // Серия
	VATNumber                 int    `json:"VATNumber"`                   // Номер
	VATDate                   string `gorm:"size:100" json:"VATDate"`     // Дата
	License                   bool   `json:"license"`                     // Лицензия
	LicenseNumber             int    `json:"licenseNumber"`               // Номер
	LicenseDate               string `gorm:"size:100" json:"licenseDate"` // Дата
	BeneficialOwnerBusinessID uint
}
