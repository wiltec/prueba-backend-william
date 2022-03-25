package Interfaces

import (
	"PruebaBackendWilliam/Models"

	"gorm.io/gorm"
)

type IPlayerLevelQuery interface {
	RetrievePlayerLevel(db *gorm.DB) []Models.PlayerLevelModel
}
