package jwtProvider

import (
	"fastFood/common"
	"fastFood/components/tokenProvider"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type jwtProvider struct {
	prefix string
	secret string
}

func NewJwtProvider(prefix, secret string) *jwtProvider {
	return &jwtProvider{prefix: prefix, secret: secret}
}

func (j *jwtProvider) SecretKey() string {
	return j.secret
}

type myClaims struct {
	Payload common.TokenPayLoad `json:"payload"`
	jwt.RegisteredClaims
}

type token struct {
	Token   string    `json:"token"`
	Created time.Time `json:"created"`
	Expiry  int       `json:"expiry"`
}

func (t *token) GetToken() string {
	return t.Token
}

func (j *jwtProvider) Generate(data tokenProvider.TokenPayload, expiry int) (tokenProvider.Token, error) {
	now := time.Now()

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims{
		common.TokenPayLoad{UId: data.UserId()},
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Local().Add(time.Second * time.Duration(expiry))),
			IssuedAt:  jwt.NewNumericDate(now.Local()),
			ID:        fmt.Sprintf("%d", now.UnixNano()),
		},
	})

	myToken, err := t.SignedString([]byte(j.secret))

	if err != nil {
		return nil, err
	}

	return &token{
		Token:   myToken,
		Expiry:  expiry,
		Created: now,
	}, nil
}

func (j *jwtProvider) Validate(myToken string) (tokenProvider.TokenPayload, error) {
	res, err := jwt.ParseWithClaims(myToken, &myClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, tokenProvider.ErrUnexpectedSigningMethod(fmt.Sprintf("unexpected signing method: %v", t.Header["alg"]))
		}
		return []byte(j.secret), nil
	})

	if err != nil {
		return nil, tokenProvider.ErrNotFound()
	}

	if !res.Valid {
		return nil, tokenProvider.ErrNotFound()
	}

	claims, ok := res.Claims.(*myClaims)

	if !ok {
		return nil, tokenProvider.ErrNotFound()
	}

	return claims.Payload, nil
}
