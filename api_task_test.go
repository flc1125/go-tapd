package tapd

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTaskService_GetTasks(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/tasks", r.URL.Path)

		_, _ = w.Write(loadData(t, ".testdata/api/task/get_tasks.json"))
	}))

	tasks, _, err := client.TaskService.GetTasks(ctx, &GetTasksRequest{
		WorkspaceID: Ptr(11112222),
		Status:      Enum(TaskStatusOpen, TaskStatusDone),
		Fields:      Multi("id", "workspace_id"),
	})
	assert.NoError(t, err)
	assert.True(t, len(tasks) > 0)
}

func TestTaskService_GetTasksCount(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/tasks/count", r.URL.Path)
		assert.Equal(t, "11112222", r.URL.Query().Get("workspace_id"))
		assert.Equal(t, "open|done", r.URL.Query().Get("status"))

		_, _ = w.Write(loadData(t, ".testdata/api/task/get_tasks_count.json"))
	}))

	count, _, err := client.TaskService.GetTasksCount(ctx, &GetTasksCountRequest{
		WorkspaceID: Ptr(11112222),
		Status:      Enum(TaskStatusOpen, TaskStatusDone),
	})
	assert.NoError(t, err)
	assert.Equal(t, 36, count)
}

func TestTaskService_GetTaskFieldsInfo(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/tasks/get_fields_info", r.URL.Path)
		assert.Equal(t, "11112222", r.URL.Query().Get("workspace_id"))

		_, _ = w.Write(loadData(t, ".testdata/api/task/get_task_fields_info.json"))
	}))

	fields, _, err := client.TaskService.GetTaskFieldsInfo(ctx, &GetTaskFieldsInfoRequest{
		WorkspaceID: Ptr(11112222),
	})
	assert.NoError(t, err)
	assert.True(t, len(fields) > 0)

	var flag1, flag2 bool

	// slice with id name
	for _, field := range fields {
		if field.Name == "id" {
			assert.Equal(t, TaskFieldsInfoHTMLTypeInput, field.HTMLType)
			assert.Equal(t, "ID", field.Label)
			flag1 = true
		}

		if field.Name == "iteration_id" {
			flag2 = true
			assert.Equal(t, "迭代", field.Label)
			assert.Equal(t, TaskFieldsInfoHTMLTypeSelect, field.HTMLType)
			assert.Contains(t, field.Options, TaskFieldsInfoOption{
				Value: "1111112222001001246",
				Label: "迭代2",
			})
			assert.Contains(t, field.PureOptions, TaskFieldsInfoPureOption{
				ParentID:    "0",
				WorkspaceID: "11112222",
				Sort:        "100124600000",
				Value:       "1111112222001001246",
				Label:       "迭代2",
				Panel:       0,
			})
		}
	}
	assert.True(t, flag1)
	assert.True(t, flag2)
}
