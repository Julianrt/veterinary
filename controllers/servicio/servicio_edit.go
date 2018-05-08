package servicio

import(
	"github.com/Julianrt/beego/veterinaria/models"
	"github.com/astaxie/beego"
	"strconv"
)

type EditServicioController struct {
	beego.Controller
}

func (this *EditServicioController) Get() {
	servicios := models.GetAllServicios()
	this.Data["s"] = servicios
	this.TplName = "servicios/service_edit.html"
}

func (this *EditServicioController) Post() {
	id,_ := strconv.Atoi(this.GetString("idd"))
	nombre := this.GetString("name")
	costo,_ := strconv.Atoi(this.GetString("costo"))
	descripcion := this.GetString("descripcion")
	
	models.UpdateServicio(id, costo, nombre, descripcion)
	this.Ctx.Redirect(302,"/")
}
