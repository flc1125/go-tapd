package tapd

import (
	"context"
	"net/http"
)

// TaskStatus 任务状态
type TaskStatus string

const (
	TaskStatusOpen        TaskStatus = "open"        // 未开始
	TaskStatusProgressing TaskStatus = "progressing" // 进行中
	TaskStatusDone        TaskStatus = "done"        // 已完成
)

// Task 任务
type Task struct {
	ID                string        `json:"id,omitempty"`               // 任务ID
	Name              string        `json:"name,omitempty"`             // 任务标题
	Description       string        `json:"description,omitempty"`      // 任务详细描述
	WorkspaceID       string        `json:"workspace_id,omitempty"`     // 项目ID
	Creator           string        `json:"creator,omitempty"`          // 创建人
	Created           string        `json:"created,omitempty"`          // 创建时间
	Modified          string        `json:"modified,omitempty"`         // 最后修改时间
	Status            TaskStatus    `json:"status,omitempty"`           // 状态
	Owner             string        `json:"owner,omitempty"`            // 任务当前处理人
	Cc                string        `json:"cc,omitempty"`               // 抄送人
	Begin             string        `json:"begin,omitempty"`            // 预计开始
	Due               string        `json:"due,omitempty"`              // 预计结束
	StoryID           string        `json:"story_id,omitempty"`         // 关联需求的ID
	IterationID       string        `json:"iteration_id,omitempty"`     // 所属迭代的ID
	Priority          string        `json:"priority,omitempty"`         // 优先级
	Progress          string        `json:"progress,omitempty"`         // 进度
	Completed         string        `json:"completed,omitempty"`        // 完成时间
	EffortCompleted   string        `json:"effort_completed,omitempty"` // 完成工时
	Exceed            string        `json:"exceed,omitempty"`           // 超出工时
	Remain            string        `json:"remain,omitempty"`           // 剩余工时
	Effort            string        `json:"effort,omitempty"`           // 预估工时
	HasAttachment     string        `json:"has_attachment,omitempty"`   // 是否有附件
	ReleaseID         string        `json:"release_id,omitempty"`       // 发布计划ID
	Label             string        `json:"label,omitempty"`            // 标签
	CustomFieldOne    string        `json:"custom_field_one,omitempty"` // 自定义字段
	CustomFieldTwo    string        `json:"custom_field_two,omitempty"`
	CustomFieldThree  string        `json:"custom_field_three,omitempty"`
	CustomFieldFour   string        `json:"custom_field_four,omitempty"`
	CustomFieldFive   string        `json:"custom_field_five,omitempty"`
	CustomFieldSix    string        `json:"custom_field_six,omitempty"`
	CustomFieldSeven  string        `json:"custom_field_seven,omitempty"`
	CustomFieldEight  string        `json:"custom_field_eight,omitempty"`
	CustomField9      string        `json:"custom_field_9,omitempty"`
	CustomField10     string        `json:"custom_field_10,omitempty"`
	CustomField11     string        `json:"custom_field_11,omitempty"`
	CustomField12     string        `json:"custom_field_12,omitempty"`
	CustomField13     string        `json:"custom_field_13,omitempty"`
	CustomField14     string        `json:"custom_field_14,omitempty"`
	CustomField15     string        `json:"custom_field_15,omitempty"`
	CustomField16     string        `json:"custom_field_16,omitempty"`
	CustomField17     string        `json:"custom_field_17,omitempty"`
	CustomField18     string        `json:"custom_field_18,omitempty"`
	CustomField19     string        `json:"custom_field_19,omitempty"`
	CustomField20     string        `json:"custom_field_20,omitempty"`
	CustomField21     string        `json:"custom_field_21,omitempty"`
	CustomField22     string        `json:"custom_field_22,omitempty"`
	CustomField23     string        `json:"custom_field_23,omitempty"`
	CustomField24     string        `json:"custom_field_24,omitempty"`
	CustomField25     string        `json:"custom_field_25,omitempty"`
	CustomField26     string        `json:"custom_field_26,omitempty"`
	CustomField27     string        `json:"custom_field_27,omitempty"`
	CustomField28     string        `json:"custom_field_28,omitempty"`
	CustomField29     string        `json:"custom_field_29,omitempty"`
	CustomField30     string        `json:"custom_field_30,omitempty"`
	CustomField31     string        `json:"custom_field_31,omitempty"`
	CustomField32     string        `json:"custom_field_32,omitempty"`
	CustomField33     string        `json:"custom_field_33,omitempty"`
	CustomField34     string        `json:"custom_field_34,omitempty"`
	CustomField35     string        `json:"custom_field_35,omitempty"`
	CustomField36     string        `json:"custom_field_36,omitempty"`
	CustomField37     string        `json:"custom_field_37,omitempty"`
	CustomField38     string        `json:"custom_field_38,omitempty"`
	CustomField39     string        `json:"custom_field_39,omitempty"`
	CustomField40     string        `json:"custom_field_40,omitempty"`
	CustomField41     string        `json:"custom_field_41,omitempty"`
	CustomField42     string        `json:"custom_field_42,omitempty"`
	CustomField43     string        `json:"custom_field_43,omitempty"`
	CustomField44     string        `json:"custom_field_44,omitempty"`
	CustomField45     string        `json:"custom_field_45,omitempty"`
	CustomField46     string        `json:"custom_field_46,omitempty"`
	CustomField47     string        `json:"custom_field_47,omitempty"`
	CustomField48     string        `json:"custom_field_48,omitempty"`
	CustomField49     string        `json:"custom_field_49,omitempty"`
	CustomField50     string        `json:"custom_field_50,omitempty"`
	CustomPlanField1  string        `json:"custom_plan_field_1,omitempty"`
	CustomPlanField2  string        `json:"custom_plan_field_2,omitempty"`
	CustomPlanField3  string        `json:"custom_plan_field_3,omitempty"`
	CustomPlanField4  string        `json:"custom_plan_field_4,omitempty"`
	CustomPlanField5  string        `json:"custom_plan_field_5,omitempty"`
	CustomPlanField6  string        `json:"custom_plan_field_6,omitempty"`
	CustomPlanField7  string        `json:"custom_plan_field_7,omitempty"`
	CustomPlanField8  string        `json:"custom_plan_field_8,omitempty"`
	CustomPlanField9  string        `json:"custom_plan_field_9,omitempty"`
	CustomPlanField10 string        `json:"custom_plan_field_10,omitempty"`
	PriorityLabel     PriorityLabel `json:"priority_label,omitempty"` // 优先级
}

// TaskService 任务服务
type TaskService struct {
	client *Client
}

// NewTaskService 创建任务服务
func NewTaskService(client *Client) *TaskService {
	return &TaskService{client}
}

// 创建任务
// 获取任务变更历史
// 获取任务变更次数
// 获取任务自定义字段配置

type GetTasksRequest struct {
	ID               *MultiType[int]       `url:"id,omitempty"`               // 支持多ID查询、模糊匹配
	Name             *string               `url:"name,omitempty"`             // 任务标题	支持模糊匹配
	Description      *string               `url:"description,omitempty"`      // 任务详细描述
	WorkspaceID      *int                  `url:"workspace_id,omitempty"`     // [必须]项目ID
	Creator          *string               `url:"creator,omitempty"`          // 创建人	支持多人员查询
	Created          *string               `url:"created,omitempty"`          // 创建时间	支持时间查询
	Modified         *string               `url:"modified,omitempty"`         // 最后修改时间	支持时间查询
	Status           *EnumType[TaskStatus] `url:"status,omitempty"`           // 状态	支持枚举查询
	Label            *EnumType[string]     `url:"label,omitempty"`            // 标签查询	支持枚举查询
	Owner            *string               `url:"owner,omitempty"`            // 任务当前处理人	支持模糊匹配
	CC               *string               `url:"cc,omitempty"`               // 抄送人
	Begin            *string               `url:"begin,omitempty"`            // 预计开始	支持时间查询
	Due              *string               `url:"due,omitempty"`              // 预计结束	支持时间查询
	StoryID          *MultiType[int]       `url:"story_id,omitempty"`         // 关联需求的ID	支持多ID查询
	IterationID      *EnumType[int]        `url:"iteration_id,omitempty"`     // 所属迭代的ID	支持枚举查询
	Priority         *string               `url:"priority,omitempty"`         //nolint:lll // 优先级。为了兼容自定义优先级，请使用 priority_label 字段，详情参考：如何兼容自定义优先级
	PriorityLabel    *PriorityLabel        `url:"priority_label,omitempty"`   // 优先级。推荐使用这个字段
	Progress         *int                  `url:"progress,omitempty"`         // 进度
	Completed        *string               `url:"completed,omitempty"`        // 完成时间	支持时间查询
	EffortCompleted  *string               `url:"effort_completed,omitempty"` // 完成工时
	Exceed           *float64              `url:"exceed,omitempty"`           // 超出工时
	Remain           *float64              `url:"remain,omitempty"`           // 剩余工时
	Effort           *string               `url:"effort,omitempty"`           // 预估工时
	CustomFieldOne   *string               `url:"custom_field_one,omitempty"`
	CustomFieldTwo   *string               `url:"custom_field_two,omitempty"`
	CustomFieldThree *string               `url:"custom_field_three,omitempty"`
	CustomFieldFour  *string               `url:"custom_field_four,omitempty"`
	CustomFieldFive  *string               `url:"custom_field_five,omitempty"`
	CustomFieldSix   *string               `url:"custom_field_six,omitempty"`
	CustomFieldSeven *string               `url:"custom_field_seven,omitempty"`
	CustomFieldEight *string               `url:"custom_field_eight,omitempty"`
	CustomField9     *string               `url:"custom_field_9,omitempty"`
	CustomField10    *string               `url:"custom_field_10,omitempty"`
	CustomField11    *string               `url:"custom_field_11,omitempty"`
	CustomField12    *string               `url:"custom_field_12,omitempty"`
	CustomField13    *string               `url:"custom_field_13,omitempty"`
	CustomField14    *string               `url:"custom_field_14,omitempty"`
	CustomField15    *string               `url:"custom_field_15,omitempty"`
	CustomField16    *string               `url:"custom_field_16,omitempty"`
	CustomField17    *string               `url:"custom_field_17,omitempty"`
	CustomField18    *string               `url:"custom_field_18,omitempty"`
	CustomField19    *string               `url:"custom_field_19,omitempty"`
	CustomField20    *string               `url:"custom_field_20,omitempty"`
	CustomField21    *string               `url:"custom_field_21,omitempty"`
	CustomField22    *string               `url:"custom_field_22,omitempty"`
	CustomField23    *string               `url:"custom_field_23,omitempty"`
	CustomField24    *string               `url:"custom_field_24,omitempty"`
	CustomField25    *string               `url:"custom_field_25,omitempty"`
	CustomField26    *string               `url:"custom_field_26,omitempty"`
	CustomField27    *string               `url:"custom_field_27,omitempty"`
	CustomField28    *string               `url:"custom_field_28,omitempty"`
	CustomField29    *string               `url:"custom_field_29,omitempty"`
	CustomField30    *string               `url:"custom_field_30,omitempty"`
	CustomField31    *string               `url:"custom_field_31,omitempty"`
	CustomField32    *string               `url:"custom_field_32,omitempty"`
	CustomField33    *string               `url:"custom_field_33,omitempty"`
	CustomField34    *string               `url:"custom_field_34,omitempty"`
	CustomField35    *string               `url:"custom_field_35,omitempty"`
	CustomField36    *string               `url:"custom_field_36,omitempty"`
	CustomField37    *string               `url:"custom_field_37,omitempty"`
	CustomField38    *string               `url:"custom_field_38,omitempty"`
	CustomField39    *string               `url:"custom_field_39,omitempty"`
	CustomField40    *string               `url:"custom_field_40,omitempty"`
	CustomField41    *string               `url:"custom_field_41,omitempty"`
	CustomField42    *string               `url:"custom_field_42,omitempty"`
	CustomField43    *string               `url:"custom_field_43,omitempty"`
	CustomField44    *string               `url:"custom_field_44,omitempty"`
	CustomField45    *string               `url:"custom_field_45,omitempty"`
	CustomField46    *string               `url:"custom_field_46,omitempty"`
	CustomField47    *string               `url:"custom_field_47,omitempty"`
	CustomField48    *string               `url:"custom_field_48,omitempty"`
	CustomField49    *string               `url:"custom_field_49,omitempty"`
	CustomField50    *string               `url:"custom_field_50,omitempty"`
	Limit            *int                  `url:"limit,omitempty"`  // 设置返回数量限制，默认为30
	Page             *int                  `url:"page,omitempty"`   // 返回当前数量限制下第N页的数据，默认为1（第一页）
	Order            *Order                `url:"order,omitempty"`  //nolint:lll // 排序规则，规则：字段名 ASC或者DESC，然后 urlencode	如按创建时间逆序：order=created%20desc
	Fields           *MultiType[string]    `url:"fields,omitempty"` // 设置获取的字段，多个字段间以','逗号隔开
}

// GetTasks 获取任务
//
// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/task/get_tasks.html
func (s *TaskService) GetTasks(
	ctx context.Context, request *GetTasksRequest, opts ...RequestOption,
) ([]*Task, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "tasks", request, opts)
	if err != nil {
		return nil, nil, err
	}

	var items []struct {
		Task *Task `json:"Task"`
	}
	resp, err := s.client.Do(req, &items)
	if err != nil {
		return nil, resp, err
	}

	tasks := make([]*Task, 0, len(items))
	for _, item := range items {
		tasks = append(tasks, item.Task)
	}

	return tasks, resp, nil
}

type GetTasksCountRequest struct {
	ID               *MultiType[int]       `url:"id,omitempty"`               // 支持多ID查询、模糊匹配
	Name             *string               `url:"name,omitempty"`             // 任务标题	支持模糊匹配
	Description      *string               `url:"description,omitempty"`      // 任务详细描述
	WorkspaceID      *int                  `url:"workspace_id,omitempty"`     // [必须]项目ID
	Creator          *string               `url:"creator,omitempty"`          // 创建人	支持多人员查询
	Created          *string               `url:"created,omitempty"`          // 创建时间	支持时间查询
	Modified         *string               `url:"modified,omitempty"`         // 最后修改时间	支持时间查询
	Status           *EnumType[TaskStatus] `url:"status,omitempty"`           // 状态	支持枚举查询
	Label            *EnumType[string]     `url:"label,omitempty"`            // 标签查询	支持枚举查询
	Owner            *string               `url:"owner,omitempty"`            // 任务当前处理人	支持模糊匹配
	CC               *string               `url:"cc,omitempty"`               // 抄送人
	Begin            *string               `url:"begin,omitempty"`            // 预计开始	支持时间查询
	Due              *string               `url:"due,omitempty"`              // 预计结束	支持时间查询
	StoryID          *MultiType[int]       `url:"story_id,omitempty"`         // 关联需求的ID	支持多ID查询
	IterationID      *EnumType[int]        `url:"iteration_id,omitempty"`     // 所属迭代的ID	支持枚举查询
	Priority         *string               `url:"priority,omitempty"`         //nolint:lll // 优先级。为了兼容自定义优先级，请使用 priority_label 字段，详情参考：如何兼容自定义优先级
	PriorityLabel    *PriorityLabel        `url:"priority_label,omitempty"`   // 优先级。推荐使用这个字段
	Progress         *int                  `url:"progress,omitempty"`         // 进度
	Completed        *string               `url:"completed,omitempty"`        // 完成时间	支持时间查询
	EffortCompleted  *string               `url:"effort_completed,omitempty"` // 完成工时
	Exceed           *float64              `url:"exceed,omitempty"`           // 超出工时
	Remain           *float64              `url:"remain,omitempty"`           // 剩余工时
	Effort           *string               `url:"effort,omitempty"`           // 预估工时
	CustomFieldOne   *string               `url:"custom_field_one,omitempty"`
	CustomFieldTwo   *string               `url:"custom_field_two,omitempty"`
	CustomFieldThree *string               `url:"custom_field_three,omitempty"`
	CustomFieldFour  *string               `url:"custom_field_four,omitempty"`
	CustomFieldFive  *string               `url:"custom_field_five,omitempty"`
	CustomFieldSix   *string               `url:"custom_field_six,omitempty"`
	CustomFieldSeven *string               `url:"custom_field_seven,omitempty"`
	CustomFieldEight *string               `url:"custom_field_eight,omitempty"`
	CustomField9     *string               `url:"custom_field_9,omitempty"`
	CustomField10    *string               `url:"custom_field_10,omitempty"`
	CustomField11    *string               `url:"custom_field_11,omitempty"`
	CustomField12    *string               `url:"custom_field_12,omitempty"`
	CustomField13    *string               `url:"custom_field_13,omitempty"`
	CustomField14    *string               `url:"custom_field_14,omitempty"`
	CustomField15    *string               `url:"custom_field_15,omitempty"`
	CustomField16    *string               `url:"custom_field_16,omitempty"`
	CustomField17    *string               `url:"custom_field_17,omitempty"`
	CustomField18    *string               `url:"custom_field_18,omitempty"`
	CustomField19    *string               `url:"custom_field_19,omitempty"`
	CustomField20    *string               `url:"custom_field_20,omitempty"`
	CustomField21    *string               `url:"custom_field_21,omitempty"`
	CustomField22    *string               `url:"custom_field_22,omitempty"`
	CustomField23    *string               `url:"custom_field_23,omitempty"`
	CustomField24    *string               `url:"custom_field_24,omitempty"`
	CustomField25    *string               `url:"custom_field_25,omitempty"`
	CustomField26    *string               `url:"custom_field_26,omitempty"`
	CustomField27    *string               `url:"custom_field_27,omitempty"`
	CustomField28    *string               `url:"custom_field_28,omitempty"`
	CustomField29    *string               `url:"custom_field_29,omitempty"`
	CustomField30    *string               `url:"custom_field_30,omitempty"`
	CustomField31    *string               `url:"custom_field_31,omitempty"`
	CustomField32    *string               `url:"custom_field_32,omitempty"`
	CustomField33    *string               `url:"custom_field_33,omitempty"`
	CustomField34    *string               `url:"custom_field_34,omitempty"`
	CustomField35    *string               `url:"custom_field_35,omitempty"`
	CustomField36    *string               `url:"custom_field_36,omitempty"`
	CustomField37    *string               `url:"custom_field_37,omitempty"`
	CustomField38    *string               `url:"custom_field_38,omitempty"`
	CustomField39    *string               `url:"custom_field_39,omitempty"`
	CustomField40    *string               `url:"custom_field_40,omitempty"`
	CustomField41    *string               `url:"custom_field_41,omitempty"`
	CustomField42    *string               `url:"custom_field_42,omitempty"`
	CustomField43    *string               `url:"custom_field_43,omitempty"`
	CustomField44    *string               `url:"custom_field_44,omitempty"`
	CustomField45    *string               `url:"custom_field_45,omitempty"`
	CustomField46    *string               `url:"custom_field_46,omitempty"`
	CustomField47    *string               `url:"custom_field_47,omitempty"`
	CustomField48    *string               `url:"custom_field_48,omitempty"`
	CustomField49    *string               `url:"custom_field_49,omitempty"`
	CustomField50    *string               `url:"custom_field_50,omitempty"`
}

// GetTasksCount 获取任务数量
//
// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/task/get_tasks_count.html
func (s *TaskService) GetTasksCount(
	ctx context.Context, request *GetTasksCountRequest, opts ...RequestOption,
) (int, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "tasks/count", request, opts)
	if err != nil {
		return 0, nil, err
	}

	var response CountResponse
	resp, err := s.client.Do(req, &response)
	if err != nil {
		return 0, resp, err
	}

	return response.Count, resp, nil
}

// 更新任务
// 获取回收站下的任务
// 获取视图对应的任务列表

// -----------------------------------------------------------------------------
// GetTaskFieldsInfo 获取任务字段信息
// -----------------------------------------------------------------------------

type GetTaskFieldsInfoRequest struct {
	WorkspaceID *int `url:"workspace_id,omitempty"` // [必须]项目ID
}

type TaskFieldsInfoHTMLType string

const (
	TaskFieldsInfoHTMLTypeInput       TaskFieldsInfoHTMLType = "input"
	TaskFieldsInfoHTMLTypeSelect      TaskFieldsInfoHTMLType = "select"
	TaskFieldsInfoHTMLTypeRichEdit    TaskFieldsInfoHTMLType = "rich_edit"
	TaskFieldsInfoHTMLTypeUserChooser TaskFieldsInfoHTMLType = "user_chooser"
	TaskFieldsInfoHTMLTypeDatetime    TaskFieldsInfoHTMLType = "datetime"
	TaskFieldsInfoHTMLTypeFloat       TaskFieldsInfoHTMLType = "float"
	TaskFieldsInfoHTMLTypeMixChooser  TaskFieldsInfoHTMLType = "mix_chooser"
	TaskFieldsInfoHTMLTypeDateInput   TaskFieldsInfoHTMLType = "dateinput"
	TaskFieldsInfoHTMLTypeCheckbox    TaskFieldsInfoHTMLType = "checkbox"
	TaskFieldsInfoHTMLTypeMultiSelect TaskFieldsInfoHTMLType = "multi_select"
)

type TaskFieldsInfoOption struct {
	Value string `json:"key,omitempty"`   // 值
	Label string `json:"label,omitempty"` // 中文名称
}

type TaskFieldsInfoColorOption struct {
	Value string `json:"value,omitempty"`
	Label string `json:"label,omitempty"`
	Color string `json:"color,omitempty"`
}

type TaskFieldsInfoPureOption struct {
	ParentID    string `json:"parent_id,omitempty"`
	WorkspaceID string `json:"workspace_id,omitempty"`
	Sort        string `json:"sort,omitempty"`
	OriginName  string `json:"origin_name,omitempty"`
	Value       string `json:"value,omitempty"`
	Label       string `json:"label,omitempty"`
	Panel       int    `json:"panel,omitempty"`
}

type TaskFieldsInfo struct {
	Name         string                      `json:"name,omitempty"`      // name
	HTMLType     TaskFieldsInfoHTMLType      `json:"html_type,omitempty"` // 类型
	Label        string                      `json:"label,omitempty"`     // 中文名称
	Options      []TaskFieldsInfoOption      `json:"options,omitempty"`   // 候选值
	ColorOptions []TaskFieldsInfoColorOption `json:"color_options,omitempty"`
	PureOptions  []TaskFieldsInfoPureOption  `json:"pure_options,omitempty"`
}

type getTaskFieldsInfoResponse map[string]struct {
	Name         string                      `json:"name,omitempty"`      // name
	HTMLType     TaskFieldsInfoHTMLType      `json:"html_type,omitempty"` // 类型
	Label        string                      `json:"label,omitempty"`     // 中文名称
	Options      any                         `json:"options,omitempty"`   // 候选值
	ColorOptions []TaskFieldsInfoColorOption `json:"color_options,omitempty"`
	PureOptions  []TaskFieldsInfoPureOption  `json:"pure_options,omitempty"`
}

// GetTaskFieldsInfo 获取任务字段信息
//
// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/task/get_task_fields_info.html
func (s *TaskService) GetTaskFieldsInfo(
	ctx context.Context, request *GetTaskFieldsInfoRequest, opts ...RequestOption,
) ([]*TaskFieldsInfo, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "tasks/get_fields_info", request, opts)
	if err != nil {
		return nil, nil, err
	}

	var response getTaskFieldsInfoResponse
	resp, err := s.client.Do(req, &response)
	if err != nil {
		return nil, resp, err
	}

	fields := make([]*TaskFieldsInfo, 0, len(response))
	for _, item := range response {
		options := make([]TaskFieldsInfoOption, 0)

		if item.Options != nil {
			if os, ok := item.Options.(map[string]any); ok {
				options = make([]TaskFieldsInfoOption, 0, len(os))
				for key, value := range os {
					if v, ok2 := value.(string); ok2 {
						options = append(options, TaskFieldsInfoOption{
							Value: key,
							Label: v,
						})
					}
				}
			}
		}

		fields = append(fields, &TaskFieldsInfo{
			Name:         item.Name,
			HTMLType:     item.HTMLType,
			Label:        item.Label,
			Options:      options,
			ColorOptions: item.ColorOptions,
			PureOptions:  item.PureOptions,
		})
	}

	return fields, resp, nil
}
