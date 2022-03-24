package Implements

import (
	"PruebaBackendWilliam/Models"
	MSG "PruebaBackendWilliam/Utils/Error/ErrorMessages"
	"PruebaBackendWilliam/Validators/Interfaces"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

//PlayerLevelValidator es un struct que almacenara todas las dependencias
//que necesiten injectarse y que usaran los validadores
type PlayerLevelValidator struct {
	Interfaces.IPlayerLevelValidator `inject:""`
}

//ValidateCreatePlayerLevel es una validacion Publica que valida que pueda crearse un nuevo elemento PlayerLevel en la base de datos
func (validator *PlayerLevelValidator) ValidateCreatePlayerLevel(model *Models.PlayerLevelModel, db *gorm.DB) (bool, MSG.MessageValidator) {

	// Validamos los campos obligatorios
	if model.GoalMonth <= 0 {
		return
	}

	// Validamos que no exista un elemento con los mismos datos en la base de datos
	if duplicated, message := validator.IsDuplicated(model, db); duplicated {

		// Si el modelo ya existe enviamos false y el mensaje de error
		return false, message
	}

	//Devolvemos la comprobación y el mensaje de éxito
	return MSG.NewWithBool(MSG.Successfull)
}

//ValidateUpdatePlayerLevel es una validación Publica que verifica que puedan actualizarse
func (validator *PlayerLevelValidator) ValidateUpdatePlayerLevel(model *Models.PlayerLevelModel, db *gorm.DB) (bool, MSG.MessageValidator) {

	// Validamos los campos que son obligatorios con ayuda del Helper validator
	ok, message := validator.HelperValidator.ValidateFields(*model, []string{

		"IdPlayerLevel",
		"Active",
	})

	// Si algún campo no pasa la validación
	if !ok {

		//Devolvemos la comprobación y el error
		return false, message
	}

	// Validamos que el modelo exista.
	if exists, message := validator.Exists(model, db); !exists {

		// Si el modelo no existe retornamos false y el mensaje de error
		return false, message
	}

	// Validamos que el modelo no este duplicado
	if isDuplicated, message := validator.IsDuplicated(model, db); isDuplicated {

		// Si el modelo esta duplicado retornamos falso y el mensaje de error.
		return false, message
	}

	// Si se pasaron las validaciones retornamos verdadero y un mensaje de Success
	return MSG.NewWithBool(MSG.Successfull)
}

//ValidateDeletePlayerLevel : es una validación Publica que valida que exista un elemento para eliminar
func (validator *PlayerLevelValidator) ValidateDeletePlayerLevel(model *Models.PlayerLevelModel, db *gorm.DB) (bool, MSG.MessageValidator) {

	// Validamos los campos que son obligatorios con ayuda del Helper validator
	ok, message := validator.HelperValidator.ValidateFields(*model, []string{

		"IdPlayerLevel",
	})

	// Si algún campo no pasa la validación
	if !ok {

		//Devolvemos la comprobación y el error
		return false, message
	}

	// Validamos que existe un elemento con el UUID
	if exists, message := validator.Exists(model, db); !exists {

		// Si no existe enviamos false, y un mensaje de validación con el error
		return false, message
	}

	// Si todo salio bien enviamos true y un mensaje de validación satisfactorio
	return MSG.NewWithBool(MSG.Successfull)
}

//IsDuplicated  : retorna verdadero y un mensaje validator con código MSG.Duplicated_Field si el registro ya existe
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

		// si existe retornamos true, y el mensaje
		return exists, message
	}

	// Si el modelo no existe retornamos false con el mensaje de success
	return false, MSG.New(MSG.Successfull)

}

//Exists : Método que indica si un modelo esta existe por su UUID
func (validator *PlayerLevelValidator) Exists(model *Models.PlayerLevelModel, db *gorm.DB) (bool, MSG.MessageValidator) {

	// Definimos la consulta y el arreglo que tendrá los resultados
	arrayModel := []Models.PlayerLevelModel{}

	//Consultamos por el IdPlayerLevel
	db.Model(&Models.PlayerLevelModel{}).Where("id_player_model = ?", model.IdPlayerLevel).Find(&arrayModel)

	// Si hay resultados significa que el modelo existe en la base de datos
	if len(arrayModel) > 0 {

		// Retornamos true, y el mensaje de duplicated
		return true, MSG.New(MSG.Duplicated_Field, "PlayerLevel", "The PlayerLevel exists")
	}

	// retornamos false y el mensaje de no existe
	return MSG.NewWithBool(MSG.Data_Not_Found, "PlayerLevel", "The PlayerLevel exists")
}

//validateExistId : el arreglo resultante para usarlo en el update
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

		// Verificamos si recibimos el UUID del modelo.
		// Si no lo recibimos asumimos que es un create por lo que solo validamos que no existen los campos únicos
		// Sin tomar en cuenta el UUID
		if strings.TrimSpace(model.IdPlayerLevel) == "" {

			// Si existe enviamos true y un mensaje de validación con duplicated code
			return true, message
		}

		// Si enviaron el UUID verificamos que el código del modelo resultante no sea el mismo que el recibido
		// Si el código es el mismo significa que no se esta duplicando, unicamente se esta actualizando el
		// modelo
		if strings.TrimSpace(model.IdPlayerLevel) != "" &&
			arrayModel[0].IdPlayerLevel != model.IdPlayerLevel {

			// Si existe enviamos true y un mensaje de validación con duplicated code
			return true, message
		}
	}

	// Si el modelo no existe retornamos false con el mensaje de success
	return false, MSG.New(MSG.Successfull)
}
