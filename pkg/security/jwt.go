// pkg/security/jwt.go
package security

import (
	"backend/internal/config"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

var cfg = config.Load()

func GenerateStaticToken(adminID uuid.UUID, roleID int) (string, error) {
	claims := jwt.MapClaims{
		"admin_id": adminID.String(),
		"role":     roleID,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(cfg.JWTSecret))
}

func GenerateDynamicToken(adminID uuid.UUID, roleID int) (string, error) {
	claims := jwt.MapClaims{
		"admin_id": adminID.String(),
		"role":     roleID,
		"exp":      time.Now().Add(15 * time.Minute).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(cfg.JWTSecret))
}

func ValidateToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.JWTSecret), nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, jwt.ErrInvalidKey
	}
	return claims, nil
}
