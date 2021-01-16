package web

import (
	"snapmatic-to-jpg/src/business"

	"github.com/gin-gonic/gin"
)

// Service contains the dependencies
type Service struct {
	router *gin.Engine
}

// Start starts the application
func Start() error {
	var service Service

	service.router = InitRouter()
	service.RegisterRoutes()

	return service.router.Run(":8080")
}

// InitRouter initializes the Gin router
func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())

	// gin.SetMode(gin.DebugMode)
	gin.SetMode(gin.ReleaseMode)

	return router
}

// RegisterRoutes registers all available endpoints
func (service *Service) RegisterRoutes() {
	var (
		baseEndpoint    = "/api/"
		convertEndpoint = baseEndpoint + "convert"
		healthEndpoint  = baseEndpoint + "health"
	)

	service.router.GET(healthEndpoint, business.Health)

	// The snapmatic files usually are less than 1 MB but buffer for 2 just to be safe
	service.router.MaxMultipartMemory = 2 << 20 // 2 MiB
	service.router.POST(convertEndpoint, business.Convert)
}
