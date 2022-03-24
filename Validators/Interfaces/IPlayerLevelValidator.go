package Interfaces

import (
	"PruebaBackendWilliam/Models"
	MSG "PruebaBackendWilliam/Utils/Error/ErrorMessages"

	"gorm.io/gorm"
)

type IPlayerLevelValidator interface {
	ValidateCreatePlayerLevel(model *Models.PlayerLevelModel, db *gorm.DB) (bool, MSG.MessageValidator)
	ValidateUpdatePlayerLevel(model *Models.PlayerLevelModel, db *gorm.DB) (bool, MSG.MessageValidator)
	ValidateDeletePlayerLevel(model *Models.PlayerLevelModel, db *gorm.DB) (bool, MSG.MessageValidator)
}
