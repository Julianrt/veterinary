package servicio

import(
	"github.com/Julianrt/beego/veterinaria/models"
	"github.com/astaxie/beego"
	"strconv"
)

type DeleteServicioController struct {
	beego.Controller
}

func (this *DeleteServicioController) Get () {
	servicios := models.GetAllServicios()
	this.Data["s"] = servicios
	this.TplName = "servicios/service_delete.html"
}

func (this *DeleteServicioController) Post() {
	id,_ := strconv.Atoi(this.GetString("idd"))
	models.DeleteServicio(id)
	this.Ctx.Redirect(302,"/")
}
