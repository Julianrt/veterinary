package models

import(
    "log"
    
    "github.com/jinzhu/gorm"
    _"github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error

func InitializeDatabase() {
    db, err = gorm.Open( "mysql", createString() )
    if err != nil {
        log.Println(err)
    }
    log.Println("la conexión con la base de datos fue exitosa")
}

func createString()string{
    return "root:coco@/veterinaria?charset=utf8&parseTime=True&loc=Local"
}

func CloseConnection(){
    db.Close()
    log.Println("la conexión con la base de datos fue cerrada")
}
