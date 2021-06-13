package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sheexiong/Go-REST-API/bin/handlers"
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
