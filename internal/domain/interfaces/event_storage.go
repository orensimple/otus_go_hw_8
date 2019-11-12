package interfaces

import (
	"context"
	"time"

	"github.com/orensimple/otus_hw1_8/internal/domain/models"
)

type EventStorage interface {
	SaveEvent(ctx context.Context, event *models.Event) error
	GetEventById(ctx context.Context, id string) (*models.Event, error)
	GetEventsByOwnerStartDate(ctx context.Context, owner string, startTime time.Time) []*models.Event
}
