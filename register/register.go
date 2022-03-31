package register

import (
	"fmt"
	"net/http"
	"os"
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

//jwt service
type JWTService interface {
	GenerateToken(email string, isUser bool) string
	ValidateToken(token string) (*jwt.Token, error)
}

type authCustomClaims struct {
	Name string `json:"name"`
	User bool   `json:"user"`
	jwt.StandardClaims
}

type jwtServices struct {
	secretKey string
	issure    string
}

//auth-jwt
func JWTAuthService() JWTService {
	return &jwtServices{
		secretKey: getSecretKey(),
		issure:    "Priya",
	}
}

func getSecretKey() string {
	secret := os.Getenv("SECRET")
	if secret == "" {
		secret = "secret"
	}
	return secret
}

func (register *jwtServices) GenerateToken(email string, isUser bool) string {
	claims := &authCustomClaims{
		email,
		isUser,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			Issuer:    register.issure,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//encoded string
	t, err := token.SignedString([]byte(register.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (register *jwtServices) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("Invalid token", token.Header["alg"])

		}
		return []byte(register.secretKey), nil
	})

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
	auth := JWTAuthService().GenerateToken(input.Email, true)
	_, err = repository.Db.Exec(`INSERT INTO Users(name,email,user_type,auth_token) VALUES (?,?,?,?)`, input.Name, input.Email, "General", auth)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Register Successfully"})

}

// ADMIN REGISTER
func (repository *RegisterRepository) RegisterAdmin(c *gin.Context) {

	input := UserRegister{}

	auth := JWTAuthService().GenerateToken(input.Email, true)

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

// LOGIN AUTHENTICATION
type LoginService interface {
	LoginUser(email string) bool
}
type loginInformation struct {
	email string
}

func (info *loginInformation) LoginUser(email string) bool {
	return info.email == email
}

func StaticLoginService() LoginService {
	return &loginInformation{
		email: "priyanka@gmail.com",
	}
}

//login contorller interface
type LoginController interface {
	Login(ctx *gin.Context) string
}

type loginController struct {
	loginService LoginService
	jWtService   JWTService
}

func LoginHandler(loginService LoginService,
	jWtService JWTService) LoginController {
	return &loginController{
		loginService: loginService,
		jWtService:   jWtService,
	}
}

func (register *loginController) Login(ctx *gin.Context) string {
	var credential UserRegister
	err := ctx.ShouldBind(&credential)
	if err != nil {
		return "no data found"
	}
	isUserAuthenticated := register.loginService.LoginUser(credential.Email)
	if isUserAuthenticated {
		return register.jWtService.GenerateToken(credential.Email, true)
	}
	return ""

}
