package db

import (
	"gorm.io/gorm"
	"sync"
)

func FindAndCount[T any](db *gorm.DB, filter CommonFilter) ([]*T, int64, error) {
	var (
		errCh = make(chan error, 2)
		wg    sync.WaitGroup
		total int64
		items []*T
	)

	countQuery := db.Session(&gorm.Session{})
	findQuery := db.Session(&gorm.Session{})

	wg.Add(1)
	go GetCount(countQuery, &total, &wg, errCh)

	wg.Add(1)
	go FindItems(findQuery, &items, filter, &wg, errCh)

	wg.Wait()
	close(errCh)

	if err := <-errCh; err != nil {
		return nil, 0, err
	}

	return items, total, nil
}
