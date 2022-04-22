package main

import (
	// "db"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func WelcomeEndpoint(c *gin.Context) {
	c.Next()
	dt := time.Now()
	fmt.Println("Current date and time is: ", dt.String())
	fmt.Println("Formated date and time", dt.Format("01-02-2006 15:04:05"))
	log.Printf("Endpoint URL is %v", c.Request.URL)

}
func VersionCheck(c *gin.Context) {

	url := c.Request.URL.String()
	fmt.Println("url:", url)
	if strings.Contains(url, "/v1") {
		fmt.Println("This is old version V1")
	}

}

func main() {
	router := gin.Default()
	authorized := router.Group("/api", WelcomeEndpoint)
	{
		authorized.Use(VersionCheck)

		authorized.GET("/v1/welcome", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Welcome to Middleware version v1",
			})
		})

		authorized.GET("/v2/hello-world", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Hello World version 2",
			})
		})

	}

	//Test ENDPOINTS
	router.GET("/api/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Pong",
		})
	})

	router.Run(":8080")
}
