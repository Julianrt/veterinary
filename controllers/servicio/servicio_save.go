package servicio

import(
    "github.com/Julianrt/beego/veterinaria/models"
	"github.com/astaxie/beego"
	"strconv"
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
	
	nombre := this.GetString("name")
	costo,_ := strconv.Atoi(this.GetString("costo"))
	descripcion := this.GetString("descripcion")
	
	servicio := models.Servicio{
		Nombre_Servicio: nombre,
		Costo: costo,
		Descripcion: descripcion,
	}
	
	models.AddServicio(servicio)
	
	this.Ctx.Redirect(302,"/")
}
