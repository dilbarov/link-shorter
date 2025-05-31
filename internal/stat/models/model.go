package stat

import (
	"github.com/google/uuid"
	"link-shorter/pkg/base"
)

type Model struct {
	base.Model
	LinkId *uuid.UUID `gorm:"type:uuid;index"`
}

func (*Model) TableName() string {
	return "stats"
}
