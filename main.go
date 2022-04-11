package main

import (
	"album"
	"artist"
	"db"
	"favourite"
	"fmt"
	"net/http"
	"playlist"
	"register"
	"track"

	"github.com/jmoiron/sqlx"

	"github.com/gin-gonic/gin"

	"github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	sqlDb := db.NewSql()
	defer sqlDb.Close()

	r := setupRouter(sqlDb)
	_ = r.Run(":8080")
}

//Middleware
func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Token")
		tokenString := authHeader
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Invalid token", token.Header["Token"])

			}
			return "secretKey", nil
		})

		if err != nil {
			c.JSON(403, gin.H{"message": "Your Token has been expired."})
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if claims["role"] == "Admin" {
				c.Request.Header.Set("Role", "Admin")
				return

			} else if claims["role"] == "General" {
				c.Request.Header.Set("Role", "General")
				return

			}
		}
		fmt.Println(err)
		c.AbortWithStatus(http.StatusUnauthorized)

	}
}

//Middleware for user
func UserAuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Token")
		tokenString := authHeader
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Invalid token", token.Header["Token"])

			}
			return "secretKey", nil
		})

		if err != nil {
			c.JSON(403, gin.H{"message": "Your Token has been expired."})
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if claims["role"] == "General" {
				c.Request.Header.Set("Role", "General")
				return

			}
		}
		fmt.Println(err)
		c.AbortWithStatus(http.StatusUnauthorized)

	}
}

func setupRouter(sqlDb *sqlx.DB) *gin.Engine {
	router := gin.Default()

	registerRepo := register.New(sqlDb)
	artistRepo := artist.New(sqlDb)
	albumRepo := album.New(sqlDb)
	trackRepo := track.New(sqlDb)
	playlistRepo := playlist.New(sqlDb)
	favRepo := favourite.New(sqlDb)

	//USER Authencation
	router.POST("/api/v1/register", registerRepo.Register)
	router.POST("/api/v1/register/admin-register", registerRepo.RegisterAdmin)
	router.GET("/api/v1/login", registerRepo.Login)

	//Adminauthorized wrapper class

	authorized := router.Group("/api", AuthorizeJWT())
	{
		authorized.POST("/v1/artist", artistRepo.Create)
		authorized.PUT("/v1/artist/:id", artistRepo.Update)
		authorized.GET("/v1/artist/show", artistRepo.Read)
		authorized.DELETE("/v1/artist/remove", artistRepo.Delete)

		//ALBUM
		authorized.POST("/v1/album", albumRepo.Create)
		authorized.PUT("/v1/album/edit", albumRepo.Update)
		authorized.DELETE("/v1/album/remove", albumRepo.Delete)
		authorized.POST("/v1/album/add", albumRepo.AddAlbum)
		authorized.DELETE("/api/v1/album/remove-track", albumRepo.RemoveAlbum)

		//Track
		authorized.POST("/v1/track", trackRepo.Create)
		authorized.PUT("/v1/track/edit", trackRepo.Update)
		authorized.DELETE("/v1/track/remove", trackRepo.Delete)

		//Playlist
		authorized.POST("/v1/playlist", playlistRepo.Create)
		authorized.PUT("/v1/playlist/edit", playlistRepo.Update)
		authorized.DELETE("/v1/playlist/remove", playlistRepo.Delete)

		authorized.POST("/v1/playlist/add-track-playlist", playlistRepo.AddPlaylist)
		authorized.DELETE("/v1/playlist/remove-track-playlist", playlistRepo.Remove)

	}

	//Favourite Track (User functionality API)

	authorizedUser := router.Group("/api", AuthorizeJWT())
	{
		authorizedUser.GET("/v1/album/show", albumRepo.Read)
		authorizedUser.GET("/v1/track/show", trackRepo.Read)
		authorizedUser.GET("/v1/playlist/show", playlistRepo.Read)
		authorizedUser.GET("/v1/playlist/get-playlist-track/:id", playlistRepo.PlaylistById)

		authorizedUser.POST("/v1/favourite-track/create", favRepo.Create)
		authorizedUser.DELETE("/v1/unfavourite-track", favRepo.Delete)
		authorizedUser.GET("/v1/favourite-track", favRepo.Read)
		authorizedUser.GET("/v1/favourite-track/:id", favRepo.FavTrackId)
	}

	//Test ENDPOINTS
	router.GET("/api/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Pong",
		})
	})

	return router
}
