package routes

import (
	addresscontroller "github.com/VieiraGabrielAlexandre/luztecnologia-cms-backend/core/controllers/address_controller"
	clientscontroller "github.com/VieiraGabrielAlexandre/luztecnologia-cms-backend/core/controllers/clients_controller"
	valuescontroller "github.com/VieiraGabrielAlexandre/luztecnologia-cms-backend/core/controllers/values_controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"} // Update this with your actual allowed origins
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"*"} // Allow any headers for simplicity; you can customize this based on your needs

	r.Use(cors.New(config))

	r.OPTIONS("/api/clients", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Status(200)
	})

	base := r.Group("/api/")
	{
		clients := base.Group("/clients")
		{
			clients.GET("", clientscontroller.List)
			clients.GET("/:id", clientscontroller.Detail)
			clients.POST("", clientscontroller.Create)
		}

		base.GET("/address/:postalcode", addresscontroller.GetViaCepPostalCode)
		base.GET("/values", valuescontroller.Calculate)
	}

	r.Run(":3001")
}
