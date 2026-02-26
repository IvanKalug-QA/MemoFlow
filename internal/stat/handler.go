package stat

import (
	"memoflow/configs"
	"memoflow/pkg/middleware"
	"memoflow/pkg/res"
	"net/http"
	"time"
)

const (
	GroupByDay   = "day"
	GroupByMonth = "month"
)

func NewStatHandler(router *http.ServeMux, deps StatHandlerDeps) {
	handler := &StatHandler{
		StatRepository: deps.StatRepository,
		Config:         deps.Config,
	}
	router.Handle("/stat", middleware.IsAuthed(handler.GetStat(), deps.Config))
}

type StatHandlerDeps struct {
	StatRepository *StatRepository
	Config         *configs.Config
}

type StatHandler struct {
	StatRepository *StatRepository
	Config         *configs.Config
}

func (s *StatHandler) GetStat() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		from, err := time.Parse("2006-01-02", r.URL.Query().Get("from"))
		if err != nil {
			http.Error(w, "Invalid from parametr", http.StatusBadRequest)
			return
		}
		to, err := time.Parse("2006-01-02", r.URL.Query().Get("to"))
		if err != nil {
			http.Error(w, "Invalid to parametr", http.StatusBadRequest)
			return
		}
		by := r.URL.Query().Get("by")
		if by != GroupByDay && by != GroupByMonth {
			http.Error(w, "Invalid by parametr", http.StatusBadRequest)
			return
		}
		stats := s.StatRepository.GetStats(by, from, to)
		res.Json(w, stats, http.StatusOK)
	}
}
