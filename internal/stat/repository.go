package stat

import (
	"memoflow/pkg/db"
	"time"

	"gorm.io/datatypes"
)

func NewStatRepository(db *db.Db) *StatRepository {
	return &StatRepository{Db: db}
}

type StatRepository struct {
	*db.Db
}

func (s *StatRepository) AddClick(memoId uint) {
	var stat Stat
	currentDate := datatypes.Date(time.Now())
	s.Db.Find(&stat, "memo_id = ? and date = ?", memoId, currentDate)
	if stat.ID == 0 {
		s.DB.Create(&Stat{
			MemoId: memoId,
			Clicks: 1,
			Date:   currentDate,
		})
	} else {
		stat.Clicks++
		s.DB.Save(&stat)
	}
}
