package main

import (
	// "db"
	"fmt"
	"log"
	"net/http"

	// "register"
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	sqlDb := db.NewSql()
	defer sqlDb.Close()

	r := setupRouter(sqlDb)
	_ = r.Run(":8080")
}

func WelcomeEndpoint(c *gin.Context) {
	c.Next()
	dt := time.Now()
	log.Printf("Current date and time is: ", dt.String())
	fmt.Println("Formated date and time", dt.Format("01-02-2006 15:04:05"))
	log.Printf("Endpoint URL is %v", c.Request.URL)

}

func setupRouter(sqlDb *sqlx.DB) *gin.Engine {
	router := gin.Default()
	registerRepo := register.New(sqlDb)

	//USER Authencation
	router.POST("/api/v1/register", registerRepo.Register)
	router.POST("/api/v1/register/admin-register", registerRepo.RegisterAdmin)
	router.GET("/login", registerRepo.Login)

	// Basic auth middleware to display Current Date, time and url path and show version detail
	router.Use(WelcomeEndpoint)

	router.GET("v1/welcome", func(c *gin.Context) {
		if c.FullPath() == "/v1/welcome" {
			log.Printf("welcome to older version")
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Middleware version v1",
		})
	})

	router.GET("v2/hello-world ", func(c *gin.Context) {
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

	return router
}
