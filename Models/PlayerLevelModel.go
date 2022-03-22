package Models

//PlayerLevelModel es un struct que establece el esquema de la tabla niveles de jugador
type PlayerLevelModel struct {
	IdPlayerLevel string `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name          string `gorm:"type:varchar(10);not null;UNIQUE"`
	GoalMonth     int    `gorm:"type:integer;not null"`
}

func (model *PlayerLevelModel) TableName() string {
	return "playerlevel"
}
