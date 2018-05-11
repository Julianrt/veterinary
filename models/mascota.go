package models

import (
	"log"
)

func init() {
	//InitializeDatabase()
}

func GetAllMascotasJoins() DataMascota {
	
	data := DataMascota{}
	razas := []Raza{}
	mascotashow := []MascotaShow{}
	clientes := []Cliente{}
	
	db.Find(&razas)
	db.Find(&clientes)
	
	rows,err := db.Table("mascota").Select("mascota.id, mascota.nombre_mascota, razas.nombre_raza, mascota.edad, clientes.nombre").Joins("inner join razas on razas.id = mascota.raza_id inner join clientes on clientes.id = mascota.cliente_id").Rows()
	if err != nil {
		log.Println(err)
	}
	for rows.Next() {
		var id, edad int
		var nombre_m, nombre_r, cli string

		err = rows.Scan(&id, &nombre_m, &nombre_r, &edad, &cli)
		if err != nil {
			log.Println(err)
		}
		mascotashow = append(mascotashow, MascotaShow{
				ID_Mascota : id,
				Nombre_Mascota : nombre_m,
				Raza : nombre_r,
				Edad : edad,
				Cliente : cli,
		})
	}

	data.Razas = razas;
	data.Mascotas = mascotashow
	data.Clientes = clientes
	
	return data
}

func GetAllMascotas() []Mascota {
	var mascotas []Mascota
	db.Find(&mascotas)
	return mascotas
}

func GetMascota(aidi int) Mascota {
	var mascota Mascota
	db.First(&mascota, aidi)
	return mascota
}

func AddMascota(m Mascota) {
	db.Create(&m)
}

func EditMascota(aidi, raza, edad, cliente int, nombre string) {
	m := Mascota{
		Nombre_Mascota:	nombre,
		RazaID:			raza,
		Edad:			edad,
		ClienteID:		cliente,
	}
	db.Model(&m).Where("id = ?",aidi).Updates(m)
}

func DeleteMascota(aidi int) {
	db.Model(&Mascota{}).Unscoped().Where("id = ?",aidi).Delete(Mascota{})
}