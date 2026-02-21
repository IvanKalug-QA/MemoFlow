package stat

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Stat struct {
	gorm.Model
	MemoId uint           `json:"memo_id"`
	Clicks int            `json:"clicks"`
	Date   datatypes.Date `json:"date"`
}
