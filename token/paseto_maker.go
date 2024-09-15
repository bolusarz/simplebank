package token

import (
	"aidanwoods.dev/go-paseto"
	"fmt"
	"golang.org/x/crypto/chacha20poly1305"
	"time"
)

type PasetoMaker struct {
	paseto       paseto.V4SymmetricKey
	symmetricKey []byte
}

func NewPasetoMaker(symmetricKey string) (Maker, error) {
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size: must be exactly %d characters", chacha20poly1305.KeySize)
	}

	maker := &PasetoMaker{
		paseto:       paseto.NewV4SymmetricKey(),
		symmetricKey: []byte(symmetricKey),
	}
	return maker, nil
}

func (maker PasetoMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", err
	}
	token := paseto.NewToken()
	token.SetExpiration(payload.ExpiredAt)
	token.SetIssuedAt(payload.IssuedAt)
	token.SetSubject(payload.Username)
	token.SetNotBefore(time.Now())
	_ = token.Set("payload", payload)

	return token.V4Encrypt(maker.paseto, maker.symmetricKey), nil
}

func (maker PasetoMaker) VerifyToken(token string) (*Payload, error) {
	parser := paseto.NewParser()
	parsedToken, err := parser.ParseV4Local(maker.paseto, token, maker.symmetricKey)

	if err != nil {
		return nil, err
	}

	payload := &Payload{}

	err = parsedToken.Get("payload", payload)

	if err != nil {
		return nil, ErrInvalidToken
	}

	return payload, nil

}
