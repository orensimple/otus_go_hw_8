package interfaces

import (
	"context"

	"github.com/orensimple/otus_hw1_8/internal/domain/models"
)
// EventStorage interface
type EventStorage interface {
	SaveEvent(ctx context.Context, event *models.Event) error
	UpdateEvent(ctx context.Context, event *models.Event) error
	DeleteEvent(ctx context.Context, event *models.Event) error
	GetEventById(ctx context.Context, id string) (*models.Event, error)
}
