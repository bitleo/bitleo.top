package utils

import (
	"bytes"
	"mime/multipart"
	"os"
	"io"
	"net/http/cookiejar"
	"net/http"
	"log"
	"io/ioutil"

)

type HttpPdfConvertDoc struct {

}


func (c *HttpPdfConvertDoc) Covert(destName string,srcName string )  {
	bodyBuffer := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuffer)
	fileWriter, _ := bodyWriter.CreateFormFile("input-file", "manual&1558173570.pdf")
	file, _ := os.Open(srcName)
	defer file.Close()
	io.Copy(fileWriter, file)
	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()
	cookieJar, _ := cookiejar.New(nil)
	httpClient :=  http.Client{Jar:cookieJar}
	resp, _ := httpClient.Post("https://www.pdfwordconvert.com/uploaddocument", contentType, bodyBuffer)
	log.Println(resp.Status)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println(string(body))
	bodyBuffer = &bytes.Buffer{}
	bodyWriter = multipart.NewWriter(bodyBuffer)
	contentType = bodyWriter.FormDataContentType()
	resp, _ = httpClient.Post("https://www.pdfwordconvert.com/convertsubmit",contentType,bodyBuffer)
	log.Println(resp.Status)
	defer resp.Body.Close()
	body, _ = ioutil.ReadAll(resp.Body)
	//log.Println(string(body))
	resp, _ = httpClient.Get("https://www.pdfwordconvert.com/getconverteddocument")
	log.Println(resp.Status)
	defer resp.Body.Close()
	body, _ = ioutil.ReadAll(resp.Body)
	if resp.StatusCode==200 {
		f, _ := os.Create(destName)
		f.Write(body)
		f.Close()
	}
}