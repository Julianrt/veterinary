package cliente

import (
	"github.com/Julianrt/beego/veterinaria/models"
	"github.com/astaxie/beego"
	"strconv"
)


type DeleteClienteController struct{
	beego.Controller
}

func (this *DeleteClienteController) Get() {
	clientes := models.GetAllClientes()
	this.Data["c"] = clientes
	this.TplName = "cliente/cliente_delete.html"
}

func (this *DeleteClienteController) Post() {
	id,_ := strconv.Atoi(this.GetString("idd"))
	models.DeleteCliente(id)
	this.Ctx.Redirect(302,"/cliente")
}
