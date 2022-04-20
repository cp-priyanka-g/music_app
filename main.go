package main

import (
	"album"
	"db"

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
	albumRepo := album.New(sqlDb)

	//basic auth for user authentication

	adminAuth := gin.BasicAuth(gin.Accounts{
		"priyanka@gmail.com": "123",
	})

	adminauthorized := router.Group("/", adminAuth)
	{

		adminauthorized.POST("/album", albumRepo.Create)
		adminauthorized.PUT("/album/edit", albumRepo.Update)
		adminauthorized.DELETE("/album/remove", albumRepo.Delete)

	}

	basicAuth := gin.BasicAuth(gin.Accounts{
		"priyanka@gmail.com": "123",
		"anisha@gmail.com":   "1234",
	})

	basicauthorized := router.Group("/", basicAuth)
	{

		basicauthorized.GET("/album/show", albumRepo.Read)

	}

	//Test ENDPOINTS
	router.GET("/api/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Pong",
		})
	})

	return router
}
