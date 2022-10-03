package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Entity

	email string `bson:"email"`
}

func NewUser(id uuid.UUID, createdAt, updatedAt time.Time, email string) User {
	return User{
		Entity: Entity{
			id:        id,
			createdAt: createdAt,
			updatedAt: updatedAt,
		},
		email: email,
	}
}

func (user *User) Email() string {
	return user.email
}

func (user *User) SetEmail(email string) {
	user.email = email
}
