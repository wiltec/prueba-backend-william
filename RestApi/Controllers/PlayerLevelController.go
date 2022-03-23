package Controllers

import (
	"PruebaBackendWilliam/Dto/Request"
	"PruebaBackendWilliam/Services/Interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PlayerLevelController struct {
	Service Interfaces.IPlayerLevelService `inject:""`
}

func (controller *PlayerLevelController) CreatePlayerLevel(context *gin.Context) {
	var request Request.CreatePlayerLevelRequest
	context.Bind(&request)

	response, errorResponse := controller.Service.CreatePlayerLevel(&request)
	if response {

		context.JSON(http.StatusOK, errorResponse)
	} else {

		context.JSON(http.StatusConflict, errorResponse)
	}
}

//UpdatePlayerLevel Metodo PUT para editar un PlayerLevel
func (controller *PlayerLevelController) UpdatePlayerLevel(context *gin.Context) {

	var request Request.UpdatePlayerLevelRequest
	context.Bind(&request)

	response, errorResponse := controller.Service.UpdatePlayerLevel(&request)
	if response {

		context.JSON(http.StatusOK, errorResponse)
	} else {

		context.JSON(http.StatusConflict, errorResponse)
	}
}

//DeletePlayerLevel Metodo DELETE para eliminar un PlayerLevel
func (controller *PlayerLevelController) DeletePlayerLevel(context *gin.Context) {

	var request Request.DeletePlayerLevelRequest
	context.Bind(&request)

	response, errorResponse := controller.Service.DeletePlayerLevel(&request)
	if response {

		context.JSON(http.StatusOK, errorResponse)
	} else {

		context.JSON(http.StatusConflict, errorResponse)
	}
}
