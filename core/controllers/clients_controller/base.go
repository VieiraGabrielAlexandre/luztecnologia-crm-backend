package clientscontroller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/VieiraGabrielAlexandre/luztecnologia-cms-backend/core/models"
	clientsrepository "github.com/VieiraGabrielAlexandre/luztecnologia-cms-backend/core/repository/client"
	"github.com/VieiraGabrielAlexandre/luztecnologia-cms-backend/core/utils"
)

func List(ctx *gin.Context) {
	page, limit, offset := utils.Paginate(ctx)

	count := clientsrepository.Count()

	clients := clientsrepository.GetAll()

	response := gin.H{
		"message": "Listagem de clientes",
		"data":    clients,
		"meta":    utils.GenerateMetadata(int(count), page, limit, offset),
	}

	ctx.JSON(http.StatusOK, response)
}

func Detail(ctx *gin.Context) {
	id := ctx.Params.ByName("id")

	clientID, err := strconv.Atoi(id)

	client, err := clientsrepository.Get(clientID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Detalhes de um cliente: %s", id),
		"data":    client,
	})
}

func Create(ctx *gin.Context) {
	var clientModel models.Client

	if err := ctx.ShouldBindJSON(&clientModel); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	message, erro := clientsrepository.Save(&clientModel)

	if erro != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": erro.Error(),
		})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{"message": message})
}
