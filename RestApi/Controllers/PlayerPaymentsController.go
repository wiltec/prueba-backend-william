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

//CalculateTeamPayments: Metodo POST para cálcular el sueldo final de una lista de jugadores
func (controller *PlayerPaymentsController) CalculateTeamPayments(context *gin.Context) {

	var request Request.CalculateTeamPaymentsRequest
	context.Bind(&request)

	result, response := controller.Service.CalculateTeamPayments(&request)
	if result {

		context.JSON(http.StatusOK, response)
	} else {

		context.JSON(http.StatusConflict, response)
	}
}

//CalculateTeamPaymentsByList: Metodo POST para cálcular el sueldo final de una lista de lista de jugadores
func (controller *PlayerPaymentsController) CalculateTeamPaymentsByList(context *gin.Context) {

	var request Request.CalculateTeamPaymentsByListRequest
	context.Bind(&request)

	result, response := controller.Service.CalculateTeamPaymentsByList(&request)
	if result {

		context.JSON(http.StatusOK, response)
	} else {

		context.JSON(http.StatusConflict, response)
	}
}
