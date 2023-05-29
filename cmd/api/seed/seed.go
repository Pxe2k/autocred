package seed

import (
	"autocredit/cmd/api/internal/storage"
	"gorm.io/gorm"
	"log"
)

func Load(db *gorm.DB) {
	bank := storage.Bank{
		Title: "МФО Orbis Finance",
	}
	err := db.Create(&bank).Error
	if err != nil {
		log.Fatal(err)
	}

	bankProducts := []storage.BankProduct{
		{Title: "С пробегом_Аннуитет", BankID: bank.ID},
		{Title: "С пробегом_Дифф_6М", BankID: bank.ID},
		{Title: "С пробегом_Дифф_12М", BankID: bank.ID},
		{Title: "С пробегом_Дифф_24М", BankID: bank.ID},
		{Title: "С пробегом_Дифф_36М", BankID: bank.ID},
		{Title: "С пробегом_Дифф_48М", BankID: bank.ID},
		{Title: "С пробегом_Дифф_60М", BankID: bank.ID},
		{Title: "С пробегом_Дифф_72М", BankID: bank.ID},
		{Title: "С пробегом_Дифф_84М", BankID: bank.ID},
		{Title: "С пробегом_Аннуитет_ПВ10%", BankID: bank.ID},
		{Title: "С пробегом_Дифф_6М_ПВ10%", BankID: bank.ID},
		{Title: "С пробегом_Дифф_12М_ПВ10%", BankID: bank.ID},
		{Title: "С пробегом_Дифф_24М_ПВ10%", BankID: bank.ID},
		{Title: "С пробегом_Дифф_36М_ПВ10%", BankID: bank.ID},
		{Title: "С пробегом_Дифф_48М_ПВ10%", BankID: bank.ID},
		{Title: "С пробегом_Дифф_60М_ПВ10%", BankID: bank.ID},
		{Title: "С пробегом_Дифф_72М_ПВ10%", BankID: bank.ID},
		{Title: "С пробегом_Дифф_84М_ПВ10%", BankID: bank.ID},
		{Title: "Новое_Аннуитет", BankID: bank.ID},
		{Title: "Новое_Дифф_6М", BankID: bank.ID},
		{Title: "Новое_Дифф_12М", BankID: bank.ID},
		{Title: "Новое_Дифф_24М", BankID: bank.ID},
		{Title: "Новое_Дифф_36М", BankID: bank.ID},
		{Title: "Новое_Дифф_48М", BankID: bank.ID},
		{Title: "Новое_Дифф_60М", BankID: bank.ID},
		{Title: "Новое_Дифф_72М", BankID: bank.ID},
		{Title: "Новое_Дифф_84М", BankID: bank.ID},
		{Title: "Новое_Аннуитет_ПВ10%", BankID: bank.ID},
		{Title: "Новое_Дифф_6М_ПВ10%", BankID: bank.ID},
		{Title: "Новое_Дифф_12М_ПВ10%", BankID: bank.ID},
		{Title: "Новое_Дифф_24М_ПВ10%", BankID: bank.ID},
		{Title: "Новое_Дифф_36М_ПВ10%", BankID: bank.ID},
		{Title: "Новое_Дифф_48М_ПВ10%", BankID: bank.ID},
		{Title: "Новое_Дифф_60М_ПВ10%", BankID: bank.ID},
		{Title: "Новое_Дифф_72М_ПВ10%", BankID: bank.ID},
		{Title: "Новое_Дифф_84М_ПВ10%", BankID: bank.ID},
		{Title: "Новое_Аннуитет_Субсидия", BankID: bank.ID},
	}
	for _, product := range bankProducts {
		err = db.Create(&product).Error
		if err != nil {
			log.Fatal(err)
		}
	}

	lifeInsurances := []storage.LifeInsurance{
		{Title: "ORBIS Life 30", Percent: 1.10, BankID: bank.ID},
		{Title: "ORBIS Life 40", Percent: 1.40, BankID: bank.ID},
		{Title: "ORBIS Life 60", Percent: 3.00, BankID: bank.ID},
	}
	for _, insurance := range lifeInsurances {
		err = db.Create(&insurance).Error
		if err != nil {
			log.Fatal(err)
		}
	}

	kaskos := []storage.Kasko{
		{Title: "Orbis Small", Percent: 1.50, BankID: bank.ID},
		{Title: "Orbis Medium", Percent: 2.15, BankID: bank.ID},
		{Title: "Orbis Large", Percent: 2.90, BankID: bank.ID},
		{Title: "VIP", Percent: 3.00, BankID: bank.ID},
		{Title: "VIP Plus", Percent: 3.40, BankID: bank.ID},
	}
	for _, kasko := range kaskos {
		err = db.Create(&kasko).Error
		if err != nil {
			log.Fatal(err)
		}
	}

	roadHelp := storage.RoadHelp{
		Title:  "ORBIS ASSISTANCE",
		Price:  150000,
		BankID: bank.ID,
	}
	err = db.Create(&roadHelp).Error
	if err != nil {
		log.Fatal(err)
	}
}
