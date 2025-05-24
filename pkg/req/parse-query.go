package req

import (
	"github.com/gorilla/schema"
	"net/http"
)

func ParseQuery[T any](req *http.Request, query *T) error {
	decoder := schema.NewDecoder()
	return decoder.Decode(query, req.URL.Query())
}
