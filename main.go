package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Producto contiene los datos de un artículo
type Producto struct {
	gorm.Model   //ID, CreatedAt, UpdatedAt, DeletedAt
	CodigoBarras string
	Precio       uint
}

// Entidad de Trabajo con todos sus datos oficiales.
type EntidadTrabajo struct {
	gorm.Model
	rif       string
	fechaSolv string
	razonSoc  string
}

func main() {

	// db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=ws_general password=123456")
	db, err := gorm.Open("mysql", "root:@/edcurso?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic("Error de Conexión")
	}
	db.Close()
	fmt.Println("Conexión Exitosa")

	db.CreateTable(&EntidadTrabajo{})
	fmt.Println("Creada la tabla")

	db.Create(&Producto{
		CodigoBarras: "otroCodigoDebarras",
		Precio:       5000,
	})

}
