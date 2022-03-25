package Implements

import (
	IQuery "PruebaBackendWilliam/DataAccess/Queries/Interfaces"
	"PruebaBackendWilliam/DataAccess/Repositories/Interfaces"
	"PruebaBackendWilliam/Dto/Request"
	"PruebaBackendWilliam/Dto/Response"
	"PruebaBackendWilliam/Infraestructure/DataBase"
	"PruebaBackendWilliam/Models"
	IServices "PruebaBackendWilliam/Services/Interfaces"
	MSG "PruebaBackendWilliam/Utils/Error/ErrorMessages"
	IValidator "PruebaBackendWilliam/Validators/Interfaces"
)

type PlayerLevelService struct {
	IServices.IPlayerLevelService `inject:""`
	Conn                          DataBase.IDHandler                `inject:""`
	Repository                    Interfaces.IPlayerLevelRepository `inject:""`
	Query                         IQuery.IPlayerLevelQuery          `inject:""`
	Validator                     IValidator.IPlayerLevelValidator  `inject:""`
}

//RetrievePlayerLevel: es una función que consulta todos los registros
func (service *PlayerLevelService) RetrievePlayerLevel() []Response.PlayerLevelResponse {

	db, err := service.Conn.Connect()

	//Verificamos que no exista ningun problema con la conexion a la base de datos
	if err != nil {
		//Tiramos el panic para cerrar todos los procesos
		panic(err)
	}

	conn, _ := db.DB()

	defer conn.Close()

	arrayResponse := []Response.PlayerLevelResponse{}

	arrayModel := service.Query.RetrievePlayerLevel(db)

	for _, item := range arrayModel {

		response := Response.PlayerLevelResponse{
			IdPlayerLevel: item.IdPlayerLevel,
			Name:          item.Name,
			GoalMonth:     item.GoalMonth,
		}

		arrayResponse = append(arrayResponse, response)
	}

	return arrayResponse

}

//CreatePlayerLevel: es una funcion para crear un elemento en la base de datos
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
	model.Name = request.Name
	model.GoalMonth = request.GoalMonth

	//validamos los datos
	if resp, msg := service.Validator.ValidateCreatePlayerLevel(&model, db); !resp {

		return resp, msg
	}

	if ok, err := service.Repository.CreatePlayerLevel(&model, db); !ok {

		return false, err
	}

	return MSG.NewWithBool(MSG.Successfull, model.IdPlayerLevel)

}

//UpdatePlayerLevel: es una funcion para actualizar un elemento en la base de datos
func (service *PlayerLevelService) UpdatePlayerLevel(request *Request.UpdatePlayerLevelRequest) (bool, MSG.MessageValidator) {

	db, err := service.Conn.Connect()

	//Verificamos que no exista ningun problema con la conexion a la base de datos
	if err != nil {
		//Tiramos el panic para cerrar todos los procesos
		panic(err)
	}

	conn, _ := db.DB()

	defer conn.Close()

	model := Models.PlayerLevelModel{
		IdPlayerLevel: request.IdPlayerLevel,
		Name:          request.Name,
		GoalMonth:     request.GoalMonth,
	}

	//validamos los datos
	if resp, msg := service.Validator.ValidateUpdatePlayerLevel(&model, db); !resp {

		return resp, msg
	}

	arrayModel := []Models.PlayerLevelModel{}
	db.Model(&Models.PlayerLevelModel{}).Where("id_player_level = ? ", request.IdPlayerLevel).Find(&arrayModel)

	if len(arrayModel) <= 0 {

		return MSG.NewWithBool(MSG.Data_Not_Found, "IdPlayerLevel", "Data not found")
	}

	//Declaramos un modelo, el cual el resultado de la consulta
	modelCurrent := arrayModel[0]

	modelCurrent.GoalMonth = model.GoalMonth
	modelCurrent.Name = model.Name

	//Realizamos la actualizacion del modelo por medio del repository
	if ok, message := service.Repository.UpdatePlayerLevel(&modelCurrent, db); !ok {

		return false, message
	}

	return MSG.NewWithBool(MSG.Successfull)
}

//DeletePlayerLevel es una funcion para eliminar un elemento en la base de datos
func (service *PlayerLevelService) DeletePlayerLevel(request *Request.DeletePlayerLevelRequest) (bool, MSG.MessageValidator) {

	db, err := service.Conn.Connect()

	//Verificamos que no exista ningun problema con la conexion a la base de datos
	if err != nil {
		//Tiramos el panic para cerrar todos los procesos
		panic(err)
	}

	conn, _ := db.DB()

	defer conn.Close()

	model := Models.PlayerLevelModel{
		IdPlayerLevel: request.IdPlayerLevel,
	}

	if response, message := service.Validator.ValidateDeletePlayerLevel(&model, db); !response {

		return response, message
	}

	//ejecutamos el delete
	if ok, message := service.Repository.DeletePlayerLevel(&model, db); !ok {

		return false, message
	}

	return MSG.NewWithBool(MSG.Successfull, "PlayerLevel", "Record deleted")
}
