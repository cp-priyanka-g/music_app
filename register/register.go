package register

import (
	"fmt"
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

func GenerateToken(uemail string) {
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
		fmt.Println("ERROR IN GENERATING TOKEN")
		return
	}

	fmt.Println("Token Created", tokenString)

}

func New(db *sqlx.DB) *RegisterRepository {
	return &RegisterRepository{Db: db}
}

func (repository *RegisterRepository) Register(c *gin.Context) {

	input := UserRegister{}
	err := c.ShouldBindWith(&input, binding.JSON)

	if err != nil {
		c.Abort()
		c.JSON(http.StatusBadRequest, gin.H{"message cannot bind the STRUCT ": err.Error()})
		return
	}

	_, err = repository.Db.Exec(`INSERT INTO Users(name,email,user_type) VALUES (?,?,?)`, input.Name, input.Email, "General")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message cannot insert ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Register Successfully"})

	GenerateToken(input.Email)

}

//

// ADMIN REGISTER
func (repository *RegisterRepository) RegisterAdmin(c *gin.Context) {

	input := UserRegister{}

	err := c.ShouldBindWith(&input, binding.JSON)

	if err != nil {
		c.Abort()
		c.JSON(http.StatusBadRequest, gin.H{"Message cannot bind the STRUCT ": err.Error()})
		return
	}

	_, err = repository.Db.Exec(`INSERT INTO Users(name,email,user_type) VALUES (?,?,?)`, input.Name, input.Email, "Admin")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message cannot insert ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Admin Registered Successfully"})

	GenerateToken(input.Email)

}

// LOGIN

func (repository *RegisterRepository) Login(c *gin.Context) {

	input := UserRegister{}
	var email, utype string

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}

	_ = repository.Db.Get(&utype, `SELECT user_type FROM Users WHERE email= ?`, input.Email)
	err := repository.Db.Get(&email, `SELECT email FROM Users WHERE email= ?`, input.Email)

	c.JSON(http.StatusOK, email)

	if input.Email != email {
		c.JSON(http.StatusUnauthorized, "Please provide valid login details")
		return
	} else if err != nil {
		c.JSON(http.StatusUnauthorized, "Please register to login")

	}

	// Creating TOken
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: input.Email,
		UserType: utype,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error in creating the JWT": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tokenString)

	//	Verifying user authentication

	tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			c.JSON(http.StatusUnauthorized, gin.H{"Signature Invalid": err.Error()})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error in creating the JWT": err.Error()})
		return
	}
	if !tkn.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"Unauthorised user": err.Error()})
		return
	}

	// Finally, return the welcome message to the user, along with their

	if claims.UserType == "admin" {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome Admin"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome User"})
	}

}
