package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/mteeeur/booking/internal/domain/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MeetingRpository struct {
	collection *mongo.Collection
}

func NewMeetingRpository(collection *mongo.Collection) *MeetingRpository {
	return &MeetingRpository{
		collection: collection,
	}
}

func (repostiory *MeetingRpository) InsertMeeting(ctx context.Context, meeting entity.Meeting) error {
	repostiory.collection.InsertOne(
		ctx,
		bson.M{
			"id":        meeting.ID(),
			"createdAt": meeting.CreatedAt(),
			"updatedAt": meeting.UpdatedAt(),
			"owner":     meeting.Owner(),
			"subject":   meeting.Subject(),
			"location":  meeting.Location(),
		},
	)
	return nil
}

func (repository *MeetingRpository) GetMeeting(ctx context.Context, id uuid.UUID) (entity.Meeting, error) {
	res := repository.collection.FindOne(
		ctx,
		bson.M{"id": id},
	)
	if err := res.Err(); err != nil {
		return entity.Meeting{}, err
	}

	var meeting entity.Meeting
	err := res.Decode(&meeting)
	if err != nil {
		return entity.Meeting{}, err
	}

	return meeting, nil
}
