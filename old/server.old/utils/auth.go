package utils

import (
	"fmt"
	"time"

	"github.com/astenmies/lychee/server/model"
	jwt "github.com/dgrijalva/jwt-go"
)

var PublicKey = []byte("secret")

func GenerateToken(user model.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * time.Duration(24)).Unix(), //This token will live for 24 hours
		"iat": time.Now().Unix(),
		"sub": user.ID,
	})

	tokenString, err := token.SignedString(PublicKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

//Helper function
func CheckToken(jwtToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return PublicKey, nil
	})

	// parts := strings.Split(jwtToken, ".")
	// log.Println(parts)
	if err != nil {
		return nil, err
	}

	if token.Valid {
		return token, nil
	}

	return token, nil
}
