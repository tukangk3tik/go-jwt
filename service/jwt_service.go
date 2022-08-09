package service

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"go-jwt/dto"
	"io/ioutil"
	"log"
	"time"
)

var JwtServiceInstance = CreateJwtServiceInstance()

func CreateJwtServiceInstance() Jwt {
	prvKeyFileName := "files/cert/jwt-private.key"
	pubKeyFileName := "files/cert/jwt-public.key"

	prvKey, err := ioutil.ReadFile(prvKeyFileName)
	if err != nil {
		log.Fatalln(err)
	}
	pubKey, err := ioutil.ReadFile(pubKeyFileName)
	if err != nil {
		log.Fatalln(err)
	}
	return NewJwtService(prvKey, pubKey)
}

type JwtService interface {
	GenerateToken(email string) dto.TokenDto
	ValidateToken(token string) (any, error)
}

type Jwt struct {
	privateKey []byte
	publicKey  []byte
}

func NewJwtService(privateKey []byte, publicKey []byte) Jwt {
	return Jwt{
		privateKey: privateKey,
		publicKey:  publicKey,
	}
}

func (j Jwt) GenerateToken(email string) dto.TokenDto {
	key, err := jwt.ParseRSAPrivateKeyFromPEM(j.privateKey)
	if err != nil {
		panic(err)
	}

	tokenType := "Bearer"
	tokenExpiresAt := time.Now().AddDate(0, 0, 7)         //token valid for 7 days
	refreshTokenExpiresAt := time.Now().AddDate(0, 0, 30) //token valid for 7 days
	now := time.Now().UTC()

	// claims for access token
	tClaims := make(jwt.MapClaims)
	tClaims["email"] = email
	tClaims["iat"] = now.Unix()
	tClaims["nbf"] = now.Unix()
	tClaims["exp"] = tokenExpiresAt.Unix()

	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodRS256, tClaims).SignedString(key)
	if err != nil {
		panic(err)
	}

	// claims for refresh token
	rtClaims := make(jwt.MapClaims)
	rtClaims["email"] = email
	rtClaims["exp"] = refreshTokenExpiresAt.Unix()

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodRS256, rtClaims).SignedString(key)
	if err != nil {
		panic(err)
	}

	tokenResult := dto.TokenDto{
		TokenType:    tokenType,
		ExpiresIn:    tokenExpiresAt.Unix(),
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	return tokenResult
}

func (j Jwt) ValidateToken(token string) (any, error) {
	key, err := jwt.ParseRSAPublicKeyFromPEM(j.publicKey)
	if err != nil {
		panic(err)
	}

	t, err := jwt.Parse(token, func(jwtToken *jwt.Token) (any, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected method: %s", jwtToken.Header["alg"])
		}
		return key, nil
	})

	if err != nil {
		return nil, fmt.Errorf("validate: %w", err)
	}

	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok || !t.Valid {
		return nil, fmt.Errorf("validate: invalid")
	}
	return claims["email"], nil
}
