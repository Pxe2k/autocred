package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"log"
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

func GeneratePdf(db *gorm.DB, uid uint32, body []byte) (*storage.Media, error) {
	r := NewRequestPdf("")
	var result map[string]interface{}

	// Unmarshal or Decode the JSON to the interface.
	json.Unmarshal(body, &result)
	fileName := fmt.Sprint(result["templateName"])
	templateFileName := "result-media/document_templates/" + fileName + ".html"
	data := result["data"]

	err := r.ParseTemplate(fmt.Sprint(templateFileName), data)
	if err != nil {
		return &storage.Media{}, err
	}

	outputPath := "result-media/outputPdf/" + fileName + ".pdf"
	_, err = r.ConvertHTMLtoPdf(outputPath)
	if err != nil {
		return &storage.Media{}, err
	}
	//fileBytes, err := os.ReadFile("outputPdf/" + fileName + ".pdf")

	//mediaCreated, err := UploadFileService(db, uid, fileName, "", fileBytes, fileBytes)
	mediaCreated, err := UploadFileToUser(db, uid, outputPath, fileName)
	return mediaCreated, nil
}

func (r *RequestPdf) ParseTemplate(templateFileName string, data interface{}) error {

	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return err
	}
	r.body = buf.String()
	return nil
}

func (r *RequestPdf) ConvertHTMLtoPdf(pdfPath string) (bool, error) {
	t := time.Now().Unix()
	// write whole the body

	if _, err := os.Stat("result-media/cloneTemplate/"); os.IsNotExist(err) {
		errDir := os.Mkdir("result-media/cloneTemplate/", 0777)
		if errDir != nil {
			log.Fatal(errDir)
		}
	}
	err1 := os.WriteFile("result-media/cloneTemplate/"+strconv.FormatInt(int64(t), 10)+".html", []byte(r.body), 0644)
	if err1 != nil {
		panic(err1)
	}

	f, err := os.Open("result-media/cloneTemplate/" + strconv.FormatInt(int64(t), 10) + ".html")
	if f != nil {
		defer f.Close()
	}
	if err != nil {
		log.Fatal(err)
	}

	wkhtmltopdf.SetPath(os.Getenv("CONVERT_TO_PDF_PATH"))

	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
	}

	pdfg.AddPage(wkhtmltopdf.NewPageReader(f))

	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)

	pdfg.Dpi.Set(300)

	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
	}

	err = pdfg.WriteFile(pdfPath)
	if err != nil {
		log.Fatal(err)
	}

	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	defer os.RemoveAll(dir + "/result-media/cloneTemplate")

	return true, nil
}

func UploadFileToUser(db *gorm.DB, uid uint32, filePath string, title string) (*storage.Media, error) {
	media := storage.Media{}
	media.File = filePath
	media.Title = title
	media.UserId = uint(uid)
	mediaCreated, err := media.SaveMedia(db)
	if err != nil {
		return &storage.Media{}, err
	}

	return mediaCreated, nil
}
