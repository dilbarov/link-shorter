package db

import (
	"gorm.io/gorm"
	"sync"
)

func GetCount(qb *gorm.DB, total *int64, wg *sync.WaitGroup, errCh chan error) {
	defer wg.Done()

	if err := qb.Count(total).Error; err != nil {
		errCh <- err
	}
}
