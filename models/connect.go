package models

import(
    "log"
	"fmt"
    
    "github.com/jinzhu/gorm"
    _"github.com/jinzhu/gorm/dialects/mysql"
	_"github.com/jinzhu/gorm/dialects/mssql"
)

var db *gorm.DB
var err error

func InitializeDatabase() {
	/*
	*Conexion para MySQL
	*/
    //db, err = gorm.Open( "mysql", createString() )
	

	/*
	*Conexion para SQL Server
	*/
	var server = "192.168.43.99"
	var port = 1433
	var user = "sa"
	var password = "Carritos1#"
	var database = "veterinary"
	connectionString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s", server, user, password, port, database)
	db, err = gorm.Open("mssql", connectionString)
	
	if err != nil {
        log.Println(err)
	}else{
		log.Println("la conexión con la base de datos fue exitosa")
		fmt.Println("Hola")
	}
    
}

func createString()string{
    return "root:coco@/veterinaria?charset=utf8&parseTime=True&loc=Local"
}

func CloseConnection(){
    db.Close()
    log.Println("la conexión con la base de datos fue cerrada")
}
