package routes

import (
	clientscontroller "github.com/VieiraGabrielAlexandre/luztecnologia-cms-backend/core/controllers/clients_controller"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()

	clients := r.Group("/api/")
	{
		clientsGroup := base.Group("clients")
		{
			clientsGroup.GET("/", clientscontroller.List)
			clientsGroup.GET("/:id", clientscontroller.Detail)
		}
	}

	r.Run(":8888")
}
