package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/mteeeur/booking/internal/domain/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(collection *mongo.Collection) *UserRepository {
	return &UserRepository{
		collection: collection,
	}
}

func (repository *UserRepository) InsertUser(ctx context.Context, user entity.User) error {
	_, err := repository.collection.InsertOne(
		ctx,
		bson.M{
			"id":        user.ID(),
			"createdAt": user.CreatedAt(),
			"updatedAt": user.UpdatedAt(),
			"email":     user.Email(),
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func (repository *UserRepository) GetUser(ctx context.Context, id uuid.UUID) (entity.User, error) {
	res := repository.collection.FindOne(
		ctx,
		bson.M{"id": id},
	)
	if err := res.Err(); err != nil {
		return entity.User{}, err
	}

	var user entity.User
	err := res.Decode(&user)
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (repository *UserRepository) UpdateUser(ctx context.Context, user entity.User) error {
	_, err := repository.collection.UpdateOne(
		ctx,
		bson.M{"id": user.ID()},
		bson.M{
			"$set": bson.M{
				"updatedAt": user.UpdatedAt(),
				"email":     user.Email(),
			},
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func (repository *UserRepository) DeleteUser(ctx context.Context, id uuid.UUID) error {
	_, err := repository.collection.DeleteOne(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
