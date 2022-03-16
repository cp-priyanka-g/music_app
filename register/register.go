package register

import (
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
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
// func (repository *RegisterRepository) Careers(c *gin.Context) {

// 	userInfo, err := repository.GetUser()

// 	if err != nil {
// 		c.AbortWithStatus(http.StatusInternalServerError)
// 		return
// 	}

// 	c.JSON(http.StatusOK,userInfo)

// }

// func (repository *RegisterRepository) GetUser() (userInfo[] UserRegister, err error) {

// 	err = repository.Db.Select(&userInfo, `SELECT user_id,email,password FROM Users`)

// 	if err != nil {
// 		 err.Error()
// 		return
// 	}

// 	return
// }

func (repository *RegisterRepository) Login(c *gin.Context) {
	input := UserRegister{}
	var userInfo []string

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}

	err := repository.Db.Select(&userInfo, `SELECT user_id,email,password FROM Users`)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "Cannot get user info ")
		return
	}
	//compare the user from the request, with the one we defined:
	if userInfo[2] != input.Email || userInfo[3] != input.Password {
		c.JSON(http.StatusUnauthorized, "Please provide valid login details")
		return
	}
	token, err := CreateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	c.JSON(http.StatusOK, token)
}

//Create Token

func CreateToken(userid uint64) (string, error) {
	var err error
	//Creating Access Token
	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd") //this should be in an env file
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userid
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}
