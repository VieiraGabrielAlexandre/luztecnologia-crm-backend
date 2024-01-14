package clientscontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func List(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Listagem de clientes")
}

func Detail(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	ctx.JSON(http.StatusOK, "Detalhes de um cliente: "+id)
}
