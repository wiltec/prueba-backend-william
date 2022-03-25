package Implements

import (
	"PruebaBackendWilliam/DataAccess/Queries/Interfaces"
	"PruebaBackendWilliam/Models"

	"gorm.io/gorm"
)

type PlayerLevelQuery struct {
	Interfaces.IPlayerLevelQuery `inject:""`
}

//RetrievePlayerLevel: obtiene los todos los datos del catálogo de niveles
func (query *PlayerLevelQuery) RetrievePlayerLevel(db *gorm.DB) []Models.PlayerLevelModel {

	arrayModel := []Models.PlayerLevelModel{}

	db.Model(&Models.PlayerLevelModel{}).Find(&arrayModel)

	return arrayModel
}
