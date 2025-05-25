package security

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var (
	staticSecretKey  []byte
	dynamicSecretKey []byte
	lastNonce        int64
)

const (
	staticTokenExpiryDuration  = 24 * time.Hour
	dynamicTokenExpiryDuration = 24 * time.Hour
)

type CustomClaims struct {
	AdminID uuid.UUID `json:"admin_id"`
	Role    int       `json:"role"`
	Nonce   string    `json:"nonce,omitempty"`
	jwt.RegisteredClaims
}

func JWTInit(staticSecret, dynamicSecret string) {
	staticSecretKey = []byte(staticSecret)
	dynamicSecretKey = []byte(dynamicSecret)
}

func GenerateStaticToken(id uuid.UUID, role int) (string, error) {
	claims := CustomClaims{
		AdminID: id,
		Role:    role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(staticTokenExpiryDuration)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(staticSecretKey)
}

func GenerateDynamicToken(id uuid.UUID, role int) (string, error) {
	nonce := strconv.FormatInt(time.Now().UnixNano(), 10)
	claims := CustomClaims{
		AdminID: id,
		Role:    role,
		Nonce:   nonce,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(dynamicTokenExpiryDuration)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(dynamicSecretKey)
}

func VerifyStaticToken(tokenString string, expectedID uuid.UUID, expectedRole int) error {
	return verifyToken(tokenString, staticSecretKey, expectedID.String(), expectedRole, false)
}

func VerifyDynamicToken(tokenString string, expectedID uuid.UUID, expectedRole int) error {
	return verifyToken(tokenString, dynamicSecretKey, expectedID.String(), expectedRole, true)
}

func verifyToken(tokenString string, secret []byte, expectedID string, expectedRole int, checkNonce bool) error {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return secret, nil
	})
	if err != nil {
		return errors.New("invalid token")
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok || !token.Valid {
		return errors.New("invalid claims or token not valid")
	}

	if claims.ExpiresAt.Time.Before(time.Now()) {
		return errors.New("token expired")
	}

	if claims.AdminID.String() != expectedID {
		return errors.New("username mismatch")
	}

	if claims.Role != expectedRole {
		return errors.New("role mismatch")
	}

	if checkNonce {
		nonce, err := strconv.ParseInt(claims.Nonce, 10, 64)
		if err != nil {
			return errors.New("invalid nonce format")
		}
		if nonce <= lastNonce {
			return errors.New("replay attack detected")
		}
		lastNonce = nonce
	}

	return nil
}
