package addresscontroller

import (
	"net/http"

	viaceprepository "github.com/VieiraGabrielAlexandre/luztecnologia-cms-backend/core/repository/via_cep_repository"
	"github.com/gin-gonic/gin"
)

func GetViaCepPostalCode(ctx *gin.Context) {
	postalcode := ctx.Params.ByName("postalcode")

	address, err := viaceprepository.ApiGetByPostalCode(postalcode) // Use the imported package

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": address})
	return
}
