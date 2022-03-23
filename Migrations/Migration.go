package main

import (
	"PruebaBackendWilliam/Migrations/AutoMigration"
	"PruebaBackendWilliam/Utils/DependencyInjector"
	"fmt"

	"gorm.io/gorm"
)

func main() {

	//Accedemos a la conexión a traves de la inyección de dependencias
	conn, err := DependencyInjector.DbHandler.Connect()

	if err != nil {
		panic(err)
	}

	db, _ := conn.DB()

	defer db.Close()

	GenerateMigration(conn)
}

//GenerateMigration: genera las tablas de los modelos
func GenerateMigration(context *gorm.DB) {

	fmt.Println("\n\n===========Inicia proceso de migración============")

	//Creación de tablas
	fmt.Println("\n\n========== CreateTable ! ==========")
	AutoMigration.CreateTable(context)

	//Carga valores por defecto
	fmt.Println("\n\n========== LoadData ! ==========")
	AutoMigration.LoadDefaultData(context)

	fmt.Println("\n\n========== Proceso terminado exitosamente ! ==========")
}
