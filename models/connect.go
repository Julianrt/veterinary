package models

import(
    "log"
    
    "github.com/jinzhu/gorm"
    _"github.com/jinzhu/gorm/dialects/mysql"
)

func InitializeDatabase() *gorm.DB{
    connection, err := gorm.Open( "mysql", createString() )
    if err != nil {
        log.Fatal()
        return nil
    }
    log.Println("la conexión con la base de datos fue exitosa")
    return connection
}

func createString()string{
    return "root:coco@/veterinaria"
}

func CloseConnection(db *gorm.DB){
    db.Close()
    log.Println("la conexión con la base de datos fue cerrada")
}
