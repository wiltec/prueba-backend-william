package V1

import (
	"PruebaBackendWilliam/RestApi/Controllers"

	"github.com/gin-gonic/gin"
)

//PlayerLevelRoutes: crea los accesos a los endpoints de playerLevel
func PlayerLevelRoutes(route *gin.RouterGroup, playerLevel *Controllers.PlayerLevelController) {

	route.GET("/retrieve", playerLevel.RetrievePlayerLevel)
	route.POST("/create", playerLevel.CreatePlayerLevel)
	route.PUT("/update", playerLevel.UpdatePlayerLevel)
	route.DELETE("/delete", playerLevel.DeletePlayerLevel)
}
