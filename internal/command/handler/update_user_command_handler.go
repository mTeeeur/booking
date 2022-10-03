package handler

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/mteeeur/booking/internal/command"
	"github.com/mteeeur/booking/internal/domain/entity"
)

type (
	UpdateUserCommandHandler struct {
		userRepository UpdateUserCommandHandlerRepository
	}

	UpdateUserCommandHandlerRepository interface {
		UpdateUser(ctx context.Context, user entity.User) error
	}
)

func NewUpdateUserCommandHandler(userRepository UpdateUserCommandHandlerRepository) UpdateUserCommandHandler {
	return UpdateUserCommandHandler{
		userRepository: userRepository,
	}
}

func (handler *UpdateUserCommandHandler) Handle(ctx context.Context, command command.UpdateUserCommand) (uuid.UUID, error) {
	user := entity.NewUser(
		command.ID(),
		time.Time{}, /* craetedAt */
		command.UpdatedAt(),
		command.Email(),
	)

	err := handler.userRepository.UpdateUser(ctx, user)
	if err != nil {
		return uuid.UUID{}, err
	}

	return user.ID(), nil
}
