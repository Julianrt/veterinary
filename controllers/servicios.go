package controllers

import (
    "github.com/Julianrt/beego/veterinaria/models"
	"github.com/astaxie/beego"
)

type ServiciosController struct {
	beego.Controller
}

func (this *ServiciosController) Get() {
    servicios := models.GetAllServicios()
    this.Data["servicios"] = servicios
	this.TplName = "index.html"
}
