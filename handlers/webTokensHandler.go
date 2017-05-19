package handlers

import (
	"errors"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

/* claims - JWT claims object*/
type claims struct {
	UserId int `json:"user_id"`
	jwt.StandardClaims
}

// generateJWT - Generates the JWT token with custom expiry time and user id
func generateJWT(user int) (Jwt string, err error) {
	mySigningKey := []byte("secret*#key#*for*#AES&encryption")
	claims := claims{
		user,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + 7200,
			Issuer:    "CRUD",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}
	return signedString, nil
}

// validateJWT - validate the JWT token and parse the user_id from the token
func validateJWT(webToken string, c *gin.Context) (valid bool) {
	parsedToken, _ := jwt.ParseWithClaims(webToken, &claims{}, func(parsedToken *jwt.Token) (interface{}, error) {
		return []byte("secret*#key#*for*#AES&encryption"), nil
	})
	if claims, ok := parsedToken.Claims.(*claims); ok && parsedToken.Valid {
		c.Set("user_id", claims.UserId)
		return true
	}
	return false
}

// TokenValidator - gin middleware function to validate the jwt tokens in the incoming requests
func TokenValidator() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !validateJWT(c.Request.Header.Get("Authorization"), c) {
			c.Error(errors.New("No auth token sent"))
			c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized,
				"message": "Auth token is invalid"})
			c.Abort()
			return
		}
		c.Set("token", c.Request.Header.Get("Authorization"))
		c.Next()
	}
}

/* TODO check for the auth token the json request also */
