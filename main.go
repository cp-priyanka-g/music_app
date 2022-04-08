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

// func AuthorizeJWT() gin.HandlerFunc {
// 	return func(c *gin.Context) {

// 		authHeader := c.GetHeader("Authorization")
// 		tokenString := authHeader
// 		token, err := register.JWTAuthService().ValidateToken(tokenString)
// 		if token.Valid {
// 			claims := token.Claims.(jwt.MapClaims)
// 			fmt.Println(claims)
// 		} else {
// 			fmt.Println(err)
// 			c.AbortWithStatus(http.StatusUnauthorized)
// 		}

// 	}
// }

func setupRouter(sqlDb *sqlx.DB) *gin.Engine {
	router := gin.Default()

	registerRepo := register.New(sqlDb)
	// artistRepo := artist.New(sqlDb)
	// albumRepo := album.New(sqlDb)
	// trackRepo := track.New(sqlDb)
	// playlistRepo := playlist.New(sqlDb)
	// favRepo := favourite.New(sqlDb)

	//USER Authencation
	router.POST("/api/v1/register", registerRepo.Register)
	router.POST("/api/v1/register/admin-register", registerRepo.RegisterAdmin)
	router.POST("/api/v1/login", registerRepo.Login)

	//Adminauthorized wrapper class
	// authorized := router.Group("/api", Adminauthorized())
	// {
	// 	authorized.POST("/v1/artist", artistRepo.Create)
	// 	authorized.PUT("/v1/artist/:id", artistRepo.Update)
	// 	authorized.GET("/v1/artist/show", artistRepo.Read)
	// 	authorized.DELETE("/v1/artist/remove", artistRepo.Delete)

	// 	//ALBUM
	// 	authorized.POST("/v1/album", albumRepo.Create)
	// 	authorized.PUT("/v1/album/edit", albumRepo.Update)
	// 	authorized.DELETE("/v1/album/remove", albumRepo.Delete)
	// 	authorized.POST("/v1/album/add", albumRepo.AddAlbum)
	// 	authorized.DELETE("/api/v1/album/remove-track", albumRepo.RemoveAlbum)

	// 	//Track
	// 	authorized.POST("/v1/track", trackRepo.Create)
	// 	authorized.PUT("/v1/track/edit", trackRepo.Update)
	// 	authorized.DELETE("/v1/track/remove", trackRepo.Delete)

	// 	//Playlist
	// 	authorized.POST("/v1/playlist", playlistRepo.Create)
	// 	authorized.PUT("/v1/playlist/edit", playlistRepo.Update)
	// 	authorized.DELETE("/v1/playlist/remove", playlistRepo.Delete)

	// 	authorized.POST("/v1/playlist/add-track-playlist", playlistRepo.AddPlaylist)
	// 	authorized.POST("/v1/playlist/get-playlistby-track", playlistRepo.GetPlaylistTrack)
	// 	authorized.DELETE("/v1/playlist/remove-track-playlist", playlistRepo.Remove)

	// }

	// //Favourite Track (User functionality API)
	// authorizedUser := router.Group("/api", Userauthorized())
	// {
	// 	authorizedUser.GET("/v1/album/show", albumRepo.Read)
	// 	authorizedUser.GET("/v1/track/show", trackRepo.Read)
	// 	authorizedUser.GET("/v1/playlist/show", playlistRepo.Read)
	// 	authorizedUser.GET("/v1/playlist/get-playlist-track/:id", playlistRepo.PlaylistById)

	// 	authorizedUser.POST("/v1/favourite-track/create", favRepo.Create)
	// 	authorizedUser.DELETE("/v1/unfavourite-track", favRepo.Delete)
	// 	authorizedUser.GET("/v1/favourite-track", favRepo.Read)
	// 	authorizedUser.GET("/v1/favourite-track/:id", favRepo.FavTrackId)
	// }

	//Test ENDPOINTS
	router.GET("/api/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Pong",
		})
	})

	return router
}
