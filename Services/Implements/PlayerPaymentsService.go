package Implements

import (
	"PruebaBackendWilliam/Dto/Request"
	"PruebaBackendWilliam/Infraestructure/DataBase"
	"PruebaBackendWilliam/Models"
	MSG "PruebaBackendWilliam/Utils/Error/ErrorMessages"
)

//PlayerPaymentsService se crea un constructor donde se implementa la interfaz para la inyeccion de dependencias
type PlayerPaymentsService struct {
	Conn DataBase.IDHandler `inject:""`
}

func (service *PlayerPaymentsService) CalculateTeamPayments(request *Request.CreatePlayerPaymentsRequest) (bool, MSG.MessageValidator) {

	db, err := service.Conn.Connect()

	//Verificamos que no exista ningun problema con la conexion a la base de datos
	if err != nil {
		//Tiramos el panic para cerrar todos los procesos
		panic(err)
	}

	conn, _ := db.DB()

	defer conn.Close()

	model := Models.PlayerPaymentsModel{}

	//Volcamos los datos del request al modelo
	//copier.Copy(&model, request)
	model.Name = request.Name
	model.GoalMonth = request.GoalMonth

	/* if resp, msg := service.Validator.ValidateCreatePermission(&model, oAuthContext, conn); !resp {

		// se retorna el error
		return false, msg
	}

	//validamos los datos del modelo, sus relaciones y lo necesario para que se cree
	if resp, msg := service.Validator.ValidateCreatePlayerPayments(&model, conn); !resp {

		// se retorna el error
		return false, msg
	} */

	//Regresamos el id del modelo insertado y regresamos un mensaje, dependiendo del tipo de error
	return MSG.NewWithBool(MSG.Successfull, model.IdPlayerPayments)

}
