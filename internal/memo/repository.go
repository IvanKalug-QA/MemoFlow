package memo

import (
	"memoflow/pkg/db"
)

type MemoRepository struct {
	Database *db.Db
}

func NewMemoRepository(database *db.Db) *MemoRepository {
	return &MemoRepository{
		Database: database,
	}
}

func (repo *MemoRepository) Create(memo *Memo) (*Memo, error) {
	result := repo.Database.DB.Create(memo)
	if result.Error != nil {
		return nil, result.Error
	}
	return memo, nil
}

func (repo *MemoRepository) GetByID(id int) (*Memo, error) {
	var memo Memo
	result := repo.Database.DB.First(&memo, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &memo, nil
}
