package DataBase

import (
	"PruebaBackendWilliam/Infraestructure/Configuration"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBHandler struct {
	IDHandler `inject:""`
}

var xmlPath = "../Infraestructure/Configuration/BDSetting.xml"

//Connect : abrimos conexion a la base de datos
func (dbh *DBHandler) Connect() (*gorm.DB, error) {

	//Abrimos archivo
	xmlFile, err := os.Open(xmlPath)

	//Verificamos que no existan errores
	if err != nil {
		panic(err)
	}

	//Por defecto siempre cerramos el archivo
	defer xmlFile.Close()

	//Leemos en un byte array el contenido del archivo
	dbSettingsByteArray, _ := ioutil.ReadAll(xmlFile)

	//Declaramos variables
	DBs := Configuration.DBSettingsModel{}

	//volcamos el contenido del byte array en las estructuras
	xml.Unmarshal(dbSettingsByteArray, &DBs)

	//Creamos un item de tipo DB
	DBItem := DBs.DataBase[0]

	//Creamos la cadena de conexion
	strcon := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=%s",
		DBItem.Server,
		DBItem.Port,
		DBItem.User,
		DBItem.Database,
		DBItem.Password,
		DBItem.SslMode,
	)

	//Abrimos una conexion a la base de datos
	db, err := gorm.Open(postgres.Open(strcon), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	//Si el error es distinto a nil tiramos el error
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Conexion abierta exitosamente a base dedatos: " + DBItem.Name)
	}

	//Creamos la extension para uuid
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")

	//Devolvemos la conexion abierta a la base de datos
	return db, nil
}
