package tapd

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStoryService_GetStoryCustomFieldsSettings(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/stories/custom_fields_settings", r.URL.Path)

		assert.Equal(t, "11112222", r.URL.Query().Get("workspace_id"))

		_, _ = w.Write(loadData(t, ".testdata/api/story/get_story_custom_fields_settings.json"))
	}))

	settings, _, err := client.StoryService.GetStoryCustomFieldsSettings(ctx, &GetStoryCustomFieldsSettingsRequest{
		WorkspaceID: Ptr(11112222),
	})
	assert.NoError(t, err)
	assert.True(t, len(settings) > 0)
	assert.Equal(t, "1111112222001000155", settings[0].ID)
	assert.Equal(t, "11112222", settings[0].WorkspaceID)
	assert.Equal(t, "1", settings[0].AppID)
	assert.Equal(t, "story", settings[0].EntryType)
	assert.Equal(t, "custom_field_100", settings[0].CustomField)
	assert.Equal(t, "user_chooser", settings[0].Type)
	assert.Equal(t, "test name", settings[0].Name)
	assert.Nil(t, settings[0].Options)
	assert.Nil(t, settings[0].ExtraConfig)
	assert.Equal(t, "1", settings[0].Enabled)
	assert.Equal(t, "0", settings[0].Freeze)
	assert.Nil(t, settings[0].Sort)
	assert.Nil(t, settings[0].Memo)
	assert.Equal(t, "", settings[0].OpenExtensionID)
	assert.Equal(t, 0, settings[0].IsOut)
	assert.Equal(t, 0, settings[0].IsUninstall)
	assert.Equal(t, "", settings[0].AppName)
}

func TestStoryService_GetStoryRelatedBugs(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/stories/get_related_bugs", r.URL.Path)

		assert.Equal(t, "11112222", r.URL.Query().Get("workspace_id"))
		assert.Equal(t, "33334444,55556666", r.URL.Query().Get("story_id"))

		_, _ = w.Write(loadData(t, ".testdata/api/story/get_story_related_bugs.json"))
	}))

	relatedBugs, _, err := client.StoryService.GetStoryRelatedBugs(ctx, &GetStoryRelatedBugsRequest{
		WorkspaceID: Ptr(11112222),
		StoryID:     NewMulti(33334444, 55556666),
	})
	assert.NoError(t, err)
	assert.True(t, len(relatedBugs) > 0)
	assert.Equal(t, 11112222, relatedBugs[0].WorkspaceID)
	assert.Equal(t, "1111112222001063941", relatedBugs[0].StoryID)
	assert.Equal(t, "1111112222001035927", relatedBugs[0].BugID)
}

func TestStoryService_GetConvertStoryIDsToQueryToken(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/stories/ids_to_query_token", r.URL.Path)

		var req struct {
			WorkspaceID int    `json:"workspace_id"`
			StoryIDs    string `json:"ids"`
		}
		assert.NoError(t, json.NewDecoder(r.Body).Decode(&req))
		assert.Equal(t, 11112222, req.WorkspaceID)
		assert.Equal(t, "33334444,55556666", req.StoryIDs)

		_, _ = w.Write(loadData(t, ".testdata/api/story/get_convert_story_ids_to_query_token.json"))
	}))

	response, _, err := client.StoryService.GetConvertStoryIDsToQueryToken(ctx, &GetConvertStoryIDsToQueryTokenRequest{
		WorkspaceID: Ptr(11112222),
		StoryIDs:    NewMulti(33334444, 55556666),
	})
	assert.NoError(t, err)
	assert.Equal(t, "11111111111", response.QueryToken)
	assert.Contains(t, response.Href, "11111111111")
}
