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

func TestStoryService_GetStoryTemplates(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/stories/template_list", r.URL.Path)

		assert.Equal(t, "11112222", r.URL.Query().Get("workspace_id"))
		assert.Equal(t, "1", r.URL.Query().Get("workitem_type_id"))

		_, _ = w.Write(loadData(t, ".testdata/api/story/get_story_templates.json"))
	}))

	templates, _, err := client.StoryService.GetStoryTemplates(ctx, &GetStoryTemplatesRequest{
		WorkspaceID:    Ptr(11112222),
		WorkitemTypeID: Ptr(1),
	})
	assert.NoError(t, err)
	assert.True(t, len(templates) > 0)
	assert.Equal(t, "1111112222001000015", templates[0].ID)
	assert.Equal(t, "System default template", templates[0].Name)
	assert.Equal(t, "Auto created by the system", templates[0].Description)
	assert.Equal(t, "0", templates[0].Sort)
	assert.Equal(t, "0", templates[0].Default)
	assert.Equal(t, "SYSTEM", templates[0].Creator)
	assert.Equal(t, "1", templates[0].EditorType)
}

func TestStoryService_GetStoryTemplateFields(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/stories/get_default_story_template", r.URL.Path)

		assert.Equal(t, "11112222", r.URL.Query().Get("workspace_id"))
		assert.Equal(t, "1111111111111", r.URL.Query().Get("template_id"))

		_, _ = w.Write(loadData(t, ".testdata/api/story/get_story_template_fields.json"))
	}))

	fields, _, err := client.StoryService.GetStoryTemplateFields(ctx, &GetStoryTemplateFieldsRequest{
		WorkspaceID: Ptr(11112222),
		TemplateID:  Ptr(1111111111111),
	})
	assert.NoError(t, err)
	assert.True(t, len(fields) > 0)
	assert.Equal(t, "1111112222001000113", fields[0].ID)
	assert.Equal(t, "11112222", fields[0].WorkspaceID)
	assert.Equal(t, "story", fields[0].Type)
	assert.Equal(t, "1111112222001000015", fields[0].TemplateID)
	assert.Equal(t, "name", fields[0].Field)
	assert.Equal(t, "", fields[0].Value)
	assert.Equal(t, "1", fields[0].Required)
	assert.Equal(t, "0", fields[0].Sort)
	assert.Equal(t, "", fields[0].LinkageRules)
}

func TestStoryService_GetRemovedStories(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/stories/get_removed_stories", r.URL.Path)

		assert.Equal(t, "11112222", r.URL.Query().Get("workspace_id"))
		assert.Equal(t, "1111111111111,1111111111112", r.URL.Query().Get("id"))
		assert.Equal(t, "creator", r.URL.Query().Get("creator"))
		assert.Equal(t, "1", r.URL.Query().Get("is_archived"))
		assert.Equal(t, "2021-01-01", r.URL.Query().Get("created"))
		assert.Equal(t, "2021-01-02", r.URL.Query().Get("deleted"))
		assert.Equal(t, "10", r.URL.Query().Get("limit"))
		assert.Equal(t, "1", r.URL.Query().Get("page"))

		_, _ = w.Write(loadData(t, ".testdata/api/story/get_removed_stories.json"))
	}))

	stories, _, err := client.StoryService.GetRemovedStories(ctx, &GetRemovedStoriesRequest{
		WorkspaceID: Ptr(11112222),
		ID:          NewMulti(1111111111111, 1111111111112),
		Creator:     Ptr("creator"),
		IsArchived:  Ptr(1),
		Created:     Ptr("2021-01-01"),
		Deleted:     Ptr("2021-01-02"),
		Limit:       Ptr(10),
		Page:        Ptr(1),
	})
	assert.NoError(t, err)
	assert.True(t, len(stories) > 0)
	assert.Equal(t, "1111112222001069791", stories[0].ID)
	assert.Equal(t, "測試測試", stories[0].Name)
	assert.Equal(t, "张三", stories[0].Creator)
	assert.Equal(t, "2024-08-20 11:22:49", stories[0].Created)
	assert.Equal(t, "张三", stories[0].OperationUser)
	assert.Equal(t, "2024-08-20 11:28:23", stories[0].Deleted)
	assert.Equal(t, "0", stories[0].IsArchived)
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
