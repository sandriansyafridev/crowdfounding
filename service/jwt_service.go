package service

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type MyCustomClaims struct {
	UserID uint64 `json:"user_id"`
	jwt.StandardClaims
}

type JWTService interface {
	GenerateToken(UserID uint64) (tokenString string, err error)
	ParseToken(tokenString string) (*jwt.Token, error)
}

type JWTServiceImpl struct {
	Issuer     string
	SigningKey []byte
}

func NewJWTServiceImpl() *JWTServiceImpl {
	return &JWTServiceImpl{
		Issuer:     "potatro",
		SigningKey: []byte("AllYourBase"),
	}
}

func (jwtService *JWTServiceImpl) GenerateToken(UserID uint64) (tokenString string, err error) {

	// Create the Claims
	claims := MyCustomClaims{
		UserID: UserID,
		StandardClaims: jwt.StandardClaims{
			Issuer:    jwtService.Issuer,
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().AddDate(1, 0, 0).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	if tokenString, err = token.SignedString(jwtService.SigningKey); err != nil {
		return tokenString, err
	} else {
		return tokenString, nil
	}

}

func (jwtService *JWTServiceImpl) ParseToken(tokenString string) (*jwt.Token, error) {

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return jwtService.SigningKey, nil
	})

	if err != nil {
		return token, err
	}

	return token, nil

}
