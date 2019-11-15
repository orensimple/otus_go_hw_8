package memory

import (
	"context"

	"github.com/orensimple/otus_hw1_8/internal/domain/models"
)

// SaveEvent asdf
func (pges *PgEventStorage) SaveEvent(ctx context.Context, event *models.Event) error {

	return nil
}

// UpdateEvent asdf
func (pges *PgEventStorage) UpdateEvent(ctx context.Context, event *models.Event) error {

	return nil
}

// DeleteEvent asdf
func (pges *PgEventStorage) DeleteEvent(ctx context.Context, event *models.Event) error {

	return nil
}

// GetEventById asdf
func (pges *PgEventStorage) GetEventById(ctx context.Context, id string) (*models.Event, error) {

	return nil, nil
}
