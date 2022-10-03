package command

import (
	"time"

	"github.com/google/uuid"
)

type CreateUserCommand struct {
	id        uuid.UUID
	createdAt time.Time
	updatedAt time.Time
	email     string
}

func NewCreateUserCommand(id uuid.UUID, createdAt, updatedAt time.Time, email string) CreateUserCommand {
	return CreateUserCommand{
		id:        id,
		createdAt: createdAt,
		updatedAt: updatedAt,
		email:     email,
	}
}

func (command CreateUserCommand) ID() uuid.UUID { return command.id }

func (command CreateUserCommand) CreatedAt() time.Time { return command.createdAt }

func (command CreateUserCommand) UpdatedAt() time.Time { return command.updatedAt }

func (command CreateUserCommand) Email() string { return command.email }
