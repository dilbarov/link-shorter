package click

import "gorm.io/datatypes"

type CountStatItem struct {
	Count int64          `json:"count"`
	Date  datatypes.Date `json:"date"`
}

func NewCountStatItem(count int64, date datatypes.Date) *CountStatItem {
	return &CountStatItem{Count: count, Date: date}
}
