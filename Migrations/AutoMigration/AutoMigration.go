package AutoMigration

import "gorm.io/gorm"

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
