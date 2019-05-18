package controllers

import "github.com/astaxie/beego"

type PdfController struct {
	beego.Controller
}

func (c *PdfController) Get() {
	c.TplName = "pdfconverttoword.html"
}