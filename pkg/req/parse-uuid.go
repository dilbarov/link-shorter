package req

import (
	"github.com/google/uuid"
	"net/http"
)

func ParseUUID(req *http.Request, field string) (*string, error) {
	id := req.PathValue(field)

	if _, err := uuid.Parse(id); err != nil {
		return nil, err
	}

	return &id, nil
}
