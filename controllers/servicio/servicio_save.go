package servicio

import(
    "github.com/Julianrt/beego/veterinaria/models"
	"github.com/astaxie/beego"
)

type SaveServicioController struct {
	beego.Controller
}

func (this *SaveServicioController) Get() {
	servicios := models.GetAllServicios()
	this.Data["s"] = servicios
	this.TplName = "servicios/service_save.html";
}

func (this *SaveServicioController) Post() {
	
}