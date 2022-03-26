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

authMiddleware := &jwt.GinJWTMiddleware{
	Realm:      "test zone",
	Key:        []byte("secret key"),
	Timeout:    time.Hour,
	MaxRefresh: time.Hour,
	Authenticator: func(userId string, password string, c *gin.Context) (string, bool) {
		if (userId == "admin" && password == "admin") || (userId == "test" && password == "test") {
			return userId, true
		}

		return userId, false
	},
	Authorizator: func(userId string, c *gin.Context) bool {
		if userId == "admin" {
			return true
		}

		return false
	},
	Unauthorized: func(c *gin.Context, code int, message string) {
		c.JSON(code, gin.H{
			"code":    code,
			"message": message,
		})
	},

	TokenLookup: "header:Authorization",
	TokenHeadName: "Bearer",
	TimeFunc: time.Now,
}

func setupRouter(sqlDb *sqlx.DB) *gin.Engine {
	router := gin.Default()

	//router.Use(AdminAuthorize)

	//authorized := router.Group("/", basicAuth)

	auth := router.Group("/auth")
    auth.Use(authMiddleware.MiddlewareFunc())

	registerRepo := register.New(sqlDb)
	artistRepo := artist.New(sqlDb)
	albumRepo := album.New(sqlDb)
	trackRepo := track.New(sqlDb)
	playlistRepo := playlist.New(sqlDb)
	favRepo := favourite.New(sqlDb)

	//USER Authencation
	router.POST("/api/v1/register", registerRepo.Register)
	router.POST("/api/v1/register/admin-register", registerRepo.RegisterAdmin)
	router.POST("/api/v1/login", registerRepo.Login)

	//ARTIST
	auth.POST("/api/v1/artist", artistRepo.Create)
	auth.PUT("/api/v1/artist/:id", artistRepo.Update)
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
