package core

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

var secret = os.Getenv("vocaportal_secret")

// JWT model
type JWT struct {
	Token string `json:"jwt"`
}

type JWTData struct {
	Name string
	jwt.StandardClaims
}

func GenerateJWT(name string) *JWT {
	data := JWTData{
		name,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(2 * time.Hour).Unix(),
			Issuer:    "vocaportal",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, data)

	signedToken, err := token.SignedString([]byte(secret))

	if err != nil {
		LogErr(err)
		return nil
	}

	return &JWT{Token: signedToken}
}
