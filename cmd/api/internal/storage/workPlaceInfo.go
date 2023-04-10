package storage

import "gorm.io/gorm"

type WorkPlaceInfo struct {
	gorm.Model
	OrganizationName  string           `gorm:"size:100;" json:"organizationName"` // Название организанции
	WorkPlace         string           `gorm:"size:100;" json:"workPlaceType"`    // Тип места работы
	WorkingActivityID uint             `json:"activityTypeID"`                    // Тип деятельности
	JobTitleID        uint             `json:"jobTitleID"`                        // Должность
	MonthlyIncome     int              `json:"monthlyIncome"`                     // Доход
	Address           string           `gorm:"size:100;" json:"address"`          // Адрес
	Experience        string           `gorm:"size:100;" json:"experience"`       // Стаж работы в организации (мес)
	EmploymentRate    string           `gorm:"size:100;" json:"employmentRate"`   // Степень занятости
	EmploymentDate    string           `gorm:"size:100;" json:"employmentDate"`   // Дата трудоустройства
	DateNextSalary    string           `gorm:"size:100;" json:"dateNextSalary"`   // Дата следующей з/п
	OrganizationPhone string           `gorm:"size:100;" json:"organizationPhone"`
	WorkingActivity   *WorkingActivity `json:"workingActivity,omitempty"`
	JobTitle          *JobTitle        `json:"jobTitle,omitempty"`
	ClientID          uint
}
