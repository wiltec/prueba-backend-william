package Interfaces

import (
	"PruebaBackendWilliam/Dto/Request"
	"PruebaBackendWilliam/Dto/Response"
	MSG "PruebaBackendWilliam/Utils/Error/ErrorMessages"
)

type IPlayerLevelService interface {
	RetrievePlayerLevel() []Response.PlayerLevelResponse
	CreatePlayerLevel(request *Request.CreatePlayerLevelRequest) (bool, MSG.MessageValidator)
	UpdatePlayerLevel(request *Request.UpdatePlayerLevelRequest) (bool, MSG.MessageValidator)
	DeletePlayerLevel(request *Request.DeletePlayerLevelRequest) (bool, MSG.MessageValidator)
}
