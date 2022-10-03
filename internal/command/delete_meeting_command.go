package command

import "github.com/google/uuid"

type DeleteMeetingCommand struct {
	id uuid.UUID
}

func NewDeleteMeetingCommand(id uuid.UUID) DeleteMeetingCommand {
	return DeleteMeetingCommand{
		id: id,
	}
}
