package http

import (
	"gostat/services/stat"

	"github.com/gin-gonic/gin"

	middlewareAuth "gostat/pkg/middleware/auth"
)

func RegisterHTTPEndpoints(router *gin.Engine, uc stat.UseCase) {
	h := NewHandler(uc)

	statEndpoints := router.Group("/api/stat")
	statEndpoints.Use(middlewareAuth.JwtAuthMiddleware())
	{
		statEndpoints.GET("/visits", h.GetVisits)
		statEndpoints.GET("/links", h.GetLinks)
	}

	statUpdateEndpoints := router.Group("/api/stat/update")
	{
		statUpdateEndpoints.PUT("/visit", h.SetVisit)
		statUpdateEndpoints.PUT("/visit/extend", h.VisitExtend)
	}
}
