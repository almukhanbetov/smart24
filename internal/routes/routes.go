package routes

import (
	"github.com/almukhanbetov/smart24/backend/internal/handlers"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Setup(pool *pgxpool.Pool) *gin.Engine {
	r := gin.Default()

	r.GET("/health", handlers.Health)

	api := r.Group("/api")
	{
		api.GET("/users", handlers.GetUsers(pool))
		api.GET("/devices", handlers.GetDevices(pool))
		api.GET("/device-set", handlers.GetDeviceSet(pool))
		api.GET("/tariffs", handlers.GetTariffs(pool))
		api.GET("/payments", handlers.GetPayments(pool))
		api.GET("/coin", handlers.GetCoin(pool))
		api.GET("/money", handlers.GetMoney(pool))
		api.GET("/devices/:account/full", handlers.GetDeviceFull(pool))
		api.GET("/dashboard", handlers.GetDashboard(pool))
		api.GET("/max", handlers.GetMaxData(pool))
	}

	return r
}
