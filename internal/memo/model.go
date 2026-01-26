package memo

import "gorm.io/gorm"

type Memo struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
}
