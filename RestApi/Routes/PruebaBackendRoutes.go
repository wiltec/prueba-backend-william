package Routes

import (
	"PruebaBackendWilliam/RestApi/Controllers"
	"PruebaBackendWilliam/RestApi/Routes/V1"

	"github.com/gin-gonic/gin"
)

type PruebaBackendRoutes struct {
	PlayerPayments Controllers.PlayerPaymentsController
	PlayerLevel    Controllers.PlayerLevelController
}

//RegisterRoutes: Crea los grupos para las rutas a los endpoints
func (routes *PruebaBackendRoutes) RegisterRoutes(route *gin.Engine) {

	apiV1 := route.Group("/v1")
	V1.PlayerPaymentsRoutes(apiV1.Group("/playerpayments"), &routes.PlayerPayments)
	V1.PlayerLevelRoutes(apiV1.Group("/playerlevel"), &routes.PlayerLevel)
}
