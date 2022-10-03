package command

import (
	"time"

	"github.com/google/uuid"
)

type UpdateUserCommand struct {
	id        uuid.UUID
	updatedAt time.Time
	email     string
}

func NewUpdateUserCommand(id uuid.UUID, updatedAt time.Time, email string) UpdateUserCommand {
	return UpdateUserCommand{
		id:        id,
		updatedAt: updatedAt,
		email:     email,
	}
}

func (command UpdateUserCommand) ID() uuid.UUID { return command.id }

func (command UpdateUserCommand) UpdatedAt() time.Time { return command.updatedAt }

func (command UpdateUserCommand) Email() string { return command.email }
