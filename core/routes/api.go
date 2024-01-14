package routes

import (
	clientscontroller "github.com/VieiraGabrielAlexandre/luztecnologia-cms-backend/core/controllers/clients_controller"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()

	clients := r.Group("/api/")
	{
		clients.GET("clients/", clientscontroller.List)
		clients.GET("clients/", clientscontroller.List)
	}

	r.Run(":8888")
}
