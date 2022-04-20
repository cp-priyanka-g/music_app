package main

import (
	// "db"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	fmt.Println("Hello World!")

	router.Use(WelcomeEndpoint)

	router.GET("/v1/welcome", func(c *gin.Context) {
		if c.FullPath() == "/v1/welcome" {
			log.Printf("welcome to older version")
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Middleware version v1",
		})
	})

	router.GET("/v2/hello-world ", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World version 2",
		})
	})

	//Test ENDPOINTS
	router.GET("/api/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Pong",
		})
	})

	router.Run(":8080")
}
func WelcomeEndpoint(c *gin.Context) {
	c.Next()
	dt := time.Now()
	log.Printf("Current date and time is: ", dt.String())
	fmt.Println("Formated date and time", dt.Format("01-02-2006 15:04:05"))
	log.Printf("Endpoint URL is %v", c.Request.URL)

}
