package main

import (
	"db"
	"register"

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

func setupRouter(sqlDb *sqlx.DB) *gin.Engine {
	router := gin.Default()

	registerRepo := register.New(sqlDb)

	router.POST("/api/v1/register", registerRepo.AddUser)
	router.POST("/api/v1/register/admin_register", registerRepo.AddAdmin)
	router.POST("/api/v1/login", registerRepo.Login)

	router.GET("/api/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Pong",
		})
	})

	return router
}
