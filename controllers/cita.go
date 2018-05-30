package controllers

import(
	"github.com/astaxie/beego"
	"github.com/Julianrt/beego/veterinaria/models"
	
	"strconv"
)

type CitaController struct{
	beego.Controller
}

func (this *CitaController) Get() {
	citas := models.GetAllCitasJoins()
	this.Data["c"] = citas
	this.TplName = "cita/cita.html"
}


//*******************************************************
type AddCitaController struct{
	beego.Controller
}

func (this *AddCitaController) Get(){
	id,_ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	mascotas := models.GetAllMascotas()
	clientes := models.GetAllClientes()
	servicio := models.GetServicio(id)
	
	this.Data ["m"] = mascotas
	this.Data ["c"] = clientes
	this.Data ["s"] = servicio
	
	this.TplName = "cita/cita_save.html"
}

func (this *AddCitaController) Post() {
	servicio_id,_ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	fecha := this.GetString("fecha")
	hora := this.GetString("hora")
	cliente_id,_ := strconv.Atoi(this.GetString("id_cliente"))
	mascota_id,_ := strconv.Atoi(this.GetString("id_mascota"))
	
	cita := models.Cita{
		Fecha		:fecha,
		Hora		:hora,
		ClienteID	:cliente_id,
		MascotaID	:mascota_id,
		ServicioID	:servicio_id,
	}
	
	models.AddCita(&cita)
	
	this.Ctx.Redirect(302,"/")
}