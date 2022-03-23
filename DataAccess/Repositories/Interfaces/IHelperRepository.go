package Interfaces

import "PruebaBackendWilliam/Utils/Error/ErrorMessages"

type IHelperRespository interface {
	GetMsg(er error, modelName string, debug bool) (bool, ErrorMessages.MessageValidator)
}
