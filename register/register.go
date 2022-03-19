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
	Password string `json:"password"`
}

type RegisterRepository struct {
	Db *sqlx.DB
}
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// Create the JWT key used to create the signature
var jwtKey = []byte("my_secret_key")

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
	var email, expectedPassword string

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}

	err := repository.Db.Get(&email, `SELECT email FROM Users WHERE email= ?`, input.Email)
	res := repository.Db.Get(&expectedPassword, `SELECT email FROM Users WHERE password= ?`, input.Password)

	c.JSON(http.StatusOK, email)
	c.JSON(http.StatusOK, res)
	//compare the user from the request, with the one we defined:

	if input.Email != email {
		c.JSON(http.StatusUnauthorized, "Please provide valid login details")
		return
	} else if err != nil {
		c.JSON(http.StatusUnauthorized, "Please register to login")

	}

	c.JSON(http.StatusOK, gin.H{"message": "LOgin Successfully"})

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: input.Email,
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

}

// func (repository *RegisterRepository) Welcome(c *gin.Context) {
// 	{

// 		// Get the JWT string from the cookie
// 		tknStr := c.Value

// 		// Initialize a new instance of `Claims`
// 		claims := &Claims{}

// 		// Parse the JWT string and store the result in `claims`.
// 		// Note that we are passing the key in this method as well. This method will return an error
// 		// if the token is invalid (if it has expired according to the expiry time we set on sign in),
// 		// or if the signature does not match
// 		tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
// 			return jwtKey, nil
// 		})
// 		if err != nil {
// 			if err == jwt.ErrSignatureInvalid {

// 				c.JSON(http.StatusUnauthorized, gin.H{"Invalid Signature": err.Error()})
// 				return
// 			}
// 			c.JSON(http.StatusBadRequest, gin.H{"BAd Request": err.Error()})
// 			return
// 		}
// 		if !tkn.Valid {
// 			c.JSON(http.StatusUnauthorized, gin.H{"Invalid Signature": err.Error()})
// 			return
// 		}

// 		// Finally, return the welcome message to the user, along with their
// 		// username given in the token
// 		c.JSON(([]byte(fmt.Sprintf("Welcome %s!", claims.Username))))
// 	}
// }
