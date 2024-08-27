package tapd

import (
	"encoding/json"
	"strings"
)

type EventType string

const (
	StoryCreateEventType      EventType = "story::create"
	StoryUpdateEventType      EventType = "story::update"
	TaskUpdateEventType       EventType = "task::update"
	BugCreateEventType        EventType = "bug::create"
	BugUpdateEventType        EventType = "bug::update"
	BugCommentUpdateEventType EventType = "bug_comment::update"
)

type ChangeFields []string

func (c *ChangeFields) UnmarshalJSON(data []byte) error {
	var raw string
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	*c = strings.Split(raw, ",")
	return nil
}

type StoryUpdateEvent struct {
	Event        string            `json:"event"`
	EventFrom    string            `json:"event_from"`
	Referer      string            `json:"referer"`
	WorkspaceID  string            `json:"workspace_id"`
	CurrentUser  string            `json:"current_user"`
	ID           string            `json:"id"`
	ChangeFields ChangeFields      `json:"change_fields"`
	Old          map[string]string `json:"old"`
	New          map[string]string `json:"new"`
	Secret       string            `json:"secret"`
	RioToken     string            `json:"rio_token"`
	DevProxyHost string            `json:"devproxy_host"`
	QueueID      string            `json:"queue_id"`
	EventID      string            `json:"event_id"`
	Created      string            `json:"created"`
}

var _ json.Unmarshaler = (*StoryUpdateEvent)(nil)

func (e *StoryUpdateEvent) UnmarshalJSON(data []byte) error {
	type Alias StoryUpdateEvent
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(e),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	var raw map[string]string
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	e.Old = make(map[string]string)
	e.New = make(map[string]string)

	for k, v := range raw {
		if strings.HasPrefix(k, "old_") {
			e.Old[strings.TrimPrefix(k, "old_")] = v
		} else if strings.HasPrefix(k, "new_") {
			e.New[strings.TrimPrefix(k, "new_")] = v
		}
	}

	return nil
}
