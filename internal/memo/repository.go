package memo

import (
	"memoflow/pkg/db"
)

type MemoResository struct {
	Database *db.Db
}

func NewMemoRepository(database *db.Db) *MemoResository {
	return &MemoResository{
		Database: database,
	}
}

func (repo *MemoResository) Create(link *Memo) {

}
