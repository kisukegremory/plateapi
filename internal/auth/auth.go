package auth

import (
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var apiSecret = []byte(os.Getenv("APICAR_SECRET"))

type UserClaim struct {
	jwt.RegisteredClaims
	Exp float64
	Sub string
	Iss string
}

func GenerateJwt() (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		UserClaim{
			RegisteredClaims: jwt.RegisteredClaims{},
			Iss:              "plate-server",
			Sub:              "client-user",
			Exp:              float64(time.Now().Add(time.Hour).Unix()),
		})
	return token.SignedString(apiSecret)
}

func ValidationMiddleware(c *gin.Context) {
	tokenString, errCookie := c.Cookie("Authorization")
	if errCookie != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	} // TODO: add a function to read the header or the cookie

	var userClaim UserClaim
	token, _ := jwt.ParseWithClaims(tokenString, &userClaim, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return apiSecret, nil
	}) // TODO: add a function to parse the JWT, and return token, userClaim, err

	if !token.Valid || float64(time.Now().Unix()) > userClaim.Exp {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	c.Next()

}
