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
	var email string

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}

	err := repository.Db.Get(&email, `SELECT email FROM Users WHERE email= ?`, input.Email)

	c.JSON(http.StatusOK, email)
	//compare the user from the request, with the one we defined:

	if input.Email != email {
		c.JSON(http.StatusUnauthorized, "Please provide valid login details")
		return
	} else if err != nil {
		c.JSON(http.StatusUnauthorized, "Please register to login")

	}

	c.JSON(http.StatusOK, gin.H{"message": "LOgin Successfully"})
}
