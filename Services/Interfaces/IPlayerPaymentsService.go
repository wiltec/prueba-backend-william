package Interfaces

import (
	"PruebaBackendWilliam/Dto/Request"

	"gorm.io/gorm"
)

type IPlayerPaymentsService interface {
	CalculateTeamPayments(request *Request.CalculateTeamPaymentsRequest) (bool, interface{})
	CalculateTeamPaymentsByList(request *Request.CalculateTeamPaymentsByListRequest) (bool, interface{})
	GetPlayerLevel(db *gorm.DB) map[string]int
}
