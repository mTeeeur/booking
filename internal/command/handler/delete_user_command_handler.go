package handler

import (
	"context"

	"github.com/google/uuid"
	"github.com/mteeeur/booking/internal/command"
)

type (
	DeleteUserCommandHandler struct {
		userRepository DeleteUserCommandHandlerRepository
	}

	DeleteUserCommandHandlerRepository interface {
		DeleteUser(ctx context.Context, id uuid.UUID) error
	}
)

func NewDeleteUserCommandHandler(userRepository DeleteUserCommandHandlerRepository) DeleteUserCommandHandler {
	return DeleteUserCommandHandler{
		userRepository: userRepository,
	}
}

func (handler DeleteUserCommandHandler) Handle(ctx context.Context, command command.DeleteUserCommand) error {
	err := handler.userRepository.DeleteUser(ctx, command.ID())
	if err != nil {
		return err
	}

	return nil
}
