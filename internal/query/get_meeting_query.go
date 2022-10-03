package query

import "github.com/google/uuid"

type GetMeetingQuery struct {
	id uuid.UUID
}

func NewGetMeetingQuery(id uuid.UUID) GetMeetingQuery {
	return GetMeetingQuery{
		id: id,
	}
}

func (query GetMeetingQuery) ID() uuid.UUID { return query.id }
