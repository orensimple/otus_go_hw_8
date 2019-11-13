package services

import (
	"context"
	"time"

	"github.com/orensimple/otus_hw1_8/internal/domain/interfaces"
	"github.com/orensimple/otus_hw1_8/internal/domain/models"
	uuid "github.com/satori/go.uuid"
)

type EventService struct {
	EventStorage interfaces.EventStorage
}

func (es *EventService) CreateEvent(ctx context.Context, owner, title, text string, startTime *time.Time, endTime *time.Time) (*models.Event, error) {
	// TODO: persistence, validation
	event := &models.Event{
		Id:        uuid.NewV4(),
		Owner:     owner,
		Title:     title,
		Text:      text,
		StartTime: startTime,
		EndTime:   endTime,
	}
	err := es.EventStorage.SaveEvent(ctx, event)
	if err != nil {
		return nil, err
	}
	return event, nil
}
