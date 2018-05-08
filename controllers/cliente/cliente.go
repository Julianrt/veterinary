package cliente

import(
	"github.com/Julianrt/beego/veterinaria/models"
	"github.com/astaxie/beego"
)

type ClienteController struct {
	beego.Controller
}

func (this *ClienteController) Get() {
	clientes := models.GetAllClientes()
	this.Data["c"] = clientes
	this.TplName = "cliente/clientes.html"
}

func (this *ClienteController) Post() {
	nombre := this.GetString("name")
	direccion := this.GetString("address")
	telefono := this.GetString("phone")
	
	cliente := models.Cliente{
		Nombre:		nombre,
		Direccion:	direccion,
		Telefono:	telefono,
	}
	
	models.AddCliente(cliente)
	this.Ctx.Redirect(302,"/cliente")
}
