package routers

import (
	"github.com/bitleo/bitleo.top/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/pdfconverttoword", &controllers.PdfController{})
	beego.Router("/pdfconverttoword/convert", &controllers.PdfController{},"get,post:Convert")
	beego.Router("/pdfconverttoword/getfile", &controllers.PdfController{},"get:GetFile")
}
