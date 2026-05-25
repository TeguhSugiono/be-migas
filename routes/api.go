package routes

import (
	"BackendEsp32/controllers"

	"BackendEsp32/middleware"

	"github.com/gin-gonic/gin"
)

func ApiRoutes(api *gin.RouterGroup) {

	api.POST("/gas", controllers.StoreGas)
	// api.POST("/register", controllers.Register)
	// api.POST("/login", controllers.Login)

	api.POST("/register", controllers.Register)
	api.POST("/login", controllers.Login)

	auth := api.Group("/auth")

	auth.Use(middleware.AuthMiddleware())

	{
		auth.GET("/me", controllers.Me)
	}

	device := api.Group("/devices")

	device.Use(middleware.AuthMiddleware())

	{
		device.POST(
			"",
			middleware.RoleMiddleware(
				"SELLER",
				"SUPER_ADMIN",
			),
			controllers.CreateDevice,
		)

		device.POST(
			"/pair",
			middleware.RoleMiddleware(
				"CUSTOMER",
			),
			controllers.PairDevice,
		)

		device.GET(
			"/my",
			middleware.RoleMiddleware(
				"CUSTOMER",
			),
			controllers.MyDevices,
		)

		device.GET(
			"/:device_id/latest",
			controllers.LatestDeviceStatus,
		)

		device.GET(
			"/:device_id/history",
			controllers.DeviceHistory,
		)
	}

}
