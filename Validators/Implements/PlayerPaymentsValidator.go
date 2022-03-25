package Implements

import (
	"PruebaBackendWilliam/Dto/Request"
	IService "PruebaBackendWilliam/Services/Interfaces"
	MSG "PruebaBackendWilliam/Utils/Error/ErrorMessages"
	"PruebaBackendWilliam/Validators/Interfaces"
	"fmt"

	"gorm.io/gorm"
)

type PlayerPaymentsValidator struct {
	Interfaces.IPlayerPaymentsValidator `inject:""`
	PlayerLevel                         IService.IPlayerPaymentsService `inject:""`
}

//CalculateTeamPaymentsValidator: valida los datos de entrada para el cálculo del salario final
func (validator *PlayerPaymentsValidator) CalculateTeamPaymentsValidator(request *Request.CalculateTeamPaymentsRequest, db *gorm.DB) (bool, MSG.MessageValidator) {

	mapPlayerLevel := validator.PlayerLevel.GetPlayerLevel(db)

	if resp, err := validator.VerifyPlayerList(&request.Jugadores, mapPlayerLevel); !resp {

		return resp, err
	}

	return MSG.NewWithBool(MSG.Successfull)
}

//CalculateTeamPaymentsValidator: valida los datos de entrada de una lista de listas para el cálculo del salario final
func (validator *PlayerPaymentsValidator) CalculateTeamPaymentsByListValidator(request *Request.CalculateTeamPaymentsByListRequest, db *gorm.DB) (bool, MSG.MessageValidator) {

	if len(request.Equipos) <= 0 {

		return MSG.NewWithBool(MSG.Required_Data, "Equipos", "The list must contain at least one element")
	}

	mapPlayerLevel := validator.PlayerLevel.GetPlayerLevel(db)

	for _, team := range request.Equipos {

		if resp, err := validator.VerifyPlayerList(&team.Jugadores, mapPlayerLevel); !resp {

			return resp, err
		}
	}

	return MSG.NewWithBool(MSG.Successfull)
}

//VerifyPlayerList: Recorre una lista para verificar los datos de un jugador
func (validator *PlayerPaymentsValidator) VerifyPlayerList(players *[]Request.PlayerRequest, mapPlayerLevel map[string]int) (bool, MSG.MessageValidator) {

	if len(*players) <= 0 {

		return MSG.NewWithBool(MSG.Required_Data, "Jugadores", "The list must contain at least one element")
	}

	for _, player := range *players {

		//Verificamos si existe el Nivel del jugador
		if _, ok := mapPlayerLevel[player.Nivel]; !ok {

			return MSG.NewWithBool(MSG.Data_Not_Found, "Nivel", fmt.Sprintf("Level '%s' not found", player.Nivel))
		}

		if resp, err := validator.VerifyPlayer(&player); !resp {

			return resp, err
		}
	}

	return MSG.NewWithBool(MSG.Successfull)
}

//VerifyPlayer: verifica los datos de entrada de un jugador
func (validator *PlayerPaymentsValidator) VerifyPlayer(player *Request.PlayerRequest) (bool, MSG.MessageValidator) {

	if len(player.Nombre) <= 0 {

		return MSG.NewWithBool(MSG.Required_Data, "Nombre", "Required data")
	}

	if len(player.Nivel) <= 0 {

		return MSG.NewWithBool(MSG.Required_Data, "Nivel", "Required data")
	}

	if player.Goles <= 0 {

		return MSG.NewWithBool(MSG.Required_Data, "Goles", "Required data")
	}

	if player.Sueldo <= 0 {

		return MSG.NewWithBool(MSG.Required_Data, "Sueldo", "Required data")
	}

	if player.Bono <= 0 {

		return MSG.NewWithBool(MSG.Required_Data, "Bono", "Required data")
	}

	if len(player.Equipo) <= 0 {

		return MSG.NewWithBool(MSG.Required_Data, "Equipo", "Required data")
	}

	return MSG.NewWithBool(MSG.Successfull)
}
