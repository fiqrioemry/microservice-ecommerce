package routes

import (
	"github.com/fiqrioemry/microservice-ecommerce/server/user-service/internal/handlers"

	"github.com/gin-gonic/gin"
)

func LocationRoutes(r *gin.Engine, handler *handlers.LocationHandler) {
	r.GET("/api/provinces", handler.GetProvinces)
	r.GET("/api/provinces/:id/cities", handler.GetCitiesByProvince)
}
