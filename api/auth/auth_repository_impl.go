package auth

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

type authRepository struct {
	jwt          *jwt.Token
	parse        *jwt.Parser
	signingKey   []byte
	expiredToken int
}

func (a authRepository) GenerateJWT(ip string) (string, error) {
	claims := a.jwt.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["ip"] = ip
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(a.expiredToken)).Unix()
	return a.jwt.SignedString(a.signingKey)
}

func (a authRepository) ClaimToken(token, ip string) error {

	tempToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there was an error in parsing")
		}
		return a.signingKey, nil
	})

	if err != nil {
		return err
	}

	claim, ok := tempToken.Claims.(jwt.MapClaims)
	if ok && tempToken.Valid {
		if claim["ip"] == ip {
			return nil
		}
		return fmt.Errorf("not Authorize")
	}
	return fmt.Errorf("not Authorize")
}

func NewAuthRepository(jwt *jwt.Token, signingKey []byte, expiredToken int) IAuthRepository {
	return &authRepository{jwt: jwt, signingKey: signingKey, expiredToken: expiredToken}
}
