package storage

import "gorm.io/gorm"

type WorkPlaceInfo struct {
	gorm.Model
	OrganizationName   string           `gorm:"size:100;" json:"organizationName"` // Название организанции
	WorkPlace          string           `gorm:"size:100;" json:"workPlaceType"`    // Тип места работы
	WorkingActivityID  uint             `json:"activityTypeID"`                    // Тип деятельности
	JobTitle           string           `json:"jobTitle"`                          // Должность
	MonthlyIncome      int              `json:"monthlyIncome"`                     // Доход
	Address            string           `gorm:"size:100;" json:"address"`          // Адрес
	Experience         string           `gorm:"size:100;" json:"experience"`       // Стаж работы в организации (мес)
	EmploymentRate     string           `gorm:"size:100;" json:"employmentRate"`   // Степень занятости
	EmploymentDate     string           `gorm:"size:100;" json:"employmentDate"`   // Дата трудоустройства
	DateNextSalary     string           `gorm:"size:100;" json:"dateNextSalary"`   // Дата следующей з/п
	OrganizationPhone  string           `gorm:"size:100;" json:"organizationPhone"`
	WorkingActivity    *WorkingActivity `json:"workingActivity,omitempty"`
	IndividualClientID uint
}

func (w *WorkPlaceInfo) Update(db *gorm.DB, workPlace *WorkPlaceInfo, clientID uint) error {
	err := db.Debug().Model(&WorkPlaceInfo{}).Where("individual_client_id = ?", clientID).Updates(workPlace).Error
	if err != nil {
		return err
	}

	return nil
}
