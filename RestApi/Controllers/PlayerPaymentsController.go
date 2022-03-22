package Controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PlayerPaymentsController struct {
}

func (controller *PlayerPaymentsController) CalculateTeamPayments(context *gin.Context) {

	context.JSON(http.StatusConflict, "success")
}
