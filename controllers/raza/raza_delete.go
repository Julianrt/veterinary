package raza

import(
	"github.com/astaxie/beego"
	"github.com/Julianrt/beego/veterinaria/models"
	"strconv"
)

type DeleteRazaController struct{
	beego.Controller
}

func (this *DeleteRazaController) Get() {
	razas := models.GetAllRazas()
	this.Data["r"] = razas
	this.TplName = "raza/delete_raza.html"
}

func (this *DeleteRazaController) Post() {
	id,_ := strconv.Atoi(this.GetString("idd"))
	models.DeleteRaza(id)
	this.Ctx.Redirect(302,"/raza")
}
