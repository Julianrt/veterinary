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
	servicio_id,_ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	servicios := models.GetServicio(servicio_id)
	this.Data["s"] = servicios
	this.TplName = "servicios/service_edit.html"
}

func (this *EditServicioController) Post() {
	//id,_ := strconv.Atoi(this.GetString("idd"))
	id,_ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	nombre := this.GetString("name")
	costo,_ := strconv.Atoi(this.GetString("costo"))
	descripcion := this.GetString("descripcion")
	
	models.UpdateServicio(id, costo, nombre, descripcion)
	this.Ctx.Redirect(302,"/")
}
