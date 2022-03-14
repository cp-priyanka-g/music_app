package register

import (
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jmoiron/sqlx"
)

type UserRegister struct {
	UserId   int    `json:"user_id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	UserType string `json:"user_type"`
}

type RegisterRepository struct {
	Db *sqlx.DB
}

func New(db *sqlx.DB) *RegisterRepository {
	return &RegisterRepository{Db: db}
}

func (repository *RegisterRepository) UserInsert(c *gin.Context) {

	var input UserRegister

	err := c.ShouldBindWith(&input, binding.JSON)
	if err != nil {
		c.Abort()
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	res, err := repository.Db.Query(`INSERT INTO Users(name,email,password,user_type) VALUES (?, ?,?,?)`, input)

	if err != nil {
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, res)

}
