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

func GeneratePdf(db *gorm.DB, body []byte, id uint) (*storage.Media, error) {
	r := NewRequestPdf("")

	requestData := requests.GenerateDocumentRequestData{}
	err := json.Unmarshal(body, &requestData)
	if err != nil {
		return &storage.Media{}, err
	}

	BCCTemplateFile := "templates/resultMedia/documentTemplates/BCCDataProcessing.html"

	client := storage.IndividualClient{}
	documentData := requests.BCCTemplateData{}
	clientGotten, err := client.Get(db, id)
	if err != nil {
		return nil, err
	}

	documentData.FIO = clientGotten.MiddleName + " " + clientGotten.FirstName + " " + clientGotten.LastName
	documentData.Phone = clientGotten.Phone
	documentData.CurrentDate = helpers.CurrentDateString()
	documentData.Place = clientGotten.User.AutoDealer.Address

	var fileName string

	for _, bankTitle := range requestData.Banks {
		if bankTitle.Title == "BCC" {
			err = r.ParseTemplate(fmt.Sprint(BCCTemplateFile), documentData)
			fileName = "bcc-data-processing" + strconv.Itoa(int(clientGotten.ID)) + "_" + documentData.CurrentDate
			if err != nil {
				return &storage.Media{}, err
			}
		} else if bankTitle.Title == "EU" {
			err = r.ParseTemplate(fmt.Sprint(BCCTemplateFile), documentData)
			fileName = "bcc-data-processing" + strconv.Itoa(int(clientGotten.ID)) + "_" + documentData.CurrentDate
			if err != nil {
				return &storage.Media{}, err
			}
		} else if bankTitle.Title == "Shinhan" {
			fileName = "bcc-data-processing" + strconv.Itoa(int(clientGotten.ID)) + "_" + documentData.CurrentDate
			err = r.ParseTemplate(fmt.Sprint(BCCTemplateFile), documentData)
			if err != nil {
				return &storage.Media{}, err
			}
		}
	}

	outputPath := "storage/" + fileName + ".pdf"
	err = r.ConvertHTMLtoPdf(outputPath)
	if err != nil {
		return &storage.Media{}, err
	}

	//var mediaDataSeeded = []storage.Media{}

	mediaCreated, err := UploadFileToUser(db, uint32(id), outputPath, fileName)
	return mediaCreated, nil
}

func ConfirmPdf(db *gorm.DB, body []byte, id uint) (*storage.Media, error) {
	r := NewRequestPdf("")
	var result map[string]interface{}

	err := json.Unmarshal(body, &result)
	if err != nil {
		return &storage.Media{}, err
	}

	BCCTemplateFile := "templates/resultMedia/documentTemplates/BCCDataProcessing.html"

	client := storage.IndividualClient{}
	documentData := requests.BCCTemplateData{}
	clientGotten, err := client.Get(db, id)
	if err != nil {
		return nil, err
	}

	documentData.FIO = clientGotten.MiddleName + " " + clientGotten.FirstName + " " + clientGotten.LastName
	documentData.Phone = clientGotten.Phone
	documentData.CurrentDate = helpers.CurrentDateString()
	documentData.OTP = fmt.Sprint(result["OTP"])
	documentData.Place = clientGotten.User.AutoDealer.Address
	fmt.Println(documentData.Place)

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

	err = r.ParseTemplate(fmt.Sprint(BCCTemplateFile), documentData)
	if err != nil {
		return &storage.Media{}, err
	}

	fileName := "bcc-data-processing" + "_" + documentData.CurrentDate

	outputPath := "storage/" + fileName + ".pdf"
	err = r.ConvertHTMLtoPdf(outputPath)
	if err != nil {
		return &storage.Media{}, err
	}

	mediaCreated, err := UploadFileToUser(db, uint32(id), outputPath, fileName)
	return mediaCreated, nil
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

func UploadFileToUser(db *gorm.DB, uid uint32, filePath string, title string) (*storage.Media, error) {
	media := storage.Media{}
	media.File = filePath
	media.Title = title
	media.IndividualClientID = uint(uid)
	mediaCreated, err := media.Save(db)
	if err != nil {
		return &storage.Media{}, err
	}

	return mediaCreated, nil
}
