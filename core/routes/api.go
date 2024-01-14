package routes

import (
	clientscontroller "github.com/VieiraGabrielAlexandre/luztecnologia-cms-backend/core/controllers/clients_controller"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()

	base := r.Group("/api/")
	{
		base.GET("clients/", clientscontroller.List)
	}

	r.Run(":8888")
}
