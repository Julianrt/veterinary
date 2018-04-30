package servicio

import(
	"github.com/Julianrt/beego/veterinaria/models"
	"github.com/astaxie/beego"
)

type DeleteServicioController struct {
	beego.Controller
}

func (this *DeleteServicioController) Get () {
	servicios := models.GetAllServicios()
	this.Data["s"] = servicios
	this.TplName = "servicios/service_delete.html"
}