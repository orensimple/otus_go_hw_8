package memory

import (
	"context"

	"github.com/orensimple/otus_hw1_8/internal/domain/models"
)

// MemEventStorage slice
type MemEventStorage struct {
	main []models.Event
}

// CreateEvent asdf
func (mem *MemEventStorage) CreateEvent(ctx context.Context, event *models.Event) error {

	return nil
}

// UpdateEvent asdf
func (mem *MemEventStorage) UpdateEvent(ctx context.Context, event *models.Event) error {

	return nil
}

// DeleteEvent asdf
func (mem *MemEventStorage) DeleteEvent(ctx context.Context, event *models.Event) error {

	return nil
}

// GetEventByID asDf
func (mem *MemEventStorage) GetEventByID(ctx context.Context, ID int64) (*models.Event, error) {

	return nil, nil
}
