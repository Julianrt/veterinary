package raza

import (
	"github.com/astaxie/beego"
	"github.com/Julianrt/beego/veterinaria/models"
)

type RazaController struct {
	beego.Controller
}

func (this *RazaController) Get() {
	razas := models.GetAllRazas()
	this.Data["r"] = razas
	this.TplName = "raza/raza.html"
}

func (this *RazaController) Post() {
	nombre := this.GetString("name")
	raza := models.Raza{
		Nombre_Raza:	nombre,
	}
	models.AddRaza(raza)
	
	this.Ctx.Redirect(302,"/raza")
}
