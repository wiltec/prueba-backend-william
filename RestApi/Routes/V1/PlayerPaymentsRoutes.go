package V1

import (
	"PruebaBackendWilliam/RestApi/Controllers"

	"github.com/gin-gonic/gin"
)

//PlayerPaymentsRoutes: crea los accesos a los endpoints de playerPayments
func PlayerPaymentsRoutes(route *gin.RouterGroup, playerPayments *Controllers.PlayerPaymentsController) {

	route.POST("/calculateteampayments", playerPayments.CalculateTeamPayments)
}
