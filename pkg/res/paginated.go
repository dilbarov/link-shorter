package res

type Paginated[T any] struct {
	Items []T `json:"items"`
	Count int `json:"count"`
}
