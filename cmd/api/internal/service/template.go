package service

import (
	"autocredit/cmd/api/helpers"
	"autocredit/cmd/api/internal/storage"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
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
	var result map[string]interface{}

	err := json.Unmarshal(body, &result)
	if err != nil {
		return &storage.Media{}, err
	}
	fileName := fmt.Sprint(result["templateName"])
	templateFileName := "templates/resultMedia/documentTemplates/" + fileName + ".html"
	//data := result["data"]

	client := storage.IndividualClient{}
	clientGotten, err := client.Get(db, id)
	if err != nil {
		return nil, err
	}

	err = r.ParseTemplate(fmt.Sprint(templateFileName), clientGotten)
	if err != nil {
		return &storage.Media{}, err
	}

	docNumber := helpers.RandEmailCode()

	outputPath := "storage/" + fileName + docNumber + ".pdf"
	err = r.ConvertHTMLtoPdf(outputPath)
	if err != nil {
		return &storage.Media{}, err
	}
	//fileBytes, err := os.ReadFile("outputPdf/" + fileName + ".pdf")

	//mediaCreated, err := UploadFileService(db, uid, fileName, "", fileBytes, fileBytes)
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

func GeneratePdfForBanks(db *gorm.DB, body []byte, id uint) (*storage.Media, error) {
	r := NewRequestPdf("")
	var result map[string]interface{}

	err := json.Unmarshal(body, &result)
	if err != nil {
		return &storage.Media{}, err
	}
	fileName := fmt.Sprint(result["templateName"])
	templateFileName := "templates/resultMedia/documentTemplates/" + fileName + ".html"
	//data := result["data"]

	client := storage.IndividualClient{}
	clientGotten, err := client.Get(db, id)
	if err != nil {
		return nil, err
	}

	err = r.ParseTemplate(fmt.Sprint(templateFileName), clientGotten)
	if err != nil {
		return &storage.Media{}, err
	}

	docNumber := helpers.RandEmailCode()

	outputPath := "storage/" + fileName + docNumber + ".pdf"
	err = r.ConvertHTMLtoPdf(outputPath)
	if err != nil {
		return &storage.Media{}, err
	}
	//fileBytes, err := os.ReadFile("outputPdf/" + fileName + ".pdf")

	//mediaCreated, err := UploadFileService(db, uid, fileName, "", fileBytes, fileBytes)
	mediaCreated, err := UploadFileToUser(db, uint32(id), outputPath, fileName)
	return mediaCreated, nil
}
