package handler

import (
	"context"

	"github.com/google/uuid"
	"github.com/mteeeur/booking/internal/domain/entity"
	"github.com/mteeeur/booking/internal/query"
)

type (
	GetUserQueryHandler struct {
		userRepository GetUserer
	}

	GetUserer interface {
		GetUser(ctx context.Context, id uuid.UUID) (entity.User, error)
	}
)

func NewGetUserQueryHandler(userRepository GetUserer) GetUserQueryHandler {
	return GetUserQueryHandler{
		userRepository: userRepository,
	}
}

func (handler *GetUserQueryHandler) Handler(ctx context.Context, query query.GetUserQuery) (entity.User, error) {
	user, err := handler.userRepository.GetUser(ctx, query.ID())
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}
