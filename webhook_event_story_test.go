package tapd

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWebhookEvent_Story_StoryUpdateEvent(t *testing.T) {
	var event StoryUpdateEvent
	assert.NoError(t, json.Unmarshal(loadData(t, ".testdata/webhook/story_update_event.json"), &event))
	assert.Equal(t, EventTypeStoryUpdate, event.Event)
	assert.Equal(t, "11112222", event.WorkspaceID)
	assert.Equal(t, "1111112222001069123", event.ID)
	assert.Equal(t, "1111112222001069123", *event.StoryUpdateEventOldFields.ID)
	assert.Len(t, event.ChangeFields, 2)
	assert.Equal(t, "owner", event.ChangeFields[0])

	bytes, err := json.MarshalIndent(event, "", "  ")
	assert.NoError(t, err)
	assert.Contains(t, string(bytes), "owner,modified")
}
