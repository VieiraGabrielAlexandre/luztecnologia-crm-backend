package routes

import (
	addresscontroller "github.com/VieiraGabrielAlexandre/luztecnologia-cms-backend/core/controllers/address_controller"
	clientscontroller "github.com/VieiraGabrielAlexandre/luztecnologia-cms-backend/core/controllers/clients_controller"
	valuescontroller "github.com/VieiraGabrielAlexandre/luztecnologia-cms-backend/core/controllers/values_controller"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()

	base := r.Group("/api/")
	{
		clientsGroup := base.Group("clients")
		{
			clientsGroup.GET("/", clientscontroller.List)
			clientsGroup.GET("/:id", clientscontroller.Detail)
			clientsGroup.POST("/", clientscontroller.Create)
		}
		base.POST("calculate", valuescontroller.Calculate)
		base.GET("address/viacep/:postalcode", addresscontroller.GetViaCepPostalCode)

	}

	r.Run(":8888")
}
