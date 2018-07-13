package models

import (
	"log"
	"fmt"
)

func init() {
	//InitializeDatabase()
}

type Convert struct{
	Clientes []Cliente
	Razas []Raza
}

func GetAllMascotasJoins() DataMascota {
	
	
	data := DataMascota{}
	razas := []Raza{}
	mascotashow := []MascotaShow{}
	clientes := []Cliente{}
	//clientesB := []ClienteB{}
	
	db.Find(&razas)
	db.Find(&clientes)
	
	//var i = Convert{Clientes:clientes}
	
	//for _, v := range []interface{}{i.Clientes} {
	//	log.Println(v.Nombre)
	//}
	
	//clientesB = ClienteB(clientes)
	
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

func GetMascotasCliente(cliente int) []Mascota{
	var mascotas []Mascota
	db.Where("cliente_id = ?",cliente).Find(&mascotas)
	
	return mascotas
}

func GetMascotaSelect(cliente int) string {
	
	var selectMascota string = `<option value="" disabled selected>escoge una mascotas</option>`
	
	rows,err := db.Table("mascota").Select("id, nombre_mascota").Where("cliente_id = ?",cliente).Rows()
	if err != nil {
		log.Println(err)
	}
	for rows.Next() {
		var id int
		var nombre string

		err = rows.Scan(&id, &nombre)
		if err != nil {
			log.Println(err)
		}
		
		selectMascota += fmt.Sprintf("<option value=\"%d\">%s</option>",id,nombre)
		
	}
	
	return selectMascota
	
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
	log.Println(m)
}

func DeleteMascota(aidi int) {
	db.Model(&Mascota{}).Unscoped().Where("id = ?",aidi).Delete(Mascota{})
}
