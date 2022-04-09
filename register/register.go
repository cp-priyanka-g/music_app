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

type Authentication struct {
	Email string `json:"email"`
}

type Token struct {
	Role        string `json:"role"`
	Email       string `json:"email"`
	TokenString string `json:"token"`
}

var secretkey string = "secretkeyjwt"

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

func GenerateJWT(email, role string) (map[string]string, error) {

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["email"] = email
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return nil, err
	}

	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["sub"] = 1
	rtClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	rt, err := refreshToken.SignedString([]byte("secret"))
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"access_token":  t,
		"refresh_token": rt,
	}, nil
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

//User Login

func (repository *RegisterRepository) Login(c *gin.Context) {

	type tokenReqBody struct {
		RefreshToken string `json:"refresh_token"`
		AccessToken  string `json:"access_token"`
	}

	tokenReq := tokenReqBody{}
	c.Bind(&tokenReq)

	input := UserRegister{}
	token := Token{}
	var email string

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")

	}

	err := repository.Db.Get(&email, `SELECT email FROM Users WHERE email= ?`, input.Email)
	fmt.Println("Email:", email)
	if err == nil {
		panic(err)
	}
	//compare the user from the request, with the one we defined:
	if email != input.Email {
		c.JSON(http.StatusUnauthorized, "Please provide valid login details")
		return
	}
	err = repository.Db.Select(&token, `SELECT email,user_type FROM Users WHERE email= ?`, input.Email)
	if err != nil {
		panic(err)
	}

	fmt.Println("Detail of token", token.Email, token.Role)

	validToken, err := GenerateJWT(token.Email, token.Role)

	fmt.Println("PAir ofToken:", validToken["AccessToken"])

	c.JSON(http.StatusOK, validToken)

	if err != nil {
		fmt.Println("Fail to generate token")
		return
	}

	var tokenDetail Token
	tokenDetail.Email = token.Email
	tokenDetail.Role = token.Role
	tokenDetail.TokenString = validToken["access_token"]
	c.JSON(http.StatusOK, token)
	fmt.Println("Token:", token)

}
