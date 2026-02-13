package memo

import (
	"memoflow/pkg/db"

	"gorm.io/gorm/clause"
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

func (repo *MemoRepository) Update(memo *Memo) (*Memo, error) {
	result := repo.Database.DB.Clauses(clause.Returning{}).Updates(memo)
	if result.Error != nil {
		return nil, result.Error
	}
	return memo, nil
}

func (repo *MemoRepository) Delete(id int) error {
	result := repo.Database.DB.Delete(&Memo{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
