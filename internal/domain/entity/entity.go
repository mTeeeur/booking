package entity

import (
	"time"

	"github.com/google/uuid"
)

type Entity struct {
	id        uuid.UUID `bson:"id"`
	createdAt time.Time `bson:"createdAt"`
	updatedAt time.Time `bson:"updatedAt"`
}

func (e *Entity) ID() uuid.UUID {
	return e.id
}

func (e *Entity) CreatedAt() time.Time {
	return e.createdAt
}

func (e *Entity) UpdatedAt() time.Time {
	return e.updatedAt
}

func (e *Entity) SetUpdatedAt(updatedAt time.Time) {
	e.updatedAt = updatedAt
}
