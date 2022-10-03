package handler

import (
	"context"

	"github.com/google/uuid"
	"github.com/mteeeur/booking/internal/command"
	"github.com/mteeeur/booking/internal/domain/entity"
)

type (
	CreateUserCommandHandler struct {
		userRepository CreateUserCommandHandlerRepository
	}

	CreateUserCommandHandlerRepository interface {
		InsertUser(ctx context.Context, user entity.User) error
	}
)

func NewCreateUserCommandHandler(userRepository CreateUserCommandHandlerRepository) CreateUserCommandHandler {
	return CreateUserCommandHandler{
		userRepository: userRepository,
	}
}

func (handler CreateUserCommandHandler) Handle(ctx context.Context, command command.CreateUserCommand) (uuid.UUID, error) {
	user := entity.NewUser(
		command.ID(),
		command.CreatedAt(),
		command.UpdatedAt(),
		command.Email(),
	)

	err := handler.userRepository.InsertUser(ctx, user)
	if err != nil {
		return uuid.UUID{}, err
	}

	return user.ID(), nil
}
