package models

import "log"

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

func GetAllServicioJoin() DataRegister {
	services := []Servicio{}
	clientes := []Cliente{}
	mascotas := []Mascota{}
	data := DataRegister{}
	
	db.Find(&services)
	db.Find(&clientes)
	db.Find(&mascotas)

	data.Servicios = services
	data.Clientes = clientes
	data.Mascotas = mascotas
	
	return data
}

func RegisterServicio(s Servicios_Cliente) {
	db.Create(&s)
}

func GetHistorial() []Historial{
	
	historial := []Historial{}

	tables := "servicios_clientes.id, servicios.nombre_servicio, clientes.nombre, mascota.nombre_mascota, servicios_clientes.fecha, servicios_clientes.hora"
	joins := "inner join servicios on servicios.id = servicios_clientes.servicio_id inner join clientes on clientes.id = servicios_clientes.cliente_id inner join mascota on mascota.id = servicios_clientes.mascota_id"
	rows,err := db.Table("servicios_clientes").Select(tables).Joins(joins).Rows()
	if err != nil {
		log.Println(err)
	}
	for rows.Next() {
		var id int
		var nombre_s, nombre_c, nombre_m, fecha, hora string

		err = rows.Scan(&id, &nombre_s, &nombre_c, &nombre_m, &fecha, &hora)
		if err != nil {
			log.Println(err)
		}
		historial = append(historial, Historial{
			ID_SC		:id,
			Servicio	:nombre_s,
			Cliente		:nombre_c,
			Mascota		:nombre_m,
			Fecha		:fecha,
			Hora		:hora,
		})
	}

	return historial	
}
