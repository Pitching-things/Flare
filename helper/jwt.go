package helper

import (
	"errors"
	"time"

	"github.com/Pitching-things/Flare/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

func JwtCreate(Data interface{}) (string, error) {
	claims := models.Claims{
		Id: Data,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(viper.GetString("SECRET_KEY")))

	return tokenString, err
}

func DataOfJwt(tokenString string) (models.Claims, error) {
	claims := models.Claims{}
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(viper.GetString("SECRET_KEY")), nil
	})

	if err != nil {
		return models.Claims{}, err
	}

	if !token.Valid {
		return models.Claims{}, errors.New("token isn't valid")
	}

	return claims, nil
}
