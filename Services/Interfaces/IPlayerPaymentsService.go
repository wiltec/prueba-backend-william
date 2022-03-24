package Interfaces

import (
	"PruebaBackendWilliam/Dto/Request"
	"PruebaBackendWilliam/Dto/Response"
)

type IPlayerPaymentsService interface {
	CalculateTeamPayments(request *Request.CalculateTeamPaymentsRequest) Response.CalculateTeamPaymentsResponse
	CalculateTeamPaymentsByList(request *Request.CalculateTeamPaymentsByListRequest) Response.CalculateTeamPaymentsByListResponse
}
