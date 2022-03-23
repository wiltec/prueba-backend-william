package Interfaces

import (
	"PruebaBackendWilliam/Dto/Request"
	MSG "PruebaBackendWilliam/Utils/Error/ErrorMessages"
)

type IPlayerLevelService interface {
	CreatePlayerLevel(request *Request.CreatePlayerLevelRequest) (bool, MSG.MessageValidator)
	UpdatePlayerLevel(request *Request.UpdatePlayerLevelRequest) (bool, MSG.MessageValidator)
	DeletePlayerLevel(request *Request.DeletePlayerLevelRequest) (bool, MSG.MessageValidator)
}
