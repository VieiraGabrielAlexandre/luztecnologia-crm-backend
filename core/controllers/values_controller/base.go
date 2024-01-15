package valuescontroller

import (
	"net/http"

	"github.com/VieiraGabrielAlexandre/luztecnologia-cms-backend/core/models"
	calculateservices "github.com/VieiraGabrielAlexandre/luztecnologia-cms-backend/core/services/calculate_services"
	"github.com/gin-gonic/gin"
)

type CalculateResponse struct {
	Message string      `json:"message"`
	Data    models.Calc `json:"data"`
}

func Calculate(ctx *gin.Context) {
	var calc models.Calc

	if err := ctx.ShouldBindJSON(&calc); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	calculateservices.Base(&calc)

	response := CalculateResponse{
		Message: "Valor calculado com sucesso",
		Data: models.Calc{
			NumberClients: calc.NumberClients,
			Value:         calc.Value,
		},
	}

	ctx.JSON(http.StatusOK, response)
}
