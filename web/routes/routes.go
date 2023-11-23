package routes

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/larnTechGeeks/pc-test/internal/services/spams"
	spam_handler "github.com/larnTechGeeks/pc-test/web/handlers/spams"
	middleware "github.com/larnTechGeeks/pc-test/web/middlewares"
)

const (
	maxLatencyDuration = 5 * time.Second
)

type appRouter struct {
	*gin.Engine
}

func BuildRouter() *appRouter {

	router := gin.Default()

	defaultMiddlewares := middleware.DefaultMiddlewares()

	router.Use(defaultMiddlewares...)

	apiRouter := router.Group("/v1")

	spamService := spams.NewSpamService()

	spam_handler.AddOpenEndpoint(apiRouter, spamService)

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"error_message": "Endpoint not found"})
	})

	return &appRouter{
		router,
	}
}
