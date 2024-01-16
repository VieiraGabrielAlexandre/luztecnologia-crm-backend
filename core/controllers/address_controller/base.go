package addresscontroller

import (
	"github.com/gin-gonic/gin"
)

func GetViaCepPostCode(ctx *gin.Context){
	postalcode := ctx.Params.ByName("postalcode")

	
}