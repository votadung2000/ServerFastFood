package tokenProvider

type TokenPayload interface {
	UserId() int
}

type Token interface {
	GetToken() string
}

type Provider interface {
	Generate(data TokenPayload, expiry int) (Token, error)
	Validate(token string) (TokenPayload, error)
	SecretKey() string
}
