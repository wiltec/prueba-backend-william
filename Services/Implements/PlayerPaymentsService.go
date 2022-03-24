package Implements

import (
	"PruebaBackendWilliam/Dto/Request"
	"PruebaBackendWilliam/Dto/Response"
	"PruebaBackendWilliam/Infraestructure/DataBase"
	"PruebaBackendWilliam/Models"
	"PruebaBackendWilliam/Services/Interfaces"

	"gorm.io/gorm"
)

//PlayerPaymentsService se crea un constructor donde se implementa la interfaz para la inyeccion de dependencias
type PlayerPaymentsService struct {
	Interfaces.IPlayerPaymentsService `inhject:""`
	Conn                              DataBase.IDHandler `inject:""`
}

func (service *PlayerPaymentsService) CalculateTeamPayments(request *Request.CalculateTeamPaymentsRequest) Response.CalculateTeamPaymentsResponse {

	db, err := service.Conn.Connect()

	//Verificamos que no exista ningun problema con la conexion a la base de datos
	if err != nil {
		//Tiramos el panic para cerrar todos los procesos
		panic(err)
	}

	conn, _ := db.DB()

	defer conn.Close()

	//response := Response.CalculateTeamPaymentsResponse{}

	/*//validamos los datos del modelo, sus relaciones y lo necesario para que se cree
	if resp, msg := service.Validator.ValidateCreatePlayerPayments(&model, conn); !resp {

		// se retorna el error
		return false, msg
	} */

	listPlayerSalaryTemp := service.CalculateFullSalary(request.Jugadores, db)

	response := Response.CalculateTeamPaymentsResponse{}
	listPlayerResponse := []Response.PlayerResponse{}

	//Se agrega cada elemento de la lista temporal a una nueva lista para el response
	for _, item := range listPlayerSalaryTemp {

		playerResponse := Response.PlayerResponse{
			Nombre:          item.Nombre,
			Goles_minimos:   item.Goles_minimos,
			Goles:           item.Goles,
			Sueldo:          item.Sueldo,
			Bono:            item.Bono,
			Sueldo_completo: item.Sueldo_completo,
			Equipo:          item.Equipo,
		}

		listPlayerResponse = append(listPlayerResponse, playerResponse)
	}

	response.Jugadores = listPlayerResponse

	return response
}

func (service *PlayerPaymentsService) CalculateTeamPaymentsByList(request *Request.CalculateTeamPaymentsByListRequest) Response.CalculateTeamPaymentsByListResponse {

	db, err := service.Conn.Connect()

	//Verificamos que no exista ningun problema con la conexion a la base de datos
	if err != nil {
		//Tiramos el panic para cerrar todos los procesos
		panic(err)
	}

	conn, _ := db.DB()

	defer conn.Close()

	//response := Response.CalculateTeamPaymentsResponse{}

	/*//validamos los datos del modelo, sus relaciones y lo necesario para que se cree
	if resp, msg := service.Validator.ValidateCreatePlayerPayments(&model, conn); !resp {

		// se retorna el error
		return false, msg
	} */

	response := Response.CalculateTeamPaymentsByListResponse{}
	ListCalculateTeamPaymentsResponse := []Response.CalculateTeamPaymentsResponse{}

	listArrayPlayerSalaryTemp := []Models.ArrayPlayerSalaryTemp{}

	for _, item := range request.Equipos {

		listPlayerSalaryTemp := service.CalculateFullSalary(item.Jugadores, db)

		arrayPlayerSalaryTemp := Models.ArrayPlayerSalaryTemp{
			Jugadores: listPlayerSalaryTemp,
		}

		listArrayPlayerSalaryTemp = append(listArrayPlayerSalaryTemp, arrayPlayerSalaryTemp)
	}

	for _, arrayPlayerSalaryTemp := range listArrayPlayerSalaryTemp {

		calculateTeamPaymentResponse := Response.CalculateTeamPaymentsResponse{}
		listPlayerResponse := []Response.PlayerResponse{}

		for _, playerSalaryTemp := range arrayPlayerSalaryTemp.Jugadores {

			playerResponse := Response.PlayerResponse{
				Nombre:          playerSalaryTemp.Nombre,
				Goles_minimos:   playerSalaryTemp.Goles_minimos,
				Goles:           playerSalaryTemp.Goles,
				Sueldo:          playerSalaryTemp.Sueldo,
				Bono:            playerSalaryTemp.Bono,
				Sueldo_completo: playerSalaryTemp.Sueldo_completo,
				Equipo:          playerSalaryTemp.Equipo,
			}

			listPlayerResponse = append(listPlayerResponse, playerResponse)

		}

		calculateTeamPaymentResponse.Jugadores = listPlayerResponse

		ListCalculateTeamPaymentsResponse = append(ListCalculateTeamPaymentsResponse, calculateTeamPaymentResponse)
	}

	response.Equipos = ListCalculateTeamPaymentsResponse

	return response
}

func (service *PlayerPaymentsService) CalculateFullSalary(players []Request.PlayerRequest, db *gorm.DB) []Models.PlayerSalaryTemp {

	//Obtenemos el catálogo de niveles de jugador
	mapPlayerLevel := service.GetPlayerLevel(db)

	listPlayerSalaryTemp := []Models.PlayerSalaryTemp{}

	//Goles por equipo
	minimumGoalsTeam := 0

	//Golesl anotados por equipo
	goalsScoredTeam := 0

	for _, item := range players {

		//Goles minimos
		minimumGoals := mapPlayerLevel[item.Nivel]

		//Goles anotados
		goalsScored := item.Goles

		//Acumular minimo de goles
		minimumGoalsTeam = minimumGoalsTeam + minimumGoals

		//Acumular goles anotados
		goalsScoredTeam = goalsScoredTeam + goalsScored

		//Porcentage del bono individual
		individualPercentage := service.CalculateBonusPercentage(float32(goalsScored), float32(minimumGoals))

		playerSalaryTemp := Models.PlayerSalaryTemp{
			Nombre:               item.Nombre,
			Nivel:                item.Nivel,
			Goles:                item.Goles,
			Sueldo:               item.Sueldo,
			Bono:                 item.Bono,
			Equipo:               item.Equipo,
			PorcentageIndividual: individualPercentage,
			Goles_minimos:        minimumGoals,
		}

		listPlayerSalaryTemp = append(listPlayerSalaryTemp, playerSalaryTemp)

	}

	//Porcentage del bono por equipo
	teamPercentage := service.CalculateBonusPercentage(float32(goalsScoredTeam), float32(minimumGoalsTeam))

	//Se obtiene el salario final por jugador
	service.GetFullSalary(teamPercentage, &listPlayerSalaryTemp)

	return listPlayerSalaryTemp
}

func (service *PlayerPaymentsService) GetPlayerLevel(db *gorm.DB) map[string]int {

	//Declarmos un map para devolver los valores
	mapPlayerLevel := make(map[string]int)

	arrayModel := []Models.PlayerLevelModel{}

	db.Model(&Models.PlayerLevelModel{}).Find(&arrayModel)

	if len(arrayModel) > 0 {

		for _, item := range arrayModel {

			mapPlayerLevel[item.Name] = item.GoalMonth
		}
	}

	return mapPlayerLevel
}

func (service *PlayerPaymentsService) CalculateBonusPercentage(goalsScored float32, minimumGoals float32) float32 {

	var percentage float32

	percentage = (goalsScored * 100) / minimumGoals

	//Si el porcentage es mayor al 100% se estable en 100
	if percentage > 100 {
		percentage = 100
	}

	return percentage
}

func (service *PlayerPaymentsService) GetFullSalary(teamPercentage float32, listPlayer *[]Models.PlayerSalaryTemp) *[]Models.PlayerSalaryTemp {

	for i := 0; i < len(*listPlayer); i++ {

		percentageTotal := (teamPercentage + (*listPlayer)[i].PorcentageIndividual) / 2

		bonusAmount := (percentageTotal * (*listPlayer)[i].Bono) / 100

		(*listPlayer)[i].Sueldo_completo = bonusAmount + (*listPlayer)[i].Sueldo
	}

	return listPlayer
}
