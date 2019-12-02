package api

import (
	"context"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"
	"net"
	"time"

	"github.com/orensimple/otus_hw1_8/internal/domain/errors"
	"github.com/orensimple/otus_hw1_8/internal/domain/services"
)

type CalendarServer struct {
	EventService *services.EventService
}

func (cs *CalendarServer) CreateEvent(ctx context.Context, req *CreateEventRequest) (*CreateEventResponse, error) {

	startTime := (*time.Time)(nil)
	if req.GetStartTime() != nil {
		st, err := ptypes.Timestamp(req.GetStartTime())
		if err != nil {
			return nil, err
		}
		startTime = &st
	}

	endTime := (*time.Time)(nil)
	if req.GetEndTime() != nil {
		et, err := ptypes.Timestamp(req.GetEndTime())
		if err != nil {
			return nil, err
		}
		endTime = &et
	}

	event, err := cs.EventService.CreateEvent(ctx, req.GetID(), req.GetOwner(), req.GetTitle(), req.GetText(), *startTime, *endTime)
	if err != nil {
		if berr, ok := err.(errors.EventError); ok {
			resp := &CreateEventResponse{
				Result: &CreateEventResponse_Error{
					Error: string(berr),
				},
			}
			return resp, nil
		} else {
			return nil, err
		}
	}
	protoEvent := &Event{
		ID:    event.ID,
		Title: event.Title,
		Text:  event.Text,
		Owner: event.Owner,
	}
	if protoEvent.StartTime, err = ptypes.TimestampProto(event.StartTime); err != nil {
		return nil, err
	}
	if protoEvent.EndTime, err = ptypes.TimestampProto(event.EndTime); err != nil {
		return nil, err
	}
	resp := &CreateEventResponse{
		Result: &CreateEventResponse_Event{
			Event: protoEvent,
		},
	}
	return resp, nil
}

func (cs *CalendarServer) UpdateEvent(ctx context.Context, req *UpdateEventRequest) (*UpdateEventResponse, error) {

	startTime := (*time.Time)(nil)
	if req.GetStartTime() != nil {
		st, err := ptypes.Timestamp(req.GetStartTime())
		if err != nil {
			return nil, err
		}
		startTime = &st
	}

	endTime := (*time.Time)(nil)
	if req.GetEndTime() != nil {
		et, err := ptypes.Timestamp(req.GetEndTime())
		if err != nil {
			return nil, err
		}
		endTime = &et
	}

	event, err := cs.EventService.UpdateEvent(ctx, req.GetID(), req.GetOwner(), req.GetTitle(), req.GetText(), *startTime, *endTime)
	if err != nil {
		if berr, ok := err.(errors.EventError); ok {
			resp := &UpdateEventResponse{
				Result: &UpdateEventResponse_Error{
					Error: string(berr),
				},
			}
			return resp, nil
		} else {
			return nil, err
		}
	}
	protoEvent := &Event{
		ID:    event.ID,
		Title: event.Title,
		Text:  event.Text,
		Owner: event.Owner,
	}
	if protoEvent.StartTime, err = ptypes.TimestampProto(event.StartTime); err != nil {
		return nil, err
	}
	if protoEvent.EndTime, err = ptypes.TimestampProto(event.EndTime); err != nil {
		return nil, err
	}
	resp := &UpdateEventResponse{
		Result: &UpdateEventResponse_Event{
			Event: protoEvent,
		},
	}
	return resp, nil
}

func (cs *CalendarServer) DeleteEvent(ctx context.Context, req *DeleteEventRequest) (*DeleteEventResponse, error) {

	err := cs.EventService.DeleteEvent(ctx, req.GetID())
	if err != nil {
		if berr, ok := err.(errors.EventError); ok {
			resp := &DeleteEventResponse{
				Result: &DeleteEventResponse_Error{
					Error: string(berr),
				},
			}
			return resp, nil
		} else {
			return nil, err
		}
	}

	resp := &DeleteEventResponse{
		Result: &DeleteEventResponse_Error{
			Error: "",
		},
	}
	return resp, nil
}

func (cs *CalendarServer) Serve(addr string) error {
	s := grpc.NewServer()
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	RegisterCalendarServiceServer(s, cs)
	return s.Serve(l)
}
