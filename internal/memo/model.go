package memo

import "gorm.io/gorm"

type Memo struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
}

func NewMemo(body *MemoRequest) *Memo {
	return &Memo{
		Name:        body.Name,
		Description: body.Description,
	}
}
