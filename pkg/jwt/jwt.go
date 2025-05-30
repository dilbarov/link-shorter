package jwt

import "github.com/golang-jwt/jwt/v5"

type Data struct {
	Sub   string
	Email string
}

type Service struct {
	Secret string
}

func NewJWTService(secret string) *Service {
	return &Service{
		Secret: secret,
	}
}

func (j *Service) Create(data Data) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   data.Sub,
		"email": data.Email,
	})

	s, err := token.SignedString([]byte(j.Secret))

	if err != nil {
		return "", err
	}

	return s, nil
}

func (j *Service) Parse(token string) (bool, *Data) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.Secret), nil
	})

	if err != nil {
		return false, nil
	}

	id := t.Claims.(jwt.MapClaims)["sub"].(string)
	email := t.Claims.(jwt.MapClaims)["email"].(string)

	return true, &Data{
		Sub:   id,
		Email: email,
	}
}
