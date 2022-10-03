package query

import "github.com/google/uuid"

type GetUserQuery struct {
	id uuid.UUID
}

func NewGetUserQuery(id uuid.UUID) GetUserQuery {
	return GetUserQuery{
		id: id,
	}
}

func (query GetUserQuery) ID() uuid.UUID { return query.id }
