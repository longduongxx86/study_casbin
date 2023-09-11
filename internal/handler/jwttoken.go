package handler

import (
	"docker_gozero/internal/types"
	"errors"

	jwt "github.com/golang-jwt/jwt/v5"

)

func CreateJWT(user *types.UserIden) (string, error) {
	mySigningKey := []byte("authdemo")

	claims := jwt.MapClaims{
		"role_id": user.RoleID,
		"id":      user.Id,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(mySigningKey)
}

func ParseJWT(tokenString string) (interface{}, error) {
	token, err := jwt.Parse(tokenString, func(*jwt.Token) (interface{}, error) {
		return []byte("authdemo"), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		var user types.UserIden
		user.RoleID = claims["role_id"].(string)
		user.Id = claims["id"].(int)

		return user, nil
	} else {
		return nil, errors.New("error when claims token")
	}
}
