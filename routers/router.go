package routers

import (
	"github.com/Julianrt/beego/veterinaria/controllers"
	"github.com/Julianrt/beego/veterinaria/controllers/servicio"
	"github.com/Julianrt/beego/veterinaria/controllers/cliente"
	"github.com/Julianrt/beego/veterinaria/controllers/raza"
	"github.com/Julianrt/beego/veterinaria/controllers/mascota"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/servicio/save", &servicio.SaveServicioController{})
	beego.Router("/servicio/edit/:id([0-9]+)", &servicio.EditServicioController{})
	beego.Router("/servicio/delete", &servicio.DeleteServicioController{})
	beego.Router("/servicio/register/:id([0-9]+)", &controllers.RegisterServicioController{})
	beego.Router("/servicio/historial", &controllers.HistorialServicioController{})
	
	beego.Router("/cliente", &cliente.ClienteController{})
	beego.Router("/cliente/edit/:id([0-9]+)", &cliente.EditClienteController{})
	beego.Router("/cliente/delete", &cliente.DeleteClienteController{})
	
	beego.Router("/raza", &raza.RazaController{})
	beego.Router("/raza/edit/:id([0-9]+)", &raza.EditRazaController{})
	beego.Router("/raza/delete", &raza.DeleteRazaController{})
	
	beego.Router("/mascota", &mascota.MascotaController{})
	beego.Router("/mascota/edit/:id([0-9]+)", &mascota.EditMascotaController{})
	beego.Router("/mascota/delete", &mascota.DeleteMascotaController{})
	
	beego.Router("/cita",&controllers.CitaController{})
	beego.Router("/cita/save/:id([0-9]+)", &controllers.AddCitaController{})
}
