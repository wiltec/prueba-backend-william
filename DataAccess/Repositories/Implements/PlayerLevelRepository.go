package Implements

import (
	"PruebaBackendWilliam/DataAccess/Repositories/Interfaces"
	"PruebaBackendWilliam/Models"
	"PruebaBackendWilliam/Utils/Error/ErrorMessages"

	"gorm.io/gorm"
)

type PlayerLevelRepository struct {
	Interfaces.IPlayerLevelRepository `inject:""`
	HelperRepository                  Interfaces.IHelperRespository `inject:""`
}

//CreatePlayerLevel es una funcion para crear un elemento en la base de datos
func (repository *PlayerLevelRepository) CreatePlayerLevel(model *Models.PlayerLevelModel, db *gorm.DB) (bool, ErrorMessages.MessageValidator) {

	er := db.Create(model).Error

	return repository.HelperRepository.GetMsg(er, "PlayerLevel", false)
}

//UpdatePlayerLevel es una funcion para actualizar un elemento en la base de datos
func (repository *PlayerLevelRepository) UpdatePlayerLevel(model *Models.PlayerLevelModel, db *gorm.DB) (bool, ErrorMessages.MessageValidator) {

	er := db.Save(model).Error

	return repository.HelperRepository.GetMsg(er, "PlayerLevel", false)
}

//DeletePlayerLevel es una funcion para eliminar un elemento en la base de datos
func (repository *PlayerLevelRepository) DeletePlayerLevel(model *Models.PlayerLevelModel, db *gorm.DB) (bool, ErrorMessages.MessageValidator) {

	er := db.Delete(model).Error

	return repository.HelperRepository.GetMsg(er, "PlayerLevel", false)
}
