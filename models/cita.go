package models

import (
	"log"
)

func Init() {
	
	
}

func AddCita(cita *Cita) {
	db.Create(cita)
}

func GetAllCitas() []Cita {
	var citas []Cita
	db.Find(&citas)
	return citas
}

func GetAllCitasJoins() []DataCita {
	
	var citas []DataCita
	
	rows,err := db.Table("cita").Select("cita.id, cita.fecha, cita.hora, clientes.nombre, mascota.nombre_mascota, servicios.nombre_servicio").Joins("inner join clientes on clientes.id = cita.cliente_id inner join mascota on mascota.id = cita.mascota_id inner join servicios on servicios.id = cita.servicio_id  ").Rows()
	if err != nil{
		log.Println(err)
	}
	for rows.Next() {
		var id int
		var fecha, hora, cliente, mascota, servicio string
		
		err = rows.Scan(&id, &fecha, &hora, &cliente, &mascota, &servicio)
		if err != nil {
			log.Println(err)
		}
		citas = append(citas, DataCita{
			ID			:id,
			Fecha		:fecha,
			Hora		:hora,
			Cliente		:cliente,
			Mascota		:mascota,
			Servicio	:servicio,
		})
	}
	
	return citas
	
}

func GetCita(aidi int) Cita {
	var cita Cita
	db.First(&cita, aidi)
	return cita
}

func UpdateCita(aidi, cliente_id, mascota_id, servicio_id int, fecha, hora string) {
	db.Model(&Cita{}).Where("id = ?",aidi).Updates(Cita{
		Fecha		:fecha,
		Hora		:hora,
		ClienteID	:cliente_id,
		MascotaID	:mascota_id,
		ServicioID	:servicio_id,
	})
}

func DeleteCita(aidi int){
	db.Model(&Cita{}).Unscoped().Where("id = ?",aidi).Delete(Cita{})
}