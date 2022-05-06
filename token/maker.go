package token

import "time"

type Maker interface {
	CreateToken(useremail string, duration time.Duration) (string, error)

	VerifyToken(token string) (*Payload, error)
}
