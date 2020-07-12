package providers

import "golang.org/x/crypto/bcrypt"

type HashProvider interface {
	GenerateHash(password string) string
	Verify(hash string, password string) bool
}

type bcryptHashProvider struct{}

func NewBCryptHashProvider() HashProvider {
	return &bcryptHashProvider{}
}

func (provider *bcryptHashProvider) GenerateHash(password string) string {
	bytesPassword := []byte(password)

	hash, err := bcrypt.GenerateFromPassword(bytesPassword, bcrypt.DefaultCost)

	if err != nil {
		return ""
	}

	return string(hash)
}

func (provider *bcryptHashProvider) Verify(hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	if err != nil {
		return false
	}

	return true
}
