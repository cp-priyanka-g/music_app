package main

import (
	"album"
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
	albumRepo := album.New(sqlDb)

	router.POST("/api/v1/register", registerRepo.Register)
	router.POST("/api/v1/register/admin_register", registerRepo.RegisterAdmin)
	router.POST("/api/v1/login", registerRepo.Login)

	//ARTIST
	router.POST("/api/v1/artist/create", artistRepo.Create)
	router.PUT("/api/v1/artist/update/:id", artistRepo.Update)
	router.DELETE("/api/v1/artist/delete", artistRepo.Delete)
	router.GET("/api/v1/artist/display", artistRepo.Read)

	//ALBUM
	router.POST("/api/v1/album/create", albumRepo.Create)
	//router.POST("/api/v1/album/update", albumRepo.Update)
	//router.POST("/api/v1/album/delete", albumRepo.Delete)
	//router.POST("/api/v1/album/read", albumRepo.Read)

	router.GET("/api/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Pong",
		})
	})

	return router
}
