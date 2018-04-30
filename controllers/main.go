package controllers

import (
    "github.com/Julianrt/beego/veterinaria/models"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
    servicios := models.GetAllServicios()
    this.Data["servicios"] = servicios
	this.TplName = "index.html"
}
