package AutoMigration

import (
	"PruebaBackendWilliam/Models"
	"strconv"

	"gorm.io/gorm"
)

func CreateTable(db *gorm.DB) {
	CreateTableFromModel(PlayerLevel, db)
}

func CreateTableFromModel(model interface{}, db *gorm.DB) {
	exist := db.Migrator().HasTable(model)
	if !exist {
		db.Migrator().CreateTable(model)
	}
}

func DropTableFromModel(model interface{}, db *gorm.DB) {
	exist := db.Migrator().HasTable(model)
	if exist {
		db.Migrator().DropTable(model)
	}
}

func LoadDefaultData(db *gorm.DB) {

	//Se agregan valores por defecto para el catálogo PlayerLevel
	for _, item := range PlayerLevelValuesDefault {

		//Convertimos el valor de goalmonth a entero
		goalMonth, err := strconv.Atoi(item["GoalMonth"])

		if err != nil {
			panic(err)
		}

		playerLevelModel := Models.PlayerLevelModel{
			Name:      item["Name"],
			GoalMonth: goalMonth,
		}

		db.Where(&playerLevelModel).FirstOrCreate(&playerLevelModel)
	}

}
