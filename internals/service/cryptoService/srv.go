package cryptoService

import "golang.org/x/crypto/bcrypt"

type cryptoSrv struct{}

type CryptoService interface {
	HashPassword(password string) (string, error)
	ComparePassword(password string, hash string) bool
}

func NewCryptoService() CryptoService {
	return &cryptoSrv{}
}

func (t *cryptoSrv) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (t *cryptoSrv) ComparePassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
