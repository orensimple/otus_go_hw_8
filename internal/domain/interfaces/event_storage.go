package interfaces

import (
	"context"

	"github.com/orensimple/otus_hw1_8/internal/domain/models"
)

// EventStorage interface
type EventStorage interface {
	CreateEvent(ctx context.Context, event *models.Event) error
	UpdateEvent(ctx context.Context, event *models.Event) error
	DeleteEvent(ctx context.Context, event *models.Event) error
	GetEventByID(ctx context.Context, ID int64) (*models.Event, error)
}
