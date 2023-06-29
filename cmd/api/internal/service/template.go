package service

import (
	"autocredit/cmd/api/helpers"
	"autocredit/cmd/api/helpers/requests"
	"autocredit/cmd/api/internal/storage"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"html/template"
	"os"
	"strconv"
	"time"
)

type RequestPdf struct {
	body string
}

func NewRequestPdf(body string) *RequestPdf {
	return &RequestPdf{
		body: body,
	}
}

func GeneratePdf(db *gorm.DB, body []byte, id uint) ([]storage.BankProcessingDocument, error) {
	r := NewRequestPdf("")

	requestData := requests.GenerateDocumentRequestData{}
	err := json.Unmarshal(body, &requestData)
	if err != nil {
		return nil, err
	}

	BCCTemplateFile := "templates/resultMedia/documentTemplates/BCCDataProcessing.html"
	EUTemplateFile := "templates/resultMedia/documentTemplates/EUDataProcessing.html"
	ShinhanTemplateFile := "templates/resultMedia/documentTemplates/ShinhanDataProcessing.html"

	client := storage.IndividualClient{}
	documentData := requests.ProcessingTemplateData{}
	clientGotten, err := client.Get(db, id)
	if err != nil {
		return nil, err
	}

	documentData.FIO = clientGotten.MiddleName + " " + clientGotten.FirstName + " " + clientGotten.LastName
	documentData.Phone = clientGotten.Phone
	documentData.CurrentDate = helpers.CurrentDateString()
	documentData.Place = clientGotten.User.AutoDealer.Address
	documentData.BirthPlace = clientGotten.Document.PlaceOfBirth
	documentData.BirthDate = clientGotten.BirthDate
	documentData.Address = clientGotten.ResidentialAddress.Address
	documentData.DocumentNumber = clientGotten.Document.Number
	documentData.DocumentIssuingAuthority = clientGotten.Document.IssuingAuthority
	documentData.DocumentIssueDate = clientGotten.Document.DocumentIssueDate
	documentData.IIN = clientGotten.Document.IIN
	documentData.OTP = " "

	var fileName string
	bankProcessingDocuments := []storage.BankProcessingDocument{}

	for _, bankTitle := range requestData.Banks {
		if bankTitle.ID == 1 {
			err = r.ParseTemplate(fmt.Sprint(BCCTemplateFile), documentData)
			if err != nil {
				return nil, err
			}

			fileName = "bcc-data-processing" + strconv.Itoa(int(clientGotten.ID)) + "_" + helpers.CurrentDateString()
			bankProcessingDocuments = append(bankProcessingDocuments, storage.BankProcessingDocument{Title: "Банк Центр Кредит", File: "storage/" + fileName + ".pdf", ApplicationID: requestData.ApplicationID, Image: "banks/bcc.png", BankID: 1})

			err = r.ConvertHTMLtoPdf("storage/" + fileName + ".pdf")
			if err != nil {
				return nil, err
			}
		} else if bankTitle.ID == 2 {
			err = r.ParseTemplate(fmt.Sprint(EUTemplateFile), documentData)
			if err != nil {
				return nil, err
			}

			fileName = "eu-data-processing" + strconv.Itoa(int(clientGotten.ID)) + "_" + helpers.CurrentDateString()
			bankProcessingDocuments = append(bankProcessingDocuments, storage.BankProcessingDocument{Title: "Евразийский Банк", File: "storage/" + fileName + ".pdf", ApplicationID: requestData.ApplicationID, Image: "banks/eu.png", BankID: 2})

			err = r.ConvertHTMLtoPdf("storage/" + fileName + ".pdf")
			if err != nil {
				return nil, err
			}
		} else if bankTitle.ID == 3 {
			err = r.ParseTemplate(fmt.Sprint(ShinhanTemplateFile), documentData)
			if err != nil {
				return nil, err
			}

			fileName = "shinhan-data-processing" + strconv.Itoa(int(clientGotten.ID)) + "_" + helpers.CurrentDateString()
			bankProcessingDocuments = append(bankProcessingDocuments, storage.BankProcessingDocument{Title: "Шинхан Банк", File: "storage/" + fileName + ".pdf", ApplicationID: requestData.ApplicationID, Image: "banks/shinhan.png", BankID: 3})

			err = r.ConvertHTMLtoPdf("storage/" + fileName + ".pdf")
			if err != nil {
				return nil, err
			}
		}
	}

	_, err = UploadFilesToUser(db, bankProcessingDocuments)
	if err != nil {
		return nil, err
	}
	return bankProcessingDocuments, nil
}

func ConfirmPdf(db *gorm.DB, body []byte, id uint) ([]storage.BankProcessingDocument, error) {
	r := NewRequestPdf("")

	requestData := requests.GenerateDocumentRequestData{}
	err := json.Unmarshal(body, &requestData)
	if err != nil {
		return nil, err
	}

	BCCTemplateFile := "templates/resultMedia/documentTemplates/BCCDataProcessing.html"
	EUTemplateFile := "templates/resultMedia/documentTemplates/EUDataProcessing.html"
	ShinhanTemplateFile := "templates/resultMedia/documentTemplates/ShinhanDataProcessing.html"

	client := storage.IndividualClient{}
	documentData := requests.ProcessingTemplateData{}
	clientGotten, err := client.Get(db, id)
	if err != nil {
		return nil, err
	}

	documentData.FIO = clientGotten.MiddleName + " " + clientGotten.FirstName + " " + clientGotten.LastName
	documentData.Phone = clientGotten.Phone
	documentData.CurrentDate = helpers.CurrentDateString()
	documentData.Place = clientGotten.User.AutoDealer.Address
	documentData.BirthPlace = clientGotten.Document.PlaceOfBirth
	documentData.BirthDate = clientGotten.BirthDate
	documentData.Address = clientGotten.ResidentialAddress.Address
	documentData.DocumentNumber = clientGotten.Document.Number
	documentData.DocumentIssuingAuthority = clientGotten.Document.IssuingAuthority
	documentData.DocumentIssueDate = clientGotten.Document.DocumentIssueDate
	documentData.OTP = requestData.OTP

	// TODO вынести Redis отдельно
	val, err := helpers.Redis.Get(helpers.Ctx, clientGotten.Phone).Result()
	if err == redis.Nil {
		fmt.Println("key does not exist")
		return nil, err
	} else if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if val != documentData.OTP {
		return nil, errors.New("code != value")
	}

	var fileName string
	var bankProcessingDocuments []storage.BankProcessingDocument

	for _, bankTitle := range requestData.Banks {
		if bankTitle.ID == 1 {
			err = r.ParseTemplate(fmt.Sprint(BCCTemplateFile), documentData)
			if err != nil {
				return nil, err
			}

			fileName = "bcc-data-processing" + strconv.Itoa(int(clientGotten.ID)) + "_" + helpers.CurrentDateString()
			bankProcessingDocuments = append(bankProcessingDocuments, storage.BankProcessingDocument{Title: "Банк Центр Кредит", File: "storage/" + fileName + ".pdf", ApplicationID: requestData.ApplicationID, Image: "storage/banks/bcc.png", BankID: 1})

			err = r.ConvertHTMLtoPdf("storage/" + fileName + ".pdf")
			if err != nil {
				return nil, err
			}
		} else if bankTitle.ID == 2 {
			err = r.ParseTemplate(fmt.Sprint(EUTemplateFile), documentData)
			if err != nil {
				return nil, err
			}

			fileName = "eu-data-processing" + strconv.Itoa(int(clientGotten.ID)) + "_" + helpers.CurrentDateString()
			bankProcessingDocuments = append(bankProcessingDocuments, storage.BankProcessingDocument{Title: "Евразийский Банк", File: "storage/" + fileName + ".pdf", ApplicationID: requestData.ApplicationID, Image: "storage/banks/eu.png", BankID: 2})

			err = r.ConvertHTMLtoPdf("storage/" + fileName + ".pdf")
			if err != nil {
				return nil, err
			}
		} else if bankTitle.ID == 3 {
			err = r.ParseTemplate(fmt.Sprint(ShinhanTemplateFile), documentData)
			if err != nil {
				return nil, err
			}

			fileName = "shinhan-data-processing" + strconv.Itoa(int(clientGotten.ID)) + "_" + helpers.CurrentDateString()
			bankProcessingDocuments = append(bankProcessingDocuments, storage.BankProcessingDocument{Title: "Шинхан Банк", File: "storage/" + fileName + ".pdf", ApplicationID: requestData.ApplicationID, Image: "storage/banks/shinhan.png", BankID: 3})

			err = r.ConvertHTMLtoPdf("storage/" + fileName + ".pdf")
			if err != nil {
				return nil, err
			}
		}
	}

	return bankProcessingDocuments, nil
}

func (r *RequestPdf) ParseTemplate(templateFileName string, client interface{}) error {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, client); err != nil {
		return err
	}
	r.body = buf.String()
	return nil
}

func (r *RequestPdf) ConvertHTMLtoPdf(pdfPath string) error {
	t := time.Now().Unix()
	// write whole the body

	if _, err := os.Stat("templates/resultMedia/cloneMedia/"); os.IsNotExist(err) {
		errDir := os.Mkdir("templates/resultMedia/cloneMedia/", 0777)
		if errDir != nil {
			fmt.Println(errDir)
			return errDir
		}
	}
	err1 := os.WriteFile("templates/resultMedia/cloneMedia/"+strconv.FormatInt(int64(t), 10)+".html", []byte(r.body), 0644)
	if err1 != nil {
		fmt.Println(err1)
		return err1
	}

	f, err := os.Open("templates/resultMedia/cloneMedia/" + strconv.FormatInt(int64(t), 10) + ".html")
	if f != nil {
		defer f.Close()
	}
	if err != nil {
		fmt.Println(err)
		return err
	}

	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		fmt.Println(err)
		return err
	}

	pdfg.AddPage(wkhtmltopdf.NewPageReader(f))

	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)

	pdfg.Dpi.Set(300)

	err = pdfg.Create()
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = pdfg.WriteFile(pdfPath)
	if err != nil {
		fmt.Println(err)
		return err
	}

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return err
	}

	defer os.RemoveAll(dir + "/templates/resultMedia/cloneMedia")

	return nil
}

func UploadFilesToUser(db *gorm.DB, mediaSeeded []storage.BankProcessingDocument) ([]storage.BankProcessingDocument, error) {
	media := storage.BankProcessingDocument{}

	mediaCreated, err := media.MultipleSave(db, mediaSeeded)
	if err != nil {
		return []storage.BankProcessingDocument{}, err
	}

	return mediaCreated, nil
}

func GetUserMedia(db *gorm.DB, uid uint) ([]storage.BankProcessingDocument, error) {
	bankProcessingDocument := storage.BankProcessingDocument{}
	bankProcessingDocuments, err := bankProcessingDocument.All(db, uid)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return bankProcessingDocuments, nil
}
