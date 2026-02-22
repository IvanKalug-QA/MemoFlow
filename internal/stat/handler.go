package stat

import (
	"fmt"
	"memoflow/configs"
	"net/http"
	"time"
)

const (
	FilterByDay   = "day"
	FilterByMonth = "month"
)

func NewStatHandler(router *http.ServeMux, deps StatHandlerDeps) {
	handler := &StatHandler{
		StatRepository: deps.StatRepository,
		Config:         deps.Config,
	}
	router.HandleFunc("/stat", handler.GetStat())
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
		by := r.URL.Query().Get("to")
		if by != FilterByDay && by != FilterByMonth {
			http.Error(w, "Invalid by parametr", http.StatusBadRequest)
			return
		}
		fmt.Println(from, to)
	}
}
