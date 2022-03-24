package Controllers

import (
	"PruebaBackendWilliam/Dto/Request"
	"PruebaBackendWilliam/Services/Interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PlayerPaymentsController struct {
	Service Interfaces.IPlayerPaymentsService `inject:""`
}

func (controller *PlayerPaymentsController) CalculateTeamPayments(context *gin.Context) {

	var request Request.CalculateTeamPaymentsRequest
	context.Bind(&request)

	response := controller.Service.CalculateTeamPayments(&request)
	context.JSON(http.StatusOK, response)
}

func (controller *PlayerPaymentsController) CalculateTeamPaymentsByList(context *gin.Context) {

	var request Request.CalculateTeamPaymentsByListRequest
	context.Bind(&request)

	response := controller.Service.CalculateTeamPaymentsByList(&request)
	context.JSON(http.StatusOK, response)
}
