package publusher

import (
	"context"
	"encoding/json"

	"github.com/mteeeur/booking/internal/domain/event"
	amqp "github.com/rabbitmq/amqp091-go"
)

type MeetingEventPublisher struct {
	channel amqp.Channel
}

func (publisher *MeetingEventPublisher) PublishCreateMeetingEvent(ctx context.Context, event event.CreateMeetingEvent) error {
	eventInBytes, err := json.Marshal(event)
	if err != nil {
		return err
	}

	err = publisher.channel.PublishWithContext(
		ctx,
		"",
		"meeting.event.created",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        eventInBytes,
		},
	)
	if err != nil {
		return err
	}

	return nil
}
