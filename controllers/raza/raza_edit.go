package raza

import(
	"github.com/astaxie/beego"
	"github.com/Julianrt/beego/veterinaria/models"
	"strconv"
)

type EditRazaController struct{
	beego.Controller
}

func (this *EditRazaController) Get(){
	razas := models.GetAllRazas()
	this.Data["r"] = razas
	this.TplName = "raza/edit_raza.html"
}

func (this *EditRazaController) Post() {
	id,_ := strconv.Atoi(this.GetString("idd"))
	nombre := this.GetString("name")
	
	models.UpdateRaza(id, nombre)
	this.Ctx.Redirect(302,"/raza")
}
