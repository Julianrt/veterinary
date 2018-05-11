package controllers

import (
    "github.com/Julianrt/beego/veterinaria/models"
	"github.com/astaxie/beego"
	"strconv"
	"time"
	"fmt"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
    servicios := models.GetAllServicios()
    this.Data["servicios"] = servicios
	this.TplName = "index.html"
}

type RegisterServicioController struct {
	beego.Controller
}

func (this *RegisterServicioController) Get() {
    servicios := models.GetAllServicioJoin()
    this.Data["s"] = servicios
	this.TplName = "servicios/service_register.html"
}

func (this *RegisterServicioController) Post() {
	serviceCliente := models.Servicios_Cliente{}
	t := time.Now()
	fecha := fmt.Sprintf("%d-%02d-%02d", t.Year(),t.Month(),t.Day())
	hora := fmt.Sprintf("%02d:%02d:%02d", t.Hour(),t.Minute(),t.Second())
	
	servicio,_ := strconv.Atoi(this.GetString("id_servicio"))
	cliente,_ := strconv.Atoi(this.GetString("id_cliente"))
	mascota,_ := strconv.Atoi(this.GetString("id_mascota"))
	
	serviceCliente.ServicioID = servicio
	serviceCliente.ClienteID = cliente
	serviceCliente.MascotaID = mascota
	serviceCliente.Fecha = fecha
	serviceCliente.Hora = hora
	
	models.RegisterServicio(serviceCliente)
	
	this.Ctx.Redirect(302,"/")
}

type HistorialServicioController struct {
	beego.Controller
}

func (this *HistorialServicioController) Get() {
    historial := models.GetHistorial()
    this.Data["h"] = historial
	this.TplName = "historial.html"
}

