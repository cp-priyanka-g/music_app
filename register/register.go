package register

import (
	"net/http"

	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type UserRegister struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	UserType string `json:"user_type"`
}

type RegisterRepository struct {
	Db *sqlx.DB
}

type Authentication struct {
	Email string `json:"email"`
}

type Token struct {
	Role        string `json:"role"`
	Email       string `json:"email"`
	TokenString string `json:"token"`
}

func generateRefreshToken() string {
	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["sub"] = 1
	rtClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	rt, err := refreshToken.SignedString([]byte("secret"))

	if err != nil {
		panic(err)
	}

	return rt
}

func New(db *sqlx.DB) *RegisterRepository {
	return &RegisterRepository{Db: db}
}

func (repository *RegisterRepository) Register(c *gin.Context) {

	input := UserRegister{}
	err := c.ShouldBindWith(&input, binding.JSON)
	var email string

	if err != nil {
		c.Abort()
		c.JSON(http.StatusBadRequest, gin.H{"Message ": err.Error()})
		return
	}

	err = repository.Db.Get(&email, `SELECT email from Users where email=?`, input.Email)

	if email != input.Email {

		auth := generateRefreshToken()

		_, err = repository.Db.Exec(`INSERT INTO Users(name,email,user_type,auth_token) VALUES (?,?,?,?)`, input.Name, input.Email, "General", auth)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, gin.H{"message": "Register Successfully"})
		c.JSON(http.StatusOK, auth)

	} else {
		c.JSON(403, gin.H{"message": "Email Already exist"})
	}

}

// ADMIN REGISTER
func (repository *RegisterRepository) RegisterAdmin(c *gin.Context) {

	input := UserRegister{}
	var email string

	err := c.ShouldBindWith(&input, binding.JSON)

	if err != nil {
		c.Abort()
		c.JSON(http.StatusBadRequest, gin.H{"Message ": err.Error()})
		return
	}

	err = repository.Db.Get(&email, `SELECT email from Users where email= ?`, input.Email)

	if email != input.Email {

		auth := generateRefreshToken()

		_, err = repository.Db.Exec(`INSERT INTO Users(name,email,user_type,auth_token) VALUES (?,?,?,?)`, input.Name, input.Email, "Admin", auth)

		c.JSON(http.StatusOK, gin.H{"message": "Admin Registered Successfully"})
		c.JSON(http.StatusOK, auth)

	} else {
		c.JSON(403, gin.H{"message": "Email Already exist"})
	}
}
