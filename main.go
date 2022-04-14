package main

import (
	"album"
	"artist"
	"db"
	"favourite"
	"fmt"
	"log"
	"net/http"
	"playlist"
	"register"
	"time"
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
			c.Set("role", claims["role"])
			return
		}

		c.AbortWithStatus(http.StatusUnauthorized)
	}
}

func AuthorizeAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		role := c.MustGet("role")
		if role == "General" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}
func WelcomeEndpoint(c *gin.Context) {
	c.Next()
	dt := time.Now()
	log.Printf("Current date and time is: ", dt.String())
	fmt.Println("Formated date and time", dt.Format("01-02-2006 15:04:05"))

	log.Printf("Endpoint URL is %v", c.Request.URL)

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
	router.GET("/login", registerRepo.Login)

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

	// Basic auth middleware to display Current Date, time and url path and show version detail
	router.Use(WelcomeEndpoint)

	router.GET("v1/welcome", func(c *gin.Context) {
		if c.FullPath() == "/v1/welcome" {
			log.Printf("welcome to older version")
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Middleware version v1",
		})
	})

	router.GET("v2/hello-world ", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World version 2",
		})
	})
	//Adminauthorized wrapper class

	authorized := router.Group("/api", AuthorizeJWT())
	authorized.Use(AuthorizeAdmin())
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
		authorized.DELETE("/v1/album/remove-track", albumRepo.RemoveAlbum)

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

	router.GET("/v1/album/show", albumRepo.Read)
	router.GET("/v1/track/show", trackRepo.Read)
	router.GET("/v1/playlist/show", playlistRepo.Read)
	router.GET("/v1/playlist/get-playlist-track/:id", playlistRepo.PlaylistById)

	router.POST("/v1/favourite-track/create", favRepo.Create)
	router.DELETE("/v1/unfavourite-track", favRepo.Delete)
	router.GET("/v1/favourite-track", favRepo.Read)
	router.GET("/v1/favourite-track/:id", favRepo.FavTrackId)

	//Test ENDPOINTS
	router.GET("/api/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Pong",
		})
	})

	return router
}
