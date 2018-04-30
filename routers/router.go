package routers

import (
	"github.com/Julianrt/beego/veterinaria/controllers"
	"github.com/Julianrt/beego/veterinaria/controllers/servicio"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/servicio/save", &servicio.SaveServicioController{})
	beego.Router("/servicio/edit", &servicio.EditServicioController{})
	beego.Router("/servicio/delete", &servicio.DeleteServicioController{})
    
}
