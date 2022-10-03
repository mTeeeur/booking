package command

import "github.com/google/uuid"

type DeleteUserCommand struct {
	id uuid.UUID
}

func NewDeleteUserCommand(id uuid.UUID) DeleteUserCommand {
	return DeleteUserCommand{
		id: id,
	}
}

func (command DeleteUserCommand) ID() uuid.UUID { return command.id }
