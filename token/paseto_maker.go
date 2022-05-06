package token

import (
	"fmt"
	"time"

	"github.com/o1egl/paseto"
	"golang.org/x/crypto/chacha20poly1305"
)

type PasetoMaker struct {
	paseto *paseto.V2
	symKey []byte
}

func NewPasetoMaker(symKey string) (Maker, error) {
	if len(symKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("secret key must be at least %d characters", chacha20poly1305.KeySize)
	}

	maker := &PasetoMaker{
		paseto: paseto.NewV2(),
		symKey: []byte(symKey),
	}

	return maker, nil
}

func (maker *PasetoMaker) CreateToken(useremail string, duration time.Duration) (string, error) {
	payload, err := NewPayload(useremail, duration)
	if err != nil {
		return "", err
	}

	return maker.paseto.Encrypt(maker.symKey, payload, nil)
}

func (maker *PasetoMaker) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}

	err := maker.paseto.Decrypt(token, maker.symKey, payload, nil)
	if err != nil {
		return nil, ErrInvalidToken
	}

	err = payload.Valid()
	if err != nil {
		return nil, err
	}

	return payload, nil
}
