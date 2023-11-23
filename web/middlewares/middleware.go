package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime/debug"
	"strings"
	"time"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/contrib/secure"
	"github.com/gin-gonic/gin"
	"github.com/larnTechGeeks/pc-test/internal/ctxutils"
	"github.com/larnTechGeeks/pc-test/internal/utils"
)

const requestIDHeaderKey = "x-request-id"
const userAgentHeaderKey = "user-agent"
const maxLatency = time.Second * 5

func DefaultMiddlewares() []gin.HandlerFunc {

	return []gin.HandlerFunc{
		// Remails as is
		securer(),
		compressor(),
		corsMiddleware(),

		// comes after compressor, cors and securer
		setRequestId(),

		maintenanceMode(),
		latencyMiddleware(),

		// panic always goes last
		panicRecovery(),
	}
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, X-Auth-Token, X-SET-AUTH-ACCOUNT-ID, X-AUTH-ACCOUNT-ID")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "X-CSRF-Token, Authorization, X-Requested-With, X-AUTH-Token, X-SET-AUTH-ACCOUNT-ID, X-AUTH-ACCOUNT-ID")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func compressor() gin.HandlerFunc {
	return gzip.Gzip(gzip.DefaultCompression)
}

func panicRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				requestID := ctxutils.RequestId(c.Request.Context())

				c.Writer.WriteHeader(http.StatusInternalServerError)
				log.Printf("Failed to recover from panic: %v", err)
				debug.PrintStack()

				if os.Getenv("ENVIRONMENT") != "production" {
					fmt.Fprintf(
						c.Writer,
						`{"error":"panic: %s","details":"See logs for more information (%s)."}`,
						err,
						requestID,
					)
				} else {

					fmt.Fprintf(
						c.Writer,
						`{"error":"internal server error (%s)"}`,
						ctxutils.RequestId(c.Request.Context()),
					)
				}
			}
		}()

		c.Next()
	}
}

func securer() gin.HandlerFunc {
	return secure.Secure(secure.Options{
		SSLRedirect:          strings.ToLower(os.Getenv("FORCE_SSL")) == "true",
		SSLProxyHeaders:      map[string]string{"X-Forwarded-Proto": "https"},
		STSSeconds:           315360000,
		STSIncludeSubdomains: true,
		FrameDeny:            true,
		ContentTypeNosniff:   true,
		BrowserXssFilter:     true,
	})
}

func setRequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestId := utils.GenerateUUID()
		ctx := ctxutils.WithRequestId(c.Request.Context(), requestId)
		c.Request = c.Request.WithContext(ctx)
		c.Header(requestIDHeaderKey, requestId)
		c.Next()
	}
}

func maintenanceMode() gin.HandlerFunc {
	return func(c *gin.Context) {
		if os.Getenv("MAINTENANCE_MODE_ENABLED") == "true" {
			c.JSON(404, gin.H{"message": "Maintenance Mode"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func latencyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		start := time.Now()

		c.Next()

		if time.Since(start) > maxLatency {
			log.Printf(
				"Slow response: %v %v took %v for user: %v",
				c.Request.Method,
				c.Request.URL,
				time.Since(start),
				"anymous", // use actual user id
			)
		}
	}
}
