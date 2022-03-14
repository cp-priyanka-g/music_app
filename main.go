package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/reflectx"
)

func main() {

	router := gin.Default()

	router.POST("/api/v1/user/userInsert", userInsert)
	router.GET("/api/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Pong",
		})
	})

	_ = router.Run(":8080")
}

type CareerRepository struct {
	Db *sqlx.DB
}

func New(db *sqlx.DB) *CareerRepository {
	return &CareerRepository{Db: db}
}

type UserRegister struct {
	UserId   int    `json:"user_id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	UserType string `json:"user_type"`
}

func userInsert(c *gin.Context) {
	var db *sqlx.DB

	db = sqlx.MustConnect("mysql", "root:canopas@tcp(127.0.0.1:3306)/musicplayer?parseTime=true")

	db.Mapper = reflectx.NewMapperFunc("json", strings.ToLower)
	db.SetConnMaxLifetime(time.Minute * 1)

	var input UserRegister

	err := c.ShouldBindWith(&input, binding.JSON)
	if err != nil {
		c.Abort()
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	res, err := db.Query(`INSERT INTO Users(name,email,password,user_type) VALUES (?, ?,?,?)`, input)

	if err != nil {
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, res)

}
