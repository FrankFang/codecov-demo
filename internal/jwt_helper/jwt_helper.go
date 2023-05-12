package jwt_helper

import (
	"crypto/rand"
	"io"
	"io/ioutil"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

func GenerateJWT(user_id int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user_id,
	})
	key, err := getHMACKey()
	if err != nil {
		return "", err
	}
	return token.SignedString(key)
}
func getHMACKey() ([]byte, error) {
	keyPath := viper.GetString("jwt.hmac.key_path")
	return ioutil.ReadFile(keyPath)
}
func GenerateHMACKey() ([]byte, error) {
	key := make([]byte, 64)
	_, err := io.ReadFull(rand.Reader, key)
	if err != nil {
		return nil, err
	}
	return key, nil
}

func ParseJWT(jwtString string) (jwt.MapClaims, error) {
	key, err := getHMACKey()
	if err != nil {
		return nil, err
	}
	token, err := jwt.Parse(jwtString, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return nil, err
	}
	tokenClaims := token.Claims.(jwt.MapClaims)
	return tokenClaims, nil
}
