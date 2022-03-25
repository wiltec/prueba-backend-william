package Implements

import (
	"PruebaBackendWilliam/Models"
	MSG "PruebaBackendWilliam/Utils/Error/ErrorMessages"
	"PruebaBackendWilliam/Validators/Interfaces"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

type PlayerLevelValidator struct {
	Interfaces.IPlayerLevelValidator `inject:""`
}

//ValidateCreatePlayerLevel: valida los datos de entrada para una creación
func (validator *PlayerLevelValidator) ValidateCreatePlayerLevel(model *Models.PlayerLevelModel, db *gorm.DB) (bool, MSG.MessageValidator) {

	// Validamos los campos obligatorios
	if len(model.Name) <= 0 {
		return MSG.NewWithBool(MSG.Required_Data, "Name", "Required data")
	}

	if model.GoalMonth <= 0 {

		return MSG.NewWithBool(MSG.Required_Data, "GoalMonth", "Required data")
	}

	// Validamos que no exista un elemento con los mismos datos en la base de datos
	if duplicated, message := validator.IsDuplicated(model, db); duplicated {

		return false, message
	}

	return MSG.NewWithBool(MSG.Successfull)
}

//ValidateUpdatePlayerLevel: valida los datos de entrada de una actualización
func (validator *PlayerLevelValidator) ValidateUpdatePlayerLevel(model *Models.PlayerLevelModel, db *gorm.DB) (bool, MSG.MessageValidator) {

	// Validamos los campos obligatorios
	if len(model.IdPlayerLevel) <= 0 {

		return MSG.NewWithBool(MSG.Required_Data, "IdPlayerLevel", "Required data")
	}

	if len(model.Name) <= 0 {

		return MSG.NewWithBool(MSG.Required_Data, "Name", "Required data")
	}

	if model.GoalMonth <= 0 {

		return MSG.NewWithBool(MSG.Required_Data, "GoalMonth", "Required data")
	}

	// Validamos que el modelo exista.
	if exists, message := validator.Exists(model, db); !exists {

		return false, message
	}

	// Validamos que el modelo no este duplicado
	if isDuplicated, message := validator.IsDuplicated(model, db); isDuplicated {

		return false, message
	}

	return MSG.NewWithBool(MSG.Successfull)
}

//ValidateDeletePlayerLevel: valida si existe el id de registro a eliminar
func (validator *PlayerLevelValidator) ValidateDeletePlayerLevel(model *Models.PlayerLevelModel, db *gorm.DB) (bool, MSG.MessageValidator) {

	// Validamos los campos obligatorios
	if len(model.IdPlayerLevel) <= 0 {
		return MSG.NewWithBool(MSG.Required_Data, "IdPlayerLevel", "Required data")
	}

	// Validamos que el modelo exista.
	if exists, message := validator.Exists(model, db); !exists {

		return false, message
	}

	return MSG.NewWithBool(MSG.Successfull)
}

//IsDuplicated  : retorna verdadero si el registro ya existe
func (validator *PlayerLevelValidator) IsDuplicated(model *Models.PlayerLevelModel, db *gorm.DB) (bool, MSG.MessageValidator) {

	// Realizamos la consulta con los parametros unicos
	arrayModel := []Models.PlayerLevelModel{}
	db.Model(&Models.PlayerLevelModel{}).
		Where("name = ? ", model.Name).
		Scan(&arrayModel)

	// Validamios el resultado de la consulta
	exists, message := validator.validateExistId(
		model,
		arrayModel,
		"Name",
		fmt.Sprintf("The %s code already exists", model.Name),
	)

	// si el elemento existe
	if exists {

		return exists, message
	}

	return false, MSG.New(MSG.Successfull)

}

//Exists : indica si un modelo esta existe por su id
func (validator *PlayerLevelValidator) Exists(model *Models.PlayerLevelModel, db *gorm.DB) (bool, MSG.MessageValidator) {

	// Definimos la consulta y el arreglo que tendrá los resultados
	arrayModel := []Models.PlayerLevelModel{}

	//Consultamos por el IdPlayerLevel
	db.Model(&Models.PlayerLevelModel{}).Where("id_player_level = ?", model.IdPlayerLevel).Find(&arrayModel)

	// Si hay resultados significa que el modelo existe en la base de datos
	if len(arrayModel) > 0 {

		return true, MSG.New(MSG.Duplicated_Field, "PlayerLevel", "The PlayerLevel exists")
	}

	return MSG.NewWithBool(MSG.Data_Not_Found, "PlayerLevel", "The PlayerLevel not exists")
}

//validateExistId : Valida si el id a actualizar es diferente del recupeado con el objetivo de detectar
func (validator *PlayerLevelValidator) validateExistId(model *Models.PlayerLevelModel, arrayModel []Models.PlayerLevelModel, name string, description string) (bool, MSG.MessageValidator) {

	// Declaramos el mensaje que se enviara si existe el elemento
	message := MSG.MessageValidator{

		Code:        MSG.Duplicated_Field,
		Message:     MSG.ReturnMessage(MSG.Duplicated_Field),
		Name:        name,
		Description: description,
	}

	// Si Hubo resultados significa que si existe
	if len(arrayModel) > 0 {

		if strings.TrimSpace(model.IdPlayerLevel) == "" {

			return true, message
		}

		if strings.TrimSpace(model.IdPlayerLevel) != "" &&
			arrayModel[0].IdPlayerLevel != model.IdPlayerLevel {

			return true, message
		}
	}

	return false, MSG.New(MSG.Successfull)
}
