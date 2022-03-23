package main

import (
	"album"
	"artist"
	"db"
	"favourite"
	"net/http"
	"playlist"
	"register"

	"track"

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

	//basicAuth
	basicAuth := gin.BasicAuth(gin.Accounts{
		"priya": "priya",
	})
	authorized := router.Group("/", basicAuth)
	authorized.GET("/welcome", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"secret": "Welcome to the middleware",
		})
	})

	registerRepo := register.New(sqlDb)
	artistRepo := artist.New(sqlDb)
	albumRepo := album.New(sqlDb)
	trackRepo := track.New(sqlDb)
	playlistRepo := playlist.New(sqlDb)
	favRepo := favourite.New(sqlDb)

	//USER Authencation
	authorized.POST("/api/v1/register", registerRepo.Register)
	authorized.POST("/api/v1/register/admin-register", registerRepo.RegisterAdmin)
	authorized.POST("/api/v1/login", registerRepo.Login)

	//ARTIST
	authorized.POST("/api/v1/artist", artistRepo.Create)
	authorized.PUT("/api/v1/artist/:id", artistRepo.Update)
	authorized.DELETE("/api/v1/artist", artistRepo.Delete)
	authorized.GET("/api/v1/artist/", artistRepo.Read)

	//ALBUM
	authorized.POST("/api/v1/album", albumRepo.Create)
	authorized.PUT("/api/v1/album", albumRepo.Update)
	authorized.DELETE("/api/v1/album", albumRepo.Delete)
	router.GET("/api/v1/album", albumRepo.Read)
	authorized.POST("/api/v1/album", albumRepo.Add)
	authorized.DELETE("/api/v1/album", albumRepo.Remove)

	//Track
	authorized.POST("/api/v1/track", trackRepo.Create)
	authorized.PUT("/api/v1/track", trackRepo.Update)
	authorized.DELETE("/api/v1/track", trackRepo.Delete)
	authorized.GET("/api/v1/track", trackRepo.Read)

	//Playlist
	authorized.POST("/api/v1/playlist", playlistRepo.Create)
	authorized.PUT("/api/v1/playlist", playlistRepo.Update)
	authorized.DELETE("/api/v1/playlist", playlistRepo.Delete)
	authorized.GET("/api/v1/playlist", playlistRepo.Read)
	authorized.POST("/api/v1/playlist", playlistRepo.Add)
	authorized.DELETE("/api/v1/playlist", playlistRepo.Remove)
	authorized.GET("/api/v1/playlist", playlistRepo.Get)
	authorized.GET("/api/v1/playlist/:id", playlistRepo.PlaylistById)

	// Favourite Track

	router.POST("/api/v1/favourite-track", favRepo.Create)
	router.DELETE("/api/v1/unfavourite-track", favRepo.Delete)
	router.GET("/api/v1/favourite-track", favRepo.Read)
	router.GET("/api/v1/favourite-track/:id", favRepo.FavTrackId)

	//Test ENDPOINTS
	router.GET("/api/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Pong",
		})
	})

	return router
}
