package stat

import (
	"log"
	"memoflow/pkg/di"
	"memoflow/pkg/event"
)

func NewStatService(deps *StatServiceDeps) *StatService {
	return &StatService{
		EventBus:       deps.EventBus,
		StatRepository: deps.StatRepository,
	}
}

type StatServiceDeps struct {
	EventBus       *event.EventBus
	StatRepository di.IStatRepository
}

type StatService struct {
	EventBus       *event.EventBus
	StatRepository di.IStatRepository
}

func (s *StatService) AddClick() {
	for msg := range s.EventBus.Subscribe() {
		if msg.Type == event.EventMemoVisited {
			id, ok := msg.Data.(uint)
			if !ok {
				log.Printf("Bad EventMemoVisited Data: %w", msg.Data)
				continue
			}
			s.StatRepository.AddClick(id)
		}
	}

}
