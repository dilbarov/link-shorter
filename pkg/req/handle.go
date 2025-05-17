package req

import (
	"io"
)

func HandleBody[T any](body io.ReadCloser) (*T, error) {
	payload, err := Decode[T](body)

	if err != nil {
		return nil, err
	}

	err = IsValid[T](payload)

	if err != nil {
		return nil, err
	}

	return payload, nil
}
