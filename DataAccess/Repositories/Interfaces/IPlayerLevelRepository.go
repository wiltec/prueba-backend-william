package Interfaces

import (
	"PruebaBackendWilliam/Models"
	"PruebaBackendWilliam/Utils/Error/ErrorMessages"

	"gorm.io/gorm"
)

type IPlayerLevelRepository interface {
	CreatePlayerLevel(model *Models.PlayerLevelModel, db *gorm.DB) (bool, ErrorMessages.MessageValidator)
	UpdatePlayerLevel(model *Models.PlayerLevelModel, db *gorm.DB) (bool, ErrorMessages.MessageValidator)
	DeletePlayerLevel(model *Models.PlayerLevelModel, db *gorm.DB) (bool, ErrorMessages.MessageValidator)
}
