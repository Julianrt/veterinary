package servicio

import(
	"github.com/Julianrt/beego/veterinaria/models"
	"github.com/astaxie/beego"
)

type EditServicioController struct {
	beego.Controller
}

func (this *EditServicioController) Get() {
	servicios := models.GetAllServicios()
	this.Data["s"] = servicios
	this.TplName = "servicios/service_edit.html"
}
