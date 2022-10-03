package handler

import (
	"context"

	"github.com/google/uuid"
	"github.com/mteeeur/booking/internal/command"
	"github.com/mteeeur/booking/internal/domain/entity"
	"github.com/mteeeur/booking/internal/domain/event"
)

type (
	CreateMeetingCommandHandler struct {
		meetingRepository     CreateMeetingCommandHandlerRepository
		meetingService        CreateMeetingCommandHandlerService
		meetingEventPublusher CreateMeetingCommandHandlerPublisher
	}

	CreateMeetingCommandHandlerRepository interface {
		InsertMeeting(ctx context.Context, meeting entity.Meeting) error
		DeleteMeeting(ctx context.Context, id uuid.UUID) error
	}

	CreateMeetingCommandHandlerService interface {
		CreateMeeting(ctx context.Context, meeting entity.Meeting) error
		DeleteMeeting(ctx context.Context, id uuid.UUID) error
	}

	CreateMeetingCommandHandlerPublisher interface {
		PublishCreateMeetingEvent(ctx context.Context, event event.CreateMeetingEvent) error
	}
)

func NewCreateMeetingCommandHandler(
	meetingRepository CreateMeetingCommandHandlerRepository,
	meetingService CreateMeetingCommandHandlerService,
	meetingEventPublisher CreateMeetingCommandHandlerPublisher,
) CreateMeetingCommandHandler {
	return CreateMeetingCommandHandler{
		meetingRepository:     meetingRepository,
		meetingService:        meetingService,
		meetingEventPublusher: meetingEventPublisher,
	}
}

func (handler *CreateMeetingCommandHandler) Handle(
	ctx context.Context,
	command command.CreateMeetingCommand,
) (uuid.UUID, error) {
	meeting := entity.NewMeeting(
		command.ID(),
		command.CreatedAt(),
		command.UpdatedAt(),
		command.Location(),
		command.Owner(),
		command.Subject(),
	)

	err := handler.meetingService.CreateMeeting(ctx, meeting)
	if err != nil {
		return uuid.UUID{}, err
	}

	err = handler.meetingRepository.InsertMeeting(ctx, meeting)
	if err != nil {
		return uuid.UUID{}, err
	}

	createMeetingEvent := event.NewCreateMeetingEvent(meeting.Owner(), meeting.ID())
	err = handler.meetingEventPublusher.PublishCreateMeetingEvent(ctx, createMeetingEvent)
	if err != nil {
		deleteMeetingErr := handler.meetingService.DeleteMeeting(ctx, meeting.ID())
		if deleteMeetingErr != nil {
			return uuid.UUID{}, deleteMeetingErr
		}
	}

	return meeting.ID(), nil
}
