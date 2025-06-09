package db

type ListResult[R any] struct {
	Items []R
	Count int
}

func NewListResult[R any](items []R, count int) *ListResult[R] {
	return &ListResult[R]{
		Items: items,
		Count: count,
	}
}
