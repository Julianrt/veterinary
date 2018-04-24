package routers

import (
	"github.com/Julianrt/beego/veterinaria/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.ServiciosController{})
}
