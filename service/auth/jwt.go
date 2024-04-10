package auth

import (
	"github.com/aryala7/ecom/config"
	"github.com/golang-jwt/jwt/v5"
	"strconv"
	"time"
)

func CreateJwt(secret []byte, userId int) (string, error) {

	expiration := time.Second * time.Duration(config.Envs.JWTExpirationInSeconds)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":    strconv.Itoa(userId),
		"expiredAt": time.Now().Add(expiration),
	})
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
