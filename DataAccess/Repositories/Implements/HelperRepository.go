package Implements

import (
	"PruebaBackendWilliam/DataAccess/Repositories/Interfaces"
	"PruebaBackendWilliam/Utils/Error/ErrorMessages"
)

type HelperRepository struct {
	Interfaces.IHelperRespository `inject:""`
}

//GetMsg : retorna un booleano y un mensaje de error
func (helper *HelperRepository) GetMsg(er error, modelName string, debug bool) (bool, ErrorMessages.MessageValidator) {

	// Definimos un error por defecto
	message := ErrorMessages.MessageValidator{
		Code:    ErrorMessages.Successfull,
		Name:    modelName,
		Message: ErrorMessages.ReturnMessage(ErrorMessages.Successfull),
	}

	// si hay erorr
	if er != nil {

		message.Description = "A unknown error Ocurred"

		// si recibimos la peticion de debug
		if debug {

			// a la descripcion le ponemos el error recibido
			message.Description = er.Error()
		}

		message.Code = ErrorMessages.Error_database
		message.Message = ErrorMessages.ReturnMessage(ErrorMessages.Error_database)

		// Para saber el error poner aqui er.Error como description
		return false, message
	}

	// si no hay error devolvemos el mensaje de true, succes como mensaje
	return true, ErrorMessages.MessageValidator{

		Code:    ErrorMessages.Successfull,
		Name:    modelName,
		Message: ErrorMessages.ReturnMessage(ErrorMessages.Successfull),
	}
}
