package main

import (
	"db"
	"fmt"
	"register"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

const (
	username = "root"
	password = "canopas"
	hostname = "127.0.0.1:3306"
	dbname   = "musicplayer"
)

func dsn(dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
}

func main() {
	sqlDb := db.NewSql()
	defer sqlDb.Close()

	r := setupRouter(sqlDb)

	_ = r.Run(":8080")
}

func setupRouter(sqlDb *sqlx.DB) *gin.Engine {
	router := gin.Default()

	jobsRepo := register.New(sqlDb)

	router.POST("/api/v1/user/userInsert", jobsRepo.UserInsert)

	router.GET("/api/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Pong",
		})
	})

	return router
}
