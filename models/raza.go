package models

func init() {
	
}

func GetAllRazas() []Raza {
	razas := []Raza{}
	db.Find(&razas)
	return razas
}

func GetRaza(aidi int) Raza {
	var raza Raza
	db.First(&raza, aidi)
	return raza
}

func AddRaza(r Raza) {
	db.Create(&r)
}

func UpdateRaza(aidi int, nombre string) {
	raza := Raza{Nombre_Raza: nombre}
	db.Model(&raza).Where("id = ?",aidi).Updates(raza)
}

func DeleteRaza(aidi int){
	db.Model(&Raza{}).Unscoped().Where("id = ?",aidi).Delete(Raza{})
}
