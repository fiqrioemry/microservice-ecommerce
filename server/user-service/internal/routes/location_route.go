package routes

import (
	"github.com/fiqrioemry/microservice-ecommerce/server/user-service/internal/handlers"

	"github.com/gin-gonic/gin"
)

func LocationRoutes(r *gin.Engine, handler *handlers.LocationHandler) {
	location := r.Group("/api/locations")
	{
		location.GET("/provinces", handler.GetProvinces)
		location.GET("/provinces/search", handler.SearchProvincesByName)
		location.GET("/provinces/:provinceId/cities", handler.GetCitiesByProvinceID)
		location.GET("/cities/search", handler.SearchCitiesByName)
		location.GET("/cities/:cityId/districts", handler.GetDistrictsByCityID)
		location.GET("/districts/:districtId/subdistricts", handler.GetSubdistrictsByDistrictID)
		location.GET("/subdistricts/:subdistrictId/postalcodes", handler.GetPostalCodesBySubdistrictID)
	}
}
