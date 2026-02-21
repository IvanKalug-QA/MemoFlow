package memo

import (
	"memoflow/internal/stat"

	"gorm.io/gorm"
)

type Memo struct {
	gorm.Model
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Stats       []stat.Stat `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func NewMemo(body *MemoRequest) *Memo {
	return &Memo{
		Name:        body.Name,
		Description: body.Description,
	}
}
