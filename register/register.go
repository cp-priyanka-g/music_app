package register

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type UserRegister struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRepository struct {
	Db *sqlx.DB
}
type UserLogin struct {
	Email    string
	Password string
}

func New(db *sqlx.DB) *RegisterRepository {
	return &RegisterRepository{Db: db}
}

func (repository *RegisterRepository) AddUser(c *gin.Context) {

	input := UserRegister{}
	err := c.ShouldBindWith(&input, binding.JSON)

	if err != nil {
		c.Abort()
		c.JSON(http.StatusBadRequest, gin.H{"message cannot bind the STRUCT ": err.Error()})
		return
	}

	_, err = repository.Db.Exec(`INSERT INTO Users(name,email,password,user_type) VALUES (?,?,?,?)`, input.Name, input.Email, input.Password, "General")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message cannot insert ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Register Successfully"})

}

// ADMIN REGISTER
func (repository *RegisterRepository) AddAdmin(c *gin.Context) {

	input := UserRegister{}

	err := c.ShouldBindWith(&input, binding.JSON)

	if err != nil {
		c.Abort()
		c.JSON(http.StatusBadRequest, gin.H{"Message cannot bind the STRUCT ": err.Error()})
		return
	}

	_, err = repository.Db.Exec(`INSERT INTO Users(name,email,password,user_type) VALUES (?,?,?,?)`, input.Name, input.Email, input.Password, "Admin")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message cannot insert ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Admin Registered Successfully"})
}

// LOGIN

func (repository *RegisterRepository) Login(c *gin.Context) {

	input := UserRegister{}
	//var email, pass string
	var userInfo UserRegister

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}

	row, err := repository.Db.Query(`SELECT email,password FROM Users WHERE email=? AND password=?`, input.Email, input.Password)
	row.Scan(&userInfo.Email, &userInfo.Password)

	c.JSON(http.StatusOK, row)
	//compare the user from the request, with the one we defined:

	if input.Email != userInfo.Email || input.Password != userInfo.Password {
		c.JSON(http.StatusUnauthorized, "Please provide valid login details")
		return
	} else if err != nil {
		c.JSON(http.StatusUnauthorized, "Please register to login")

	}

	c.JSON(http.StatusOK, row)
}
