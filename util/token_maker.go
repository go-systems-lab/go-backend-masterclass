package util

import (
	"fmt"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/google/uuid"
	"github.com/o1egl/paseto"
)

type TokenMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

func NewTokenMaker(symmetricKey string) (*TokenMaker, error) {
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size: must be %d characters", chacha20poly1305.KeySize)
	}

	paseto := paseto.NewV2()
	tokenMaker := &TokenMaker{
		paseto:       paseto,
		symmetricKey: []byte(symmetricKey),
	}

	return tokenMaker, nil
}

type Token struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func (tokenMaker *TokenMaker) CreateToken(username string, duration time.Duration) (string, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	token := Token{
		ID:        tokenID,
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}

	return tokenMaker.paseto.Encrypt(tokenMaker.symmetricKey, token, nil)

}

var (
	ErrInvalidToken = fmt.Errorf("token is invalid")
	ErrExpiredToken = fmt.Errorf("token has expired")
)

func (tokenMaker *TokenMaker) VerifyToken(tokenString string) (*Token, error) {
	token := &Token{}

	err := tokenMaker.paseto.Decrypt(tokenString, tokenMaker.symmetricKey, token, nil)
	if err != nil {
		return nil, ErrInvalidToken
	}

	if time.Now().After(token.ExpiredAt) {
		return nil, ErrExpiredToken
	}

	return token, nil
}
