package routes

import (
	"time"

	addresscontroller "github.com/VieiraGabrielAlexandre/luztecnologia-cms-backend/core/controllers/address_controller"
	clientscontroller "github.com/VieiraGabrielAlexandre/luztecnologia-cms-backend/core/controllers/clients_controller"
	valuescontroller "github.com/VieiraGabrielAlexandre/luztecnologia-cms-backend/core/controllers/values_controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"POST", "GET", "PUT", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization", "Accept", "User-Agent", "Cache-Control", "Pragma"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowCredentials = true
	config.MaxAge = 12 * time.Hour

	r.Use(cors.New(config))
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

	r.Run(":3001")
}
