package mascota

import (
	"github.com/astaxie/beego"
	"github.com/Julianrt/beego/veterinaria/models"
	"strconv"
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
	datos := models.GetAllMascotasJoins()
	this.Data["d"] = datos
	this.TplName = "mascota/mascota_edit.html"
}

func (this *EditMascotaController) Post() {
	
	id,_ := strconv.Atoi(this.GetString("idd"))
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
