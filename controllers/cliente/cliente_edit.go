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
	cliente_id,_ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	clientes := models.GetAllClientes()
	cliente := models.GetCliente(cliente_id)
	this.Data["c"] = clientes
	this.Data["cc"] = cliente
	this.TplName = "cliente/cliente_edit.html"
}

func (this *EditClienteController) Post() {
	id,_ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	nombre := this.GetString("name")
	direccion := this.GetString("address")
	telefono := this.GetString("phone")
	
	models.UpdateCliente(id, nombre, direccion, telefono)
	
	this.Ctx.Redirect(302,"/cliente")
}
