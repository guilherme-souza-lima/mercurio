package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"ssMercurio/entities"
)

type TokenJwt struct {
	AccessSecret string
}

func NewTokenJwt(AccessSecret string) TokenJwt {
	return TokenJwt{AccessSecret}
}

func (t TokenJwt) Create(id, email, cellphone, idPoints string) (string, error) {
	var mySigningKey = []byte(t.AccessSecret)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["id"] = id
	claims["email"] = email
	claims["cellphone"] = cellphone
	claims["idPoints"] = idPoints
	//claims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (t TokenJwt) Validation(tokenString string) (entities.Token, error) {
	var client entities.Token
	mySigningKey := []byte(t.AccessSecret)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there was an error in parsing")
		}
		return mySigningKey, nil
	})
	if err != nil {
		return client, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		claimsToken := entities.Token{
			ID:        claims["id"].(string),
			Email:     claims["email"].(string),
			Cellphone: claims["cellphone"].(string),
			IDPoints:  claims["idPoints"].(string),
		}
		return claimsToken, err
	}
	return client, err
}
