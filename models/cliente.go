package models

func init(){
    //InitializeDatabase()
}

func GetCliente(aidi int) Cliente {
    cliente := Cliente{}
    db.First(&cliente, aidi)
    
    return cliente
}

func GetAllClientes() []Cliente {
    cliente := []Cliente{}
    db.Find(&cliente)
    return cliente
}

func AddCliente(s Cliente){
	db.Create(&s)
}

func UpdateCliente(aidi int, nombre, direccion, telefono string) {
    var cliente Cliente
    db.Model(&cliente).Where("id = ?",aidi).Updates(Cliente{
        Nombre:   	 nombre,
        Direccion:   direccion,
        Telefono:    telefono,
    })
}

func DeleteCliente(aidi int){
	db.Model(Cliente{}).Unscoped().Where("id = ?",aidi).Delete(Cliente{})
}
