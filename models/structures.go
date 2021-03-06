package models

import(
    "github.com/jinzhu/gorm"
    _"github.com/jinzhu/gorm/dialects/mysql"
)

type Raza struct{
  gorm.Model
  Nombre_Raza string
  Selected bool
}

type Mascota struct{
  gorm.Model
  Nombre_Mascota string
  RazaID int
  Edad int
  ClienteID int
}
type MascotaShow struct{
  ID_Mascota int
  Nombre_Mascota string
  Raza string
  Edad int
  Cliente string
  Selected bool
}

type Cliente struct{
  gorm.Model
  Nombre string
  Direccion string
  Telefono string
  Selected bool
}
/*type ClienteB struct{
  gorm.Model
  Nombre string
  Direccion string
  Telefono string
  Selected bool
}*/

type Servicio struct{
  gorm.Model
  Nombre_Servicio string
  Costo int
  Descripcion string
}

type Servicios_Cliente struct{
  gorm.Model
  ServicioID int
  ClienteID int
  MascotaID int
  Fecha string
  Hora string
}
type Historial struct{
  ID_SC int
  Servicio string
  Cliente string
  Mascota string
  Fecha string
  Hora string
}

type Cita struct{
	gorm.Model
	Fecha string
	Hora string
	ClienteID int
	MascotaID int
	ServicioID int
}
type DataCita struct{
	ID int
	Fecha string
	Hora string
	Cliente string
	Mascota string
	Servicio string
}

type Puesto struct{
  ID_Puesto int
  Nombre_Puesto string
}

type Empleado struct{
  ID_Empleado int
  Nombre_Empleado string
  Email string
  Pass string
  Direccion string
  Telefono string
  ID_Puesto int
}

type DataMascota struct {
  Razas []Raza
  Mascotas []MascotaShow
  Clientes []Cliente
}
type DataRegister struct{
  Servicios []Servicio
  Clientes []Cliente
  Mascotas []Mascota
}
