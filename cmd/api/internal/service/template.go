package service

import (
	"autocredit/cmd/api/helpers"
	"autocredit/cmd/api/helpers/requests"
	"autocredit/cmd/api/helpers/responses"
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

func GeneratePdf(db *gorm.DB, body []byte, id uint) ([]responses.BankDocumentsCreated, error) {
	r := NewRequestPdf("")

	requestData := requests.GenerateDocumentRequestData{}
	err := json.Unmarshal(body, &requestData)
	if err != nil {
		return nil, err
	}

	BCCTemplateFile := "templates/resultMedia/documentTemplates/BCCDataProcessing.html"
	EUTemplateFile := "templates/resultMedia/documentTemplates/EUDataProcessing.html"

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
	mediaSeeded := []storage.Media{}
	var mediaData []responses.BankDocumentsCreated

	for _, bankTitle := range requestData.Banks {
		if bankTitle.ID == 1 {
			err = r.ParseTemplate(fmt.Sprint(BCCTemplateFile), documentData)
			if err != nil {
				return nil, err
			}

			fileName = "bcc-data-processing" + strconv.Itoa(int(clientGotten.ID)) + "_" + helpers.CurrentDateString()
			mediaSeeded = append(mediaSeeded, storage.Media{Title: fileName, File: "storage/" + fileName + ".pdf", IndividualClientID: id})
			mediaData = append(mediaData, responses.BankDocumentsCreated{ID: 1, Title: "Банк Центр Кредит", File: "storage/" + fileName + ".pdf"})

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
			mediaSeeded = append(mediaSeeded, storage.Media{Title: fileName, File: "storage/" + fileName + ".pdf", IndividualClientID: id})
			mediaData = append(mediaData, responses.BankDocumentsCreated{ID: 2, Title: "Евразийский Банк", File: "storage/" + fileName + ".pdf"})

			err = r.ConvertHTMLtoPdf("storage/" + fileName + ".pdf")
			if err != nil {
				return nil, err
			}
		} else if bankTitle.ID == 3 {
			err = r.ParseTemplate(fmt.Sprint(BCCTemplateFile), documentData)
			if err != nil {
				return nil, err
			}

			fileName = "shinhan-data-processing" + strconv.Itoa(int(clientGotten.ID)) + "_" + helpers.CurrentDateString()
			mediaSeeded = append(mediaSeeded, storage.Media{Title: fileName, File: "storage/" + fileName + ".pdf", IndividualClientID: id})
			mediaData = append(mediaData, responses.BankDocumentsCreated{ID: 3, Title: "Шинхан Банк", File: "storage/" + fileName + ".pdf"})

			err = r.ConvertHTMLtoPdf("storage/" + fileName + ".pdf")
			if err != nil {
				return nil, err
			}
		}
	}

	_, err = UploadFilesToUser(db, mediaSeeded)
	if err != nil {
		return nil, err
	}
	return mediaData, nil
}

func ConfirmPdf(db *gorm.DB, body []byte, id uint) ([]responses.BankDocumentsCreated, error) {
	r := NewRequestPdf("")

	requestData := requests.GenerateDocumentRequestData{}
	err := json.Unmarshal(body, &requestData)
	if err != nil {
		return nil, err
	}

	BCCTemplateFile := "templates/resultMedia/documentTemplates/BCCDataProcessing.html"
	EUTemplateFile := "templates/resultMedia/documentTemplates/EUDataProcessing.html"

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
	var mediaData []responses.BankDocumentsCreated

	for _, bankTitle := range requestData.Banks {
		if bankTitle.ID == 1 {
			err = r.ParseTemplate(fmt.Sprint(BCCTemplateFile), documentData)
			if err != nil {
				return nil, err
			}

			fileName = "bcc-data-processing" + strconv.Itoa(int(clientGotten.ID)) + "_" + helpers.CurrentDateString()
			mediaData = append(mediaData, responses.BankDocumentsCreated{ID: 1, Title: "Банк Центр Кредит", File: "storage/" + fileName + ".pdf"})

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
			mediaData = append(mediaData, responses.BankDocumentsCreated{ID: 2, Title: "Евразийский Банк", File: "storage/" + fileName + ".pdf"})

			err = r.ConvertHTMLtoPdf("storage/" + fileName + ".pdf")
			if err != nil {
				return nil, err
			}
		} else if bankTitle.ID == 3 {
			err = r.ParseTemplate(fmt.Sprint(BCCTemplateFile), documentData)
			if err != nil {
				return nil, err
			}

			fileName = "shinhan-data-processing" + strconv.Itoa(int(clientGotten.ID)) + "_" + helpers.CurrentDateString()
			mediaData = append(mediaData, responses.BankDocumentsCreated{ID: 3, Title: "Шинхан Банк", File: "storage/" + fileName + ".pdf"})

			err = r.ConvertHTMLtoPdf("storage/" + fileName + ".pdf")
			if err != nil {
				return nil, err
			}
		}
	}

	return mediaData, nil
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

func UploadFilesToUser(db *gorm.DB, mediaSeeded []storage.Media) ([]storage.Media, error) {
	media := storage.Media{}

	mediaCreated, err := media.MultipleSave(db, mediaSeeded)
	if err != nil {
		return []storage.Media{}, err
	}

	return mediaCreated, nil
}

func GetUserMedia(db *gorm.DB, uid uint) ([]responses.BankDocumentsCreated, error) {
	var mediaData []responses.BankDocumentsCreated



	mediaData = append(mediaData, responses.BankDocumentsCreated{ID: 1, Title: "Банк Центр Кредит", File: "storage/bcc-data-processing" + strconv.Itoa(int(uid)) + "_" + helpers.CurrentDateString() + ".pdf"})
	mediaData = append(mediaData, responses.BankDocumentsCreated{ID: 2, Title: "Евразийский Банк", File: "storage/eu-data-processing" + strconv.Itoa(int(uid)) + "_" + helpers.CurrentDateString() + ".pdf"})
	mediaData = append(mediaData, responses.BankDocumentsCreated{ID: 3, Title: "Шинхан Банк", File: "storage/shinhan-data-processing" + strconv.Itoa(int(uid)) + "_" + helpers.CurrentDateString() + ".pdf"})

	return mediaData, nil
}
