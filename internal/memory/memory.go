package memory

import (
	"context"
	"sync"

	"github.com/orensimple/otus_hw1_8/internal/domain/errors"
	"github.com/orensimple/otus_hw1_8/internal/domain/models"
)

// MemEventStorage slice
type MemEventStorage struct {
	events map[int64]*models.Event
	mutex  *sync.Mutex
}

// NewMemEventStorage init
func NewMemEventStorage() *MemEventStorage {
	return &MemEventStorage{
		events: make(map[int64]*models.Event),
		mutex:  new(sync.Mutex),
	}
}

// SaveEvent to mem
func (mem *MemEventStorage) SaveEvent(ctx context.Context, event *models.Event) error {
	if _, ok := mem.events[event.ID]; ok {
		return errors.ErrEventExist

	}
	mem.mutex.Lock()
	mem.events[event.ID] = event
	mem.mutex.Unlock()

	return nil
}

// UpdateEvent to mem
func (mem *MemEventStorage) UpdateEvent(ctx context.Context, event *models.Event) (*models.Event, error) {
	if _, ok := mem.events[event.ID]; ok {
		mem.mutex.Lock()
		mem.events[event.ID] = event
		mem.mutex.Unlock()

		return event, nil
	}

	return nil, errors.ErrEventNotFound
}

// GetEvents to mem
func (mem *MemEventStorage) GetEvents(ctx context.Context) ([]*models.Event, error) {
	Events := make([]*models.Event, 0)
	mem.mutex.Lock()
	for _, bm := range mem.events {
		Events = append(Events, bm)
	}
	mem.mutex.Unlock()

	return Events, nil
}

// DeleteEvent from mem
func (mem *MemEventStorage) DeleteEvent(ctx context.Context, ID int64) error {
	mem.mutex.Lock()
	defer mem.mutex.Unlock()

	_, ex := mem.events[ID]
	if ex {
		delete(mem.events, ID)
		return nil
	}

	return errors.ErrEventNotFound
}
