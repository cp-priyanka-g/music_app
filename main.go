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
	authorized.POST("/api/v1/register/admin_register", registerRepo.RegisterAdmin)
	authorized.POST("/api/v1/login", registerRepo.Login)

	//ARTIST
	authorized.POST("/api/v1/artist/create", artistRepo.Create)
	authorized.PUT("/api/v1/artist/update/:id", artistRepo.Update)
	authorized.DELETE("/api/v1/artist/delete", artistRepo.Delete)
	authorized.GET("/api/v1/artist/display", artistRepo.Read)

	//ALBUM
	authorized.POST("/api/v1/album/create", albumRepo.Create)
	authorized.PUT("/api/v1/album/update", albumRepo.Update)
	authorized.DELETE("/api/v1/album/delete", albumRepo.Delete)
	router.GET("/api/v1/album/read", albumRepo.Read)
	authorized.POST("/api/v1/album/add-track", albumRepo.Add)
	authorized.DELETE("/api/v1/album/remove-track", albumRepo.Remove)

	//Track
	authorized.POST("/api/v1/track/create", trackRepo.Create)
	authorized.PUT("/api/v1/track/update", trackRepo.Update)
	authorized.DELETE("/api/v1/track/delete", trackRepo.Delete)
	router.GET("/api/v1/track/read", trackRepo.Read)

	//Playlist
	authorized.POST("/api/v1/playlist/create", playlistRepo.Create)
	authorized.PUT("/api/v1/playlist/update", playlistRepo.Update)
	authorized.DELETE("/api/v1/playlist/delete", playlistRepo.Delete)
	authorized.GET("/api/v1/playlist/read", playlistRepo.Read)
	authorized.POST("/api/v1/playlist/add-playlist-track", playlistRepo.Add)
	authorized.DELETE("/api/v1/playlist/delete-playlist-track", playlistRepo.Remove)
	router.GET("/api/v1/playlist/get-playlist", playlistRepo.Get)
	authorized.GET("/api/v1/playlist/get-playlist/:id", playlistRepo.PlaylistById)

	// Favourite Track

	router.POST("/api/v1/favourite-track/create", favRepo.Create)
	router.DELETE("/api/v1/favourite-track/delete", favRepo.Delete)
	router.GET("/api/v1/favourite-track/read", favRepo.Read)
	router.GET("/api/v1/favourite-track/:id", favRepo.FavTrackId)

	//Test ENDPOINTS
	router.GET("/api/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Pong",
		})
	})

	return router
}
