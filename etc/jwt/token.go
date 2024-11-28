package jwt

import (
	"backend/etc/Utime"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"os"
	"time"
)

var JwtSecret = []byte(os.Getenv("SECRET_KEY"))

type Claims struct {
	UserID      string `json:"user_id"`
	Role        string `json:"role"`
	AccessLevel int    `json:"access_level"`
	jwt.RegisteredClaims
}

func GenerateToken(userID string, username string, accessLevel int) (string, error) {
	expirationTime := Utime.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID:      userID,
		AccessLevel: accessLevel,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtSecret, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
