package addresscontroller

import (
	"github.com/gin-gonic/gin"
)

func GetViaCepPostalCode(ctx *gin.Context){
	postalcode := ctx.Params.ByName("postalcode")

	viaceprepository.ApiGetByPostalCode(postalcode)
}