package token

import (
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gitnyasha/go-hekani-backend/util"
	"github.com/stretchr/testify/require"
)

func TestJWTMAker(t *testing.T) {
	maker, err := NewJWTMaker(util.RandomLongString(32))
	require.NoError(t, err)

	useremail := util.RandomEmail()
	duration := time.Minute

	issueAt := time.Now()
	expiredAt := issueAt.Add(duration)

	token, err := maker.CreateToken(useremail, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.NotZero(t, payload.ID)
	require.Equal(t, useremail, payload.Email)
	require.WithinDuration(t, issueAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.Exp, time.Second)
}

// check if the token is expired
func TestExpToken(t *testing.T) {
	maker, err := NewJWTMaker(util.RandomLongString(32))
	require.NoError(t, err)

	token, err := maker.CreateToken(util.RandomEmail(), -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.Error(t, err)
	require.Equal(t, ErrExpToken, err)
	require.Nil(t, payload)
}

// check if the token is valid
func TestInvalidToken(t *testing.T) {
	payload, err := NewPayload(util.RandomEmail(), time.Minute)
	require.NoError(t, err)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodNone, payload)
	token, err := jwtToken.SignedString(jwt.UnsafeAllowNoneSignatureType)
	require.NoError(t, err)

	maker, err := NewJWTMaker(util.RandomLongString(32))
	require.NoError(t, err)

	payload, err = maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrInvalidToken.Error())
	require.Nil(t, payload)

}
