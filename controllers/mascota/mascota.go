package mascota

import (
	"github.com/astaxie/beego"
	"github.com/Julianrt/beego/veterinaria/models"
	"strconv"
	"log"
)

type MascotaController struct {
	beego.Controller
}

func (this *MascotaController) Get() {
	datos := models.GetAllMascotasJoins()
	this.Data["d"] = datos
	this.TplName = "mascota/mascota.html"
}

func (this *MascotaController) Post() {
	nombre := this.GetString("name")
	raza,_ := strconv.Atoi(this.GetString("raza"))
	edad,_ := strconv.Atoi(this.GetString("edad"))
	cliente,_ := strconv.Atoi(this.GetString("cliente_dueno"))
	
	mascota := models.Mascota{
		Nombre_Mascota:	nombre,
		RazaID:			raza,
		Edad:			edad,
		ClienteID:		cliente,
	}
	
	models.AddMascota(mascota)
	this.Ctx.Redirect(302,"/mascota")
}

type EditMascotaController struct {
	beego.Controller
}

func (this *EditMascotaController) Get() {
	mascota_id,_ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	mascota := models.GetMascota(mascota_id)
	datos := models.GetAllMascotasJoins()
	for i,_ := range(datos.Razas) {
		if datos.Razas[i].Model.ID == uint(mascota.RazaID) {
			datos.Razas[i].Selected = true
		}
		log.Println(datos.Razas[i])
	}
	for i,_ := range(datos.Clientes) {
		if datos.Clientes[i].Model.ID == uint(mascota.ClienteID) {
			datos.Clientes[i].Selected = true
		}
		log.Println(datos.Razas[i])
	}
	this.Data["d"] = datos
	this.Data["m"] = mascota
	this.TplName = "mascota/mascota_edit.html"
}

func (this *EditMascotaController) Post() {
	
	id,_ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	log.Println(id)
	nombre := this.GetString("name")
	raza,_ := strconv.Atoi(this.GetString("raza"))
	edad,_ := strconv.Atoi(this.GetString("edad"))
	cliente,_ := strconv.Atoi(this.GetString("cliente_dueno"))
	
	models.EditMascota(id,raza,edad,cliente,nombre)
	this.Ctx.Redirect(302,"/mascota")
}

type DeleteMascotaController struct {
	beego.Controller
}

func (this *DeleteMascotaController) Get() {
	mascotas := models.GetAllMascotas()
	this.Data["m"] = mascotas
	this.TplName = "mascota/mascota_delete.html"
}

func (this *DeleteMascotaController) Post() {
	id,_ := strconv.Atoi(this.GetString("idd"))
	models.DeleteMascota(id)
	this.Ctx.Redirect(302,"/mascota")
}
