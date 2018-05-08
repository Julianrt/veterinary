package routers

import (
	"github.com/Julianrt/beego/veterinaria/controllers"
	"github.com/Julianrt/beego/veterinaria/controllers/servicio"
	"github.com/Julianrt/beego/veterinaria/controllers/cliente"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/servicio/save", &servicio.SaveServicioController{})
	beego.Router("/servicio/edit", &servicio.EditServicioController{})
	beego.Router("/servicio/delete", &servicio.DeleteServicioController{})
	
	beego.Router("/cliente", &cliente.ClienteController{})
	beego.Router("/cliente/edit", &cliente.EditClienteController{})
	beego.Router("/cliente/delete", &cliente.DeleteClienteController{})
	
    
}
