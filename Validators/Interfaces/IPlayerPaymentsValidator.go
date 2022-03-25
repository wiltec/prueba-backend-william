package Interfaces

import (
	"PruebaBackendWilliam/Dto/Request"
	MSG "PruebaBackendWilliam/Utils/Error/ErrorMessages"

	"gorm.io/gorm"
)

type IPlayerPaymentsValidator interface {
	CalculateTeamPaymentsValidator(request *Request.CalculateTeamPaymentsRequest, db *gorm.DB) (bool, MSG.MessageValidator)
	CalculateTeamPaymentsByListValidator(request *Request.CalculateTeamPaymentsByListRequest, db *gorm.DB) (bool, MSG.MessageValidator)
}
