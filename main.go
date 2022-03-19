package main

import (
	"artist"
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
	artistRepo := artist.New(sqlDb)

	router.POST("/api/v1/register", registerRepo.AddUser)
	router.POST("/api/v1/register/admin_register", registerRepo.AddAdmin)
	router.POST("/api/v1/login", registerRepo.Login)
	// router.GET("/api/v1/elcome", registerRepo.Welcome)
	router.POST("/api/v1/artist/add", artistRepo.CreateArtist)
	router.POST("/api/v1/artist/update/:id", artistRepo.UpdateArtist)
	router.POST("/api/v1/artist/delete/", artistRepo.DeleteArtist)

	router.GET("/api/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Pong",
		})
	})

	return router
}
