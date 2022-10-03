package event

import "github.com/google/uuid"

type CreateMeetingEvent struct {
	Owner     string    `json:"owner"`
	MeetingID uuid.UUID `json:"meetingId"`
}

func NewCreateMeetingEvent(owner string, meetingID uuid.UUID) CreateMeetingEvent {
	return CreateMeetingEvent{
		Owner:     owner,
		MeetingID: meetingID,
	}
}
