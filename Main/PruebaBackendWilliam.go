package main

import "PruebaBackendWilliam/Infraestructure/DataBase"

func main() {

	dbHandler := DataBase.DBHandler{}
	conn, err := dbHandler.Connect()

	if err != nil {

		panic(err)
	}

	db, _ := conn.DB()

	defer db.Close()

	
}
