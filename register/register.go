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
type Claims struct {
	Username string `json:"username"`
	UserType string `json:"user_type"`
	jwt.StandardClaims
}

// Create the JWT key used to create the signature
var jwtKey = []byte("my_secret_key")

func GenerateToken(uemail string) string {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: uemail,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		panic(err)
	}

	return tokenString

}

func New(db *sqlx.DB) *RegisterRepository {
	return &RegisterRepository{Db: db}
}

func (repository *RegisterRepository) Register(c *gin.Context) {

	input := UserRegister{}
	err := c.ShouldBindWith(&input, binding.JSON)

	if err != nil {
		c.Abort()
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	auth := GenerateToken(input.Email)
	_, err = repository.Db.Exec(`INSERT INTO Users(name,email,user_type,auth_token) VALUES (?,?,?,?)`, input.Name, input.Email, "General", auth)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Register Successfully"})

}

//

// ADMIN REGISTER
func (repository *RegisterRepository) RegisterAdmin(c *gin.Context) {

	input := UserRegister{}

	auth := GenerateToken(input.Email)

	err := c.ShouldBindWith(&input, binding.JSON)

	if err != nil {
		c.Abort()
		c.JSON(http.StatusBadRequest, gin.H{"Message ": err.Error()})
		return
	}

	_, err = repository.Db.Exec(`INSERT INTO Users(name,email,user_type,auth_token) VALUES (?,?,?,?)`, input.Name, input.Email, "Admin", auth)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Admin Registered Successfully"})

}

// LOGIN

func (repository *RegisterRepository) Login(c *gin.Context) {

	input := UserRegister{}
	var email, utype, token string

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}

	_ = repository.Db.Get(&utype, `SELECT user_type FROM Users WHERE email= ?`, input.Email)
	err := repository.Db.Get(&email, `SELECT email FROM Users WHERE email= ? AND auth_token IS NOT NULL`, input.Email)
	_ = repository.Db.Get(&token, `SELECT auth_token FROM Users WHERE email= ?`, input.Email)

	if email != input.Email {
		c.JSON(http.StatusUnauthorized, "Please provide valid login details")
		return
	} else if err != nil {
		panic(err)

	}

	if utype == "Admin" {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome Admin"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome User"})
	}

}
