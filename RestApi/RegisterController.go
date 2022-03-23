package RestApi

import (
	"PruebaBackendWilliam/RestApi/Routes"
	"PruebaBackendWilliam/Utils/DependencyInjector"

	"github.com/gin-gonic/gin"
)

func RegisterControllers() bool {
	route := gin.Default()
	var pruebaBackendRoutes Routes.PruebaBackendRoutes

	pruebaBackendRoutes.PlayerPayments = DependencyInjector.PlayerPaymentsController
	pruebaBackendRoutes.PlayerLevel = DependencyInjector.PlayerLevelController

	pruebaBackendRoutes.RegisterRoutes(route)
	port := "3001"
	route.Run(":" + port)

	return true
}
