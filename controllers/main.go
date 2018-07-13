package controllers

import (
    "github.com/Julianrt/beego/veterinaria/models"
	"github.com/astaxie/beego"
	"strconv"
	"time"
	"fmt"
)


/*
*
*/
type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
    servicios := models.GetAllServicios()
    this.Data["servicios"] = servicios
	this.TplName = "index.html"
}

//************************************************************************
/*
*
*/
type RegisterServicioController struct {
	beego.Controller
}

func (this *RegisterServicioController) Get() {
	servicio_id,_ := strconv.Atoi(this.Ctx.Input.Param(":id"))
    servicios := models.GetAllServicioJoin()
    this.Data["s"] = servicios
	this.Data["id"] = servicio_id
	this.TplName = "servicios/service_register.html"
}

func (this *RegisterServicioController) Post() {
	serviceCliente := models.Servicios_Cliente{}
	t := time.Now()
	fecha := fmt.Sprintf("%d-%02d-%02d", t.Year(),t.Month(),t.Day())
	hora := fmt.Sprintf("%02d:%02d:%02d", t.Hour(),t.Minute(),t.Second())
	
	servicio,_ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	cliente,_ := strconv.Atoi(this.GetString("id_cliente"))
	mascota,_ := strconv.Atoi(this.GetString("id_mascota"))
	
	serviceCliente.ServicioID = servicio
	serviceCliente.ClienteID = cliente
	serviceCliente.MascotaID = mascota
	serviceCliente.Fecha = fecha
	serviceCliente.Hora = hora
	
	models.RegisterServicio(serviceCliente)
	
	if this.IsAjax(){
		//servicio,_ := strconv.Atoi(this.Ctx.Input.Param(":id"))
		cliente_id,_ := strconv.Atoi(this.GetString("id"))

		lista_mascotas := models.GetMascotaSelect(cliente_id)

		//mascotas := models.GetMascotasCliente(cliente_id)
		//fmt.Println(mascotas)
		
		fmt.Println(lista_mascotas)
		this.Data["mascotas"] = lista_mascotas
	}
	
	this.Ctx.Redirect(302,"/")
}

//************************************************************************************
/*
*
*/
type HistorialServicioController struct {
	beego.Controller
}

func (this *HistorialServicioController) Get() {
    historial := models.GetHistorial()
    this.Data["h"] = historial
	this.TplName = "historial.html"
}

