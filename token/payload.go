package token

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

var (
	ErrInvalidToken = errors.New("invalid token")
	ErrExpiredToken = errors.New("expired token")
)

func (payload Payload) GetExpirationTime() (*jwt.NumericDate, error) {
	return jwt.NewNumericDate(payload.ExpiredAt), nil
}

func (payload Payload) GetIssuedAt() (*jwt.NumericDate, error) {
	return jwt.NewNumericDate(payload.IssuedAt), nil
}

func (payload Payload) GetNotBefore() (*jwt.NumericDate, error) {
	return jwt.NewNumericDate(time.Now()), nil
}

func (payload Payload) GetIssuer() (string, error) {
	return "", nil
}

func (payload Payload) GetSubject() (string, error) {
	return payload.Username, nil
}

func (payload Payload) GetAudience() (jwt.ClaimStrings, error) {
	return nil, nil
}

func NewPayload(username string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:        tokenID,
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload, nil
}
