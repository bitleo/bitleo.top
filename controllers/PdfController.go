package controllers

import (
	"github.com/astaxie/beego"
	"net/url"
	"github.com/astaxie/beego/logs"
	"path"
	"time"
	"strings"
	"strconv"
	"net/http"
	"github.com/bitleo/bitleo.top/utils"

)

type PdfController struct {
	beego.Controller
}

func (c *PdfController) Get() {
	c.TplName = "pdfconverttoword.html"
}

type ConvertResult struct {
	Code int `json:"name"`
	Message  string `json:"message"`
	Data interface{} `json:"data"`
}

type Jar struct {
	cookies []*http.Cookie
}

func (jar *Jar) SetCookies(u *url.URL, cookies []*http.Cookie) {
	jar.cookies = cookies
}

func (jar *Jar) Cookies(u *url.URL) []*http.Cookie {
	return jar.cookies
}


func (c *PdfController) Convert() {

	//image，这是一个key值，对应的是html中input type-‘file’的name属性值
	files,_:=c.GetFiles("input-file")
	file := files[0]
	//得到文件的名称
	fileName := file.Filename
	ext := path.Ext(fileName)
	now := time.Now().Unix()
	pdfDir:=beego.AppConfig.String("pdf_dir")
	docDir:=beego.AppConfig.String("doc_dir")
	saveName:=strings.TrimSuffix(fileName, ext)+"&"+strconv.FormatInt(now,10)
	logs.Debug(fileName,ext,pdfDir,saveName)
	err := c.SaveToFile("input-file",pdfDir+saveName+ext)
	if err != nil {
		c.Ctx.WriteString( "upload faild" )
	}else{

		inName:=pdfDir+saveName+ext
		outName:=docDir+saveName+".docx"
        httpConvert := &utils.HttpPdfConvertDoc{}
        httpConvert.Covert(outName,inName)

		id:= url.QueryEscape(saveName+".docx")
		result := &ConvertResult{200, "SUCCESS",id}
		c.Data["json"] = result
		c.ServeJSON()


	}



}
func (c *PdfController) GetFile() {
	id := c.GetString("id")
	fileName,err := url.QueryUnescape(id)
	if !(fileName=="" || err != nil) {
		docDir:=beego.AppConfig.String("doc_dir")
		destFileName := docDir+fileName
		logs.Debug("get file name is", fileName,destFileName)
		//第一个参数是文件的地址，第二个参数是下载显示的文件的名称
		c.Ctx.Output.Download(destFileName ,fileName)
	}else{
		c.Redirect("/",404)
	}

}