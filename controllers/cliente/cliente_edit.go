package cliente

import(
	"github.com/Julianrt/beego/veterinaria/models"
	"github.com/astaxie/beego"
	"strconv"
)

type EditClienteController struct {
	beego.Controller
}

func (this *EditClienteController) Get() {
	clientes := models.GetAllClientes()
	this.Data["c"] = clientes
	this.TplName = "cliente/cliente_edit.html"
}

func (this *EditClienteController) Post() {
	id,_ := strconv.Atoi(this.GetString("idd"))
	nombre := this.GetString("name")
	direccion := this.GetString("address")
	telefono := this.GetString("phone")
	
	models.UpdateCliente(id, nombre, direccion, telefono)
	
	this.Ctx.Redirect(302,"/cliente")
	
}
