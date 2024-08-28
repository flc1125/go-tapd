package tapd

import (
	"encoding/json"
	"errors"
	"strings"
)

// EventType represents the type of webhook event.
type EventType string

const (
	EventTypeStoryCreate      EventType = "story::create"
	EventTypeStoryUpdate      EventType = "story::update"
	EventTypeTaskUpdate       EventType = "task::update"
	EventTypeBugCreate        EventType = "bug::create"
	EventTypeBugUpdate        EventType = "bug::update"
	EventTypeBugCommentUpdate EventType = "bug_comment::update"
)

func (e EventType) String() string {
	return string(e)
}

// ParseWebhookEvent parses the webhook event from the payload.
func ParseWebhookEvent(payload []byte) (EventType, any, error) {
	var raw map[string]interface{}
	if err := json.Unmarshal(payload, &raw); err != nil {
		return "", nil, err
	}

	// get event type
	eventType, ok := raw["event"].(string)
	if !ok {
		return "", nil, errors.New("tapd: webhook event type not found")
	}

	// decode event
	switch EventType(eventType) {
	// todo: add more event types
	case EventTypeStoryCreate:
		return decodeWebhookEvent[StoryCreateEvent](EventTypeStoryCreate, payload)
	case EventTypeStoryUpdate:
		return decodeWebhookEvent[StoryUpdateEvent](EventTypeStoryUpdate, payload)
	case EventTypeBugCreate:
		return decodeWebhookEvent[BugCreateEvent](EventTypeBugCreate, payload)
	default:
		return "", nil, errors.New("tapd: webhook event not supported")
	}
}

// decodeWebhookEvent decodes the webhook event from the payload.
func decodeWebhookEvent[T any](eventType EventType, payload []byte) (EventType, *T, error) {
	var event T
	if err := json.Unmarshal(payload, &event); err != nil {
		return eventType, nil, err
	}
	return eventType, &event, nil
}

// EventChangeFields represents the changed fields in the webhook event.
type EventChangeFields []string

var _ json.Marshaler = (*EventChangeFields)(nil)
var _ json.Unmarshaler = (*EventChangeFields)(nil)

func (f EventChangeFields) MarshalJSON() ([]byte, error) {
	if f == nil {
		return json.Marshal(nil)
	}
	return json.Marshal(strings.Join(f, ","))
}

func (f *EventChangeFields) UnmarshalJSON(data []byte) error {
	if f == nil {
		return errors.New("tapd: unmarshal nil pointer")
	}

	var raw string
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	*f = strings.Split(raw, ",")
	return nil
}
