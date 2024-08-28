package tapd

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommentService_GetComments(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/comments", r.URL.Path)
		assert.Equal(t, "111,222", r.URL.Query().Get("id"))
		assert.Equal(t, "title", r.URL.Query().Get("title"))
		assert.Equal(t, "description", r.URL.Query().Get("description"))
		assert.Equal(t, "author", r.URL.Query().Get("author"))
		assert.Equal(t, CommentEntryTypeStories.String(), r.URL.Query().Get("entry_type"))
		assert.Equal(t, "123", r.URL.Query().Get("entry_id"))
		assert.Equal(t, "created", r.URL.Query().Get("created"))
		assert.Equal(t, "modified", r.URL.Query().Get("modified"))
		assert.Equal(t, "111", r.URL.Query().Get("workspace_id"))
		assert.Equal(t, "222", r.URL.Query().Get("root_id"))
		assert.Equal(t, "333", r.URL.Query().Get("reply_id"))
		assert.Equal(t, "10", r.URL.Query().Get("limit"))
		assert.Equal(t, "1", r.URL.Query().Get("page"))
		assert.Equal(t, "id desc", r.URL.Query().Get("order"))
		assert.Equal(t, "id,title", r.URL.Query().Get("fields"))

		_, _ = w.Write(loadData(t, ".testdata/api/comment/get_comments.json"))
	}))

	comments, _, err := client.CommentService.GetComments(ctx, &GetCommentsRequest{
		ID:          NewID(111, 222),
		Title:       Ptr("title"),
		Description: Ptr("description"),
		Author:      Ptr("author"),
		EntryType:   Ptr(CommentEntryTypeStories),
		EntryID:     Ptr(123),
		Created:     Ptr("created"),
		Modified:    Ptr("modified"),
		WorkspaceID: Ptr(111),
		RootID:      Ptr(222),
		ReplyID:     Ptr(333),
		Limit:       Ptr(10),
		Page:        Ptr(1),
		Order:       NewOrder("id", OrderDesc),
		Fields:      NewFields("id", "title"),
	})
	assert.NoError(t, err)
	assert.NotNil(t, comments)
	assert.True(t, len(comments) > 0)
	assert.Equal(t, "1123402991001033061", comments[0].ID)
	assert.Equal(t, "title1", comments[0].Title)
	assert.Equal(t, "description1", comments[0].Description)
	assert.Equal(t, "author1", comments[0].Author)
	assert.Equal(t, CommentEntryTypeBug, comments[0].EntryType)
	assert.Equal(t, "1123402991001037223", comments[0].EntryID)
	assert.Equal(t, "0", comments[0].ReplyID)
	assert.Equal(t, "0", comments[0].RootID)
	assert.Equal(t, "2024-08-28 10:06:24", comments[0].Created)
	assert.Equal(t, "2024-08-28 10:06:24", comments[0].Modified)
	assert.Equal(t, "11112222", comments[0].WorkspaceID)
}
