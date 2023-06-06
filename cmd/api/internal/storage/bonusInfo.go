package storage

import "gorm.io/gorm"

type BonusInfo struct {
	gorm.Model
	AmountIncome                       int     `json:"amountIncome"`                         // Сумма доходов
	AmountExpenses                     int     `json:"amountExpenses"`                       // Сумма расходов
	FinAnalysis                        bool    `json:"finAnalysis"`                          // Фин Анализ
	NetIncome                          int     `json:"netIncome"`                            // Чистый доход
	BankTitle                          *string `gorm:"size:100;" json:"bankTitle,omitempty"` // Название банков
	Returnee                           bool    `json:"returnee"`                             // Статус Оралман
	ThirdPersons                       bool    `json:"thirdPerson"`                          // Проведение операции клиентом под руководством третьего лица и/или лиц присутствующих при операции
	NonProfitOrganizations             bool    `json:"nonProfitOrganization"`                // Некомерческие и благотворительные организации, религиозные объедения
	BadHabits                          bool    `json:"badHabits"`                            // По внешнему виду (признаки лица без определенного места жительства, признаки наркомании и (или) алкоголизма) лица осуществляющего
	IPDLStatus                         bool    `json:"IPDLStatus"`                           // Флаг: Статус ИПДЛ
	CriminalRecord                     bool    `json:"criminalRecord"`                       // Судимость
	Lawsuits                           bool    `json:"lawsuits"`                             // Судебные иски
	OverdueDebts                       bool    `json:"overdueDebts"`                         // Имеются ли просроченные долги
	IncapacitationDecisions            bool    `json:"incapacitationDecisions"`              // Существуют ли или существовали в прошлом решения суда об ограничении дееспособности клиента или об установлении над клиентом опекунства
	UnfulfilledJudgment                bool    `json:"unfulfilledJudgment"`                  // Существует невыполненное судебное решение
	RegisteredPsychiatrist             bool    `json:"registeredPsychiatrist"`               // Состоит на учете у психиатра или нарколога
	RestrictionsRightCreditTransaction bool    `json:"restrictionsRightCreditTransaction"`   // Установлены ли какие-либо ограничения права заключать кредитную сделку (в т.ч. брачным договором)
	Pregnancy                          bool    `json:"pregnancy"`                            // Беременность
	TaxArrears                         bool    `json:"taxArrears"`                           // Имеется ли задолженность по налогам и сборам
	ClientID                           uint
}
