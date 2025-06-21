package click

import (
	"github.com/google/uuid"
	"link-shorter/pkg/base"
)

type Model struct {
	base.Model
	LinkId *uuid.UUID `gorm:"type:uuid;index"`
}

func (*Model) TableName() string {
	return "clicks"
}

func NewClick(linkId string) *Model {
	parsedUuid, err := uuid.Parse(linkId)
	if err != nil {
		return nil
	}

	return &Model{
		LinkId: &parsedUuid,
	}
}
