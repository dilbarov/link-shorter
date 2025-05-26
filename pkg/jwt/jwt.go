package jwt

import "github.com/golang-jwt/jwt/v5"

type Service struct {
	secret string
}

func NewJWTService(secret string) *Service {
	return &Service{
		secret: secret,
	}
}

func (j *Service) Create(id string, email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   id,
		"email": email,
	})

	s, err := token.SignedString([]byte(j.secret))

	if err != nil {
		return "", err
	}

	return s, nil
}

//func (j *Service) Validate(token string) bool {
//
//}
