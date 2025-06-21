package click

import "gorm.io/datatypes"

type GetCountParams struct {
	LinkId string
}

type GetStatsParams struct {
	LinkId    string          `json:"linkId"`
	StartDate *datatypes.Date `schema:"startDate" validate:"required,datetime=2006-01-02"`
	EndDate   *datatypes.Date `schema:"endDate" validate:"required,datetime=2006-01-02"`
	Timezone  *string         `json:"timezone"`
}

type AddPayload struct {
	LinkId string
}
