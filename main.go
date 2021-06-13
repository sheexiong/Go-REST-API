package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sheexiong/Go-REST-API/bin/handlers"
	limit "github.com/yangxikun/gin-limit-by-key"
	"golang.org/x/time/rate"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

func main() {
	r := gin.Default()

	r.Use(Cors())

	// REST API Throttling
	r.Use(limit.NewRateLimiter(func(c *gin.Context) string {
		return c.ClientIP() // limit rate by client ip
	}, func(c *gin.Context) (*rate.Limiter, time.Duration) {
		return rate.NewLimiter(rate.Every(100*time.Millisecond), 10), time.Hour // limit 10 qps/clientIp and permit bursts of at most 10 tokens, and the limiter liveness time duration is 1 hour
	}, func(c *gin.Context) {
		c.AbortWithStatus(429) // handle exceed rate limit request
	}))

	// REST API Versioning
	v1 := r.Group("api/v1")
	{
		v1.GET("/properties", handlers.GetProperties)
		v1.GET("/properties/:id", handlers.GetProperty)
		v1.GET("/countries", handlers.GetCountries)
		v1.GET("/countries/:id", handlers.GetCountry)
		v1.POST("/property", handlers.PostProperty)
		v1.POST("/country", handlers.PostCountry)
		v1.PUT("/properties/:id", handlers.UpdateProperty)
		v1.DELETE("/properties/:id", handlers.DeleteProperty)
	}

	r.Run(":8080")
}
