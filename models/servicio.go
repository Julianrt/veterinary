package models

func init(){
    InitializeDatabase()
}

func GetServicio(aidi int) Servicio {
    servicio := Servicio{}
    db.First(&servicio, aidi)
    
    return servicio
}

func GetAllServicios() []Servicio {
    servicios := []Servicio{}
    db.Find(&servicios)
    return servicios
}

func AddServicio(s Servicio){
	db.Create(&s)
}

func UpdateServicio(aidi, costo int, nombre, descripcion string) {
    var servicio Servicio
    db.Model(&servicio).Where("id = ?",aidi).Updates(Servicio{
        Nombre_Servicio:    nombre,
        Costo:              costo,
        Descripcion:        descripcion,
    })
}

func DeleteServicio(aidi int){
	db.Model(Servicio{}).Unscoped().Where("id = ?",aidi).Delete(Servicio{})
}
