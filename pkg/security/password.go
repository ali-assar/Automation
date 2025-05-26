// pkg/security/password.go
package security

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func ComparePasswords(hashed, plain string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
	return err == nil, err
}
