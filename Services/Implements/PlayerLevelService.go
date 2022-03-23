package Implements

import (
	"PruebaBackendWilliam/DataAccess/Repositories/Interfaces"
	"PruebaBackendWilliam/Dto/Request"
	"PruebaBackendWilliam/Infraestructure/DataBase"
	"PruebaBackendWilliam/Models"
	MSG "PruebaBackendWilliam/Utils/Error/ErrorMessages"
)

//PlayerLevelService se crea un constructor donde se implementa la interfaz para la inyeccion de dependencias
type PlayerLevelService struct {
	Conn       DataBase.IDHandler                `inject:""`
	Repository Interfaces.IPlayerLevelRepository `inject:""`
}

//CreatePlayerLevel es una funcion para crear un elemento en la base de datos
func (service *PlayerLevelService) CreatePlayerLevel(request *Request.CreatePlayerLevelRequest) (bool, MSG.MessageValidator) {

	db, err := service.Conn.Connect()

	//Verificamos que no exista ningun problema con la conexion a la base de datos
	if err != nil {
		//Tiramos el panic para cerrar todos los procesos
		panic(err)
	}

	conn, _ := db.DB()

	defer conn.Close()

	model := Models.PlayerLevelModel{}

	//Volcamos los datos del request al modelo
	//copier.Copy(&model, request)
	model.Name = request.Name
	model.GoalMonth = request.GoalMonth

	/* if resp, msg := service.Validator.ValidateCreatePermission(&model, oAuthContext, conn); !resp {

		// se retorna el error
		return false, msg
	}

	//validamos los datos del modelo, sus relaciones y lo necesario para que se cree
	if resp, msg := service.Validator.ValidateCreatePlayerLevel(&model, conn); !resp {

		// se retorna el error
		return false, msg
	} */

	if ok, err := service.Repository.CreatePlayerLevel(&model, db); !ok {

		//Si existe algun error, entonces enviamos el mensaje correspondiente
		return false, err
	}

	//Regresamos el id del modelo insertado y regresamos un mensaje, dependiendo del tipo de error
	return MSG.NewWithBool(MSG.Successfull, model.IdPlayerLevel)

}

//UpdatePlayerLevel es una funcion para actualizar un elemento en la base de datos
func (service *PlayerLevelService) UpdatePlayerLevel(request *Request.UpdatePlayerLevelRequest) (bool, MSG.MessageValidator) {

	db, err := service.Conn.Connect()

	//Verificamos que no exista ningun problema con la conexion a la base de datos
	if err != nil {
		//Tiramos el panic para cerrar todos los procesos
		panic(err)
	}

	conn, _ := db.DB()

	defer conn.Close()

	//Verificamos el id
	if len(request.IdPlayerLevel) <= 0 {

		// Si no se envía el id retornamos false y el mensaje de error
		return MSG.NewWithBool(MSG.Data_Not_Found, "IdPlayerLevel", "There is no PlayerLevel indicated")

	}

	arrayModel := []Models.PlayerLevelModel{}
	db.Model(&Models.PlayerLevelModel{}).Where("id_player_level = ? ", request.IdPlayerLevel).Find(&arrayModel)

	if len(arrayModel) <= 0 {

		// Si no existe el elmento retornamos false y el mensaje de error
		return MSG.NewWithBool(MSG.Data_Not_Found, "IdPlayerLevel", "Data not found")
	}

	//Declaramos un modelo, el cual el resultado de la consulta
	model := arrayModel[0]

	model.GoalMonth = request.GoalMonth
	model.Name = request.Name

	//Verificamos que no se duplique la información y si existen cada uno de los Id's de las tablas relacionadas.
	/* resp, msg := service.Validator.ValidateUpdatePlayerLevel(&model, db)

	if !resp {

		// retorna false y el mensaje de error
		return false, msg
	} */

	//Realizamos la actualizacion del modelo por medio del repository
	if ok, message := service.Repository.UpdatePlayerLevel(&model, db); !ok {

		//Si existe algun error, entonces enviamos el mensaje correspondiente
		return false, message
	}

	//Devolvemos una respuesta boleana para afirmar la actualizacion, y el mensaje de exito
	return MSG.NewWithBool(MSG.Successfull)
}

//DeletePlayerLevel es una funcion para eliminar un elemento en la base de datos
func (service *PlayerLevelService) DeletePlayerLevel(request *Request.DeletePlayerLevelRequest) (bool, MSG.MessageValidator) {

	model := Models.PlayerLevelModel{
		IdPlayerLevel: request.IdPlayerLevel,
	}

	db, err := service.Conn.Connect()

	//Verificamos que no exista ningun problema con la conexion a la base de datos
	if err != nil {
		//Tiramos el panic para cerrar todos los procesos
		panic(err)
	}

	conn, _ := db.DB()

	defer conn.Close()

	/* if response, message := service.Validator.ValidateDeletePlayerLevel(&model, conn); !response {

		//retornamos false y elmensaje de error
		return response, message
	} */

	//ejecutamos el delete
	if ok, message := service.Repository.DeletePlayerLevel(&model, db); !ok {

		//Si existe algun error, entonces enviamos el mensaje correspondiente
		return false, message
	}

	// retornamos  true, y el mensaje sucess
	return MSG.NewWithBool(MSG.Successfull, "PlayerLevel", "Record deleted")

}
