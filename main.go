package main

import (
	"album"
	"artist"
	"db"
	"favourite"
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

	//auth := router.Group("/auth")

	registerRepo := register.New(sqlDb)
	artistRepo := artist.New(sqlDb)
	albumRepo := album.New(sqlDb)
	trackRepo := track.New(sqlDb)
	playlistRepo := playlist.New(sqlDb)
	favRepo := favourite.New(sqlDb)

	// authorized:=registerRepo.AdminAuthorize
	// Adminauthorized := router.Group("/",authorized)
	// router.Use(authorized)

	//USER Authencation
	router.POST("/api/v1/register", registerRepo.Register)
	router.POST("/api/v1/register/admin-register", registerRepo.RegisterAdmin)
	router.POST("/api/v1/login", registerRepo.Login)

	//ARTISTAdminauthorized
	router.POST("/api/v1/artist", artistRepo.Create)
	router.PUT("/api/v1/artist/:id", artistRepo.Update)
	router.DELETE("/api/v1/artist/remove", artistRepo.Delete)
	router.GET("/api/v1/artist/show", artistRepo.Read)

	//ALBUM
	router.POST("/api/v1/album", albumRepo.Create)
	router.PUT("/api/v1/album/edit", albumRepo.Update)
	router.DELETE("/api/v1/album/remove", albumRepo.Delete)
	router.GET("/api/v1/album/show", albumRepo.Read)
	router.POST("/api/v1/album/add", albumRepo.Add)
	router.DELETE("/api/v1/album/remove-track", albumRepo.Remove)

	//Track
	router.POST("/api/v1/track", trackRepo.Create)
	router.PUT("/api/v1/track/edit", trackRepo.Update)
	router.DELETE("/api/v1/track/remove", trackRepo.Delete)
	router.GET("/api/v1/track/show", trackRepo.Read)

	//Playlist
	router.POST("/api/v1/playlist", playlistRepo.Create)
	router.PUT("/api/v1/playlist/edit", playlistRepo.Update)
	router.DELETE("/api/v1/playlist/remove", playlistRepo.Delete)
	router.GET("/api/v1/playlist/show", playlistRepo.Read)
	router.POST("/api/v1/playlist/add-track-playlist", playlistRepo.Add)
	router.DELETE("/api/v1/playlist/remove-track-playlist", playlistRepo.Remove)
	router.GET("/api/v1/playlist/get", playlistRepo.Get)
	router.GET("/api/v1/playlist/:id", playlistRepo.PlaylistById)

	// Favourite Track

	router.POST("/api/v1/favourite-track/create", favRepo.Create)
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
