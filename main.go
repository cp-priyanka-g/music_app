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

	// var loginService register.LoginService = register.StaticLoginService()
	// var jwtService register.JWTService = register.JWTAuthService()
	// var loginController register.LoginController = register.LoginHandler(loginService, jwtService)

	// router.POST("/api/login", func(ctx *gin.Context) {
	// 	token := loginController.Login(ctx)
	// 	if token != "" {
	// 		ctx.JSON(http.StatusOK, gin.H{
	// 			"token": token,
	// 		})
	// 	} else {
	// 		ctx.JSON(http.StatusUnauthorized, nil)
	// 	}
	// })

	registerRepo := register.New(sqlDb)
	// artistRepo := artist.New(sqlDb)
	// albumRepo := album.New(sqlDb)
	// trackRepo := track.New(sqlDb)
	// playlistRepo := playlist.New(sqlDb)
	// favRepo := favourite.New(sqlDb)

	//USER Authencation
	router.POST("/api/v1/register", registerRepo.Register)
	router.POST("/api/v1/register/admin-register", registerRepo.RegisterAdmin)

	//Adminauthorized wrapper class
	// apiRoutes := router.Group("/api", AuthorizeJWT())
	// {
	// 	apiRoutes.POST("/v1/artist", artistRepo.Create)
	// 	apiRoutes.PUT("/v1/artist/:id", artistRepo.Update)
	// 	apiRoutes.GET("/v1/artist/show", artistRepo.Read)
	// 	apiRoutes.DELETE("/v1/artist/remove", artistRepo.Delete)

	// 	//ALBUM
	// 	apiRoutes.POST("/v1/album", albumRepo.Create)
	// 	apiRoutes.PUT("/v1/album/edit", albumRepo.Update)
	// 	apiRoutes.DELETE("/v1/album/remove", albumRepo.Delete)
	// 	apiRoutes.GET("/v1/album/show", albumRepo.Read)
	// 	apiRoutes.POST("/v1/album/add", albumRepo.AddAlbum)
	// 	apiRoutes.DELETE("/api/v1/album/remove-track", albumRepo.RemoveAlbum)

	// 	//Track
	// 	apiRoutes.POST("/v1/track", trackRepo.Create)
	// 	apiRoutes.PUT("/v1/track/edit", trackRepo.Update)
	// 	apiRoutes.DELETE("/v1/track/remove", trackRepo.Delete)
	// 	apiRoutes.GET("/v1/track/show", trackRepo.Read)

	// 	//Playlist
	// 	apiRoutes.POST("/v1/playlist", playlistRepo.Create)
	// 	apiRoutes.PUT("/v1/playlist/edit", playlistRepo.Update)
	// 	apiRoutes.DELETE("/v1/playlist/remove", playlistRepo.Delete)
	// 	apiRoutes.GET("/v1/playlist/show", playlistRepo.Read)
	// 	apiRoutes.POST("/v1/playlist/add-track-playlist", playlistRepo.AddPlaylist)
	// 	apiRoutes.DELETE("/v1/playlist/remove-track-playlist", playlistRepo.Remove)
	// 	apiRoutes.GET("/v1/playlist/get", playlistRepo.Get)
	// }

	// Favourite Track (User functionality API)

	// router.POST("/api/v1/favourite-track/create", favRepo.Create)
	// router.DELETE("/api/v1/unfavourite-track", favRepo.Delete)
	// router.GET("/api/v1/favourite-track", favRepo.Read)
	// router.GET("/api/v1/favourite-track/:id", favRepo.FavTrackId)

	//Test ENDPOINTS
	router.GET("/api/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Pong",
		})
	})

	return router
}
