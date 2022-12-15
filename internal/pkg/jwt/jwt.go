package jwt

import (
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	SecretKey = []byte("secret")
)

func GenerateToken(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Minute * 1).Unix()
	tokenString, err := token.SignedString(SecretKey)
	print(tokenString)
	print(err)
	if err != nil {
		log.Fatal("Error in generating key")
		return "", err
	}
	println()
	println("--------------------------------")
	print(ParseToken(tokenString))
	println("--------------------------------")
	return tokenString, nil
}
func ParseToken(tokenStr string) (string, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) { return SecretKey, nil })
	println(token)
	println(token.Header)
	println(token.Claims)
	println(token.Claims.(jwt.MapClaims)["username"].(string))
	println(token.Claims.(jwt.MapClaims)["exp"].(float64))
	println(token.Signature)
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username := claims["username"].(string)
		return username, nil
	} else {
		return "", err
	}
}
