package command

import (
	"time"

	"github.com/google/uuid"
)

type CreateMeetingCommand struct {
	id        uuid.UUID
	createdAt time.Time
	updatedAt time.Time
	owner     string
	location  string
	subject   string
}

func NewCreateMeetingCommand(
	id uuid.UUID,
	createdAt time.Time,
	updatedAt time.Time,
	location string,
	owner string,
	subject string,
) CreateMeetingCommand {
	return CreateMeetingCommand{
		id:        id,
		createdAt: createdAt,
		updatedAt: updatedAt,
		owner:     owner,
		location:  location,
		subject:   subject,
	}
}

func (command CreateMeetingCommand) ID() uuid.UUID { return command.id }

func (command CreateMeetingCommand) CreatedAt() time.Time { return command.createdAt }

func (command CreateMeetingCommand) UpdatedAt() time.Time { return command.updatedAt }

func (command CreateMeetingCommand) Owner() string { return command.owner }

func (command CreateMeetingCommand) Location() string { return command.location }

func (command CreateMeetingCommand) Subject() string { return command.subject }
