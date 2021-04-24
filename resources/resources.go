package resources

import (
	"github.com/baranius/godiator-echo/resources/healthcheck"
	"github.com/gin-gonic/gin"
)

func Register(e *gin.Engine) {
	v1 := e.Group("/v1")

	// Health Check
	healthCheckResource := healthcheck.NewResource()
	v1.GET("/health-check", healthCheckResource.GetHealthCheck)
}