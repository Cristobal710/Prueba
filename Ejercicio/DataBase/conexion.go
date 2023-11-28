package DataBase

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var StringDataBase = "host=localhost user=Quito password=Quito123 dbname=gorm port=5432"

var DB *gorm.DB

func Conexion() {
	var err error
	DB, err = gorm.Open(postgres.Open(StringDataBase), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("DataBase Conectado")
}
