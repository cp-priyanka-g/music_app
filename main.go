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

	"github.com/dgrijalva/jwt-go"
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

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer"
		authHeader := c.GetHeader("Authorization")
		tokenString := authHeader[len(BEARER_SCHEMA):]
		token, err := register.JWTAuthService().ValidateToken(tokenString)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			fmt.Println(claims)
		} else {
			fmt.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}

	}
}
func setupRouter(sqlDb *sqlx.DB) *gin.Engine {
	router := gin.Default()
	//router.Use(AuthorizeJWT)

	var loginService register.LoginService = register.StaticLoginService()
	var jwtService register.JWTService = register.JWTAuthService()
	var loginController register.LoginController = register.LoginHandler(loginService, jwtService)

	router.POST("/api/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	})

	registerRepo := register.New(sqlDb)
	artistRepo := artist.New(sqlDb)
	albumRepo := album.New(sqlDb)
	trackRepo := track.New(sqlDb)
	playlistRepo := playlist.New(sqlDb)
	favRepo := favourite.New(sqlDb)

	//USER Authencation
	router.POST("/api/v1/register", registerRepo.Register)
	router.POST("/api/v1/register/admin-register", registerRepo.RegisterAdmin)
	//router.POST("/api/v1/login", registerRepo.Login)

	//ARTISTAdminauthorized
	apiRoutes := router.Group("/api", AuthorizeJWT())
	{
		apiRoutes.POST("/v1/artist", artistRepo.Create)
		apiRoutes.PUT("/v1/artist/:id", artistRepo.Update)

		apiRoutes.GET("/v1/artist/show", artistRepo.Read)
		apiRoutes.DELETE("/v1/artist/remove", artistRepo.Delete)
	}

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
