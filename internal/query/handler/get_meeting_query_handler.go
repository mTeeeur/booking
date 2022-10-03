package handler

import (
	"context"

	"github.com/google/uuid"
	"github.com/mteeeur/booking/internal/domain/entity"
	"github.com/mteeeur/booking/internal/query"
)

type (
	GetMeetingQueryHandler struct {
		meetingRepository GetMeetinger
	}

	GetMeetinger interface {
		GetMeeting(ctx context.Context, id uuid.UUID) (entity.Meeting, error)
	}
)

func NewGetMeetingQueryHandler(meetingRepository GetMeetinger) GetMeetingQueryHandler {
	return GetMeetingQueryHandler{
		meetingRepository: meetingRepository,
	}
}

func (handler *GetMeetingQueryHandler) Handle(ctx context.Context, query query.GetMeetingQuery) (entity.Meeting, error) {
	meeting, err := handler.meetingRepository.GetMeeting(ctx, query.ID())
	if err != nil {
		return entity.Meeting{}, err
	}

	return meeting, nil
}
