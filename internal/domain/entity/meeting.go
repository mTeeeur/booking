package entity

import (
	"time"

	"github.com/google/uuid"
)

type Meeting struct {
	Entity

	location string `bson:"location"`
	owner    string `bson:"owner"`
	subject  string `bson:"subject"`
}

func NewMeeting(
	id uuid.UUID,
	createdAt time.Time,
	updatedAt time.Time,
	location string,
	owner string,
	subject string,
) Meeting {
	return Meeting{
		Entity: Entity{
			id:        id,
			createdAt: createdAt,
			updatedAt: updatedAt,
		},
		location: location,
		owner:    owner,
		subject:  subject,
	}
}

func (meeting *Meeting) Location() string { return meeting.location }

func (meeting *Meeting) Owner() string { return meeting.owner }

func (meeting *Meeting) Subject() string { return meeting.subject }
