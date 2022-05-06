package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrExpToken     = errors.New("token expired")
	ErrInvalidToken = errors.New("invalid token")
)

type Payload struct {
	ID       uuid.UUID `json:"id"`
	Email    string    `json:"email"`
	IssuedAt time.Time `json:"iat"`
	Exp      time.Time `json:"exp"`
}

func NewPayload(useremail string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:       tokenID,
		Email:    useremail,
		IssuedAt: time.Now(),
		Exp:      time.Now().Add(duration),
	}
	return payload, nil
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.Exp) {
		return ErrExpToken
	}
	return nil
}
