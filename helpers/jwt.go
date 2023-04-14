package helpers

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"log"
)

const key = "P4NC4$!L4"

func GenerateToken(id uint, username string) (res string, err error) {
	claims := jwt.MapClaims{
		"id":       id,
		"username": username,
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := parseToken.SignedString([]byte(key))
	if err != nil {
		log.Fatal("Error parse")
		return
	}

	res = signedToken

	return
}

func VerifyToken(token string) (res interface{}, err error) {
	parseToken, err := jwt.Parse(token, func(t *jwt.Token) (res interface{}, err error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			err = errors.New("Invalid method header alg")
			return
		}
		res, err = []byte(key), nil
		return
	})

	if _, ok := parseToken.Claims.(jwt.MapClaims); !ok && !parseToken.Valid {
		err = errors.New("token invalid")
		return
	}

	res = parseToken.Claims.(jwt.MapClaims)
	return
}
