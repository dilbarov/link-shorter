package db

import (
	"gorm.io/gorm"
	"sync"
)

type CommonFilter struct {
	Limit  int
	Offset int
}

func FindItems[T any](qb *gorm.DB, items *T, filter CommonFilter, wg *sync.WaitGroup, errCh chan error) {
	defer wg.Done()

	if filter.Limit > 0 {
		qb = qb.Limit(filter.Limit)
	}
	if filter.Offset > 0 {
		qb = qb.Offset(filter.Offset)
	}
	qb = qb.Order("created_at ASC")

	if err := qb.Find(items).Error; err != nil {
		errCh <- err
	}

}
