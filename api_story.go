package tapd

import (
	"context"
	"encoding/json"
	"net/http"
)

type StoryService struct {
	client *Client
}

func NewStoryService(client *Client) *StoryService {
	return &StoryService{
		client: client,
	}
}

type CreateStoryRequest struct {
	WorkspaceID     *int           `json:"workspace_id,omitempty"`     // [必须]项目ID
	Name            *string        `json:"name,omitempty"`             // [必须]标题
	Priority        *string        `json:"priority,omitempty"`         // 优先级
	PriorityLabel   *PriorityLabel `json:"priority_label,omitempty"`   // 优先级。推荐使用这个字段
	BusinessValue   *int           `json:"business_value,omitempty"`   // 业务价值
	Version         *string        `json:"version,omitempty"`          // 版本
	Module          *string        `json:"module,omitempty"`           // 模块
	TestFocus       *string        `json:"test_focus,omitempty"`       // 测试重点
	Size            *int           `json:"size,omitempty"`             // 规模
	Owner           *string        `json:"owner,omitempty"`            // 处理人
	CC              *string        `json:"cc,omitempty"`               // 抄送人
	Creator         *string        `json:"creator,omitempty"`          // 创建人
	Developer       *string        `json:"developer,omitempty"`        // 开发人员
	Begin           *string        `json:"begin,omitempty"`            // 预计开始
	Due             *string        `json:"due,omitempty"`              // 预计结束
	IterationID     *string        `json:"iteration_id,omitempty"`     // 迭代ID
	TemplatedID     *int           `json:"templated_id,omitempty"`     // 模板ID
	ParentID        *int           `json:"parent_id,omitempty"`        // 父需求ID
	Effort          *string        `json:"effort,omitempty"`           // 预估工时
	EffortCompleted *string        `json:"effort_completed,omitempty"` // 完成工时
	Remain          *float64       `json:"remain,omitempty"`           // 剩余工时
	Exceed          *float64       `json:"exceed,omitempty"`           // 超出工时
	CategoryID      *int           `json:"category_id,omitempty"`      // 需求分类
	WorkitemTypeID  *int           `json:"workitem_type_id,omitempty"` // 需求类别
	ReleaseID       *int           `json:"release_id,omitempty"`       // 发布计划
	Source          *string        `json:"source,omitempty"`           // 来源
	Type            *string        `json:"type,omitempty"`             // 类型
	Description     *string        `json:"description,omitempty"`      // 详细描述
	Label           *string        `json:"label,omitempty"`            // 标签，标签不存在时将自动创建，多个以英文坚线分格
}

// CreateStory 创建需求
// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/story/add_story.html
func (s *StoryService) CreateStory(
	ctx context.Context, request *CreateStoryRequest, opts ...RequestOption,
) (*Story, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodPost, "stories", request, opts)
	if err != nil {
		return nil, nil, err
	}

	var response struct {
		Story *Story `json:"story"`
	}
	resp, err := s.client.Do(req, &response)
	if err != nil {
		return nil, resp, err
	}

	return response.Story, resp, nil
}

// 创建需求分类
// 复制需求
// 获取需求与其它需求的所有关联关系

type GetStoriesRequest struct {
	CustomFieldsRequest
	CustomPlanFieldsRequest

	ID              *MultiType[int] `url:"id,omitempty"`               // ID	支持多ID查询,多个ID用逗号分隔
	Name            *string         `url:"name,omitempty"`             // 标题	支持模糊匹配
	Priority        *string         `url:"priority,omitempty"`         // 优先级
	PriorityLabel   *PriorityLabel  `url:"priority_label,omitempty"`   // 优先级。推荐使用这个字段
	BusinessValue   *int            `url:"business_value,omitempty"`   // 业务价值
	Status          *string         `url:"status,omitempty"`           // 状态	支持枚举查询
	VStatus         *string         `url:"v_status,omitempty"`         // 状态(支持传入中文状态名称)
	WithVStatus     *string         `url:"with_v_status,omitempty"`    // 值=1可以返回中文状态
	Label           *string         `url:"label,omitempty"`            // 标签查询	支持枚举查询
	WorkitemTypeID  *string         `url:"workitem_type_id,omitempty"` // 需求类别ID	支持枚举查询
	Version         *string         `url:"version,omitempty"`          // 版本
	Module          *string         `url:"module,omitempty"`           // 模块
	Feature         *string         `url:"feature,omitempty"`          // 特性
	TestFocus       *string         `url:"test_focus,omitempty"`       // 测试重点
	Size            *int            `url:"size,omitempty"`             // 规模
	Owner           *string         `url:"owner,omitempty"`            // 处理人	支持模糊匹配
	CC              *string         `url:"cc,omitempty"`               // 抄送人	支持模糊匹配
	Creator         *string         `url:"creator,omitempty"`          // 创建人	支持多人员查询
	Developer       *string         `url:"developer,omitempty"`        // 开发人员
	Begin           *string         `url:"begin,omitempty"`            // 预计开始	支持时间查询
	Due             *string         `url:"due,omitempty"`              // 预计结束	支持时间查询
	Created         *string         `url:"created,omitempty"`          // 创建时间	支持时间查询
	Modified        *string         `url:"modified,omitempty"`         // 最后修改时间	支持时间查询
	Completed       *string         `url:"completed,omitempty"`        // 完成时间	支持时间查询
	IterationID     *string         `url:"iteration_id,omitempty"`     // 迭代ID	支持不等于查询
	Effort          *string         `url:"effort,omitempty"`           // 预估工时
	EffortCompleted *string         `url:"effort_completed,omitempty"` // 完成工时
	Remain          *float64        `url:"remain,omitempty"`           // 剩余工时
	Exceed          *float64        `url:"exceed,omitempty"`           // 超出工时
	CategoryID      *string         `url:"category_id,omitempty"`      // 需求分类	支持枚举查询
	ReleaseID       *string         `url:"release_id,omitempty"`       // 发布计划
	Source          *string         `url:"source,omitempty"`           // 需求来源
	Type            *string         `url:"type,omitempty"`             // 需求类型
	ParentID        *string         `url:"parent_id,omitempty"`        // 父需求
	ChildrenID      *string         `url:"children_id,omitempty"`      // 子需求	为空查询传：丨
	Description     *string         `url:"description,omitempty"`      // 详细描述	支持模糊匹配
	WorkspaceID     *int            `url:"workspace_id,omitempty"`     // 项目ID
	Limit           *int            `url:"limit,omitempty"`            // 设置返回数量限制，默认为30
	Page            *int            `url:"page,omitempty"`             // 返回当前数量限制下第N页的数据，默认为1（第一页）
	Order           *Order          `url:"order,omitempty"`            // 排序规则，规则：字段名 ASC或者DESC
	Fields          *string         `url:"fields,omitempty"`           // 设置获取的字段，多个字段间以','逗号隔开
}

type Story struct {
	CustomFields
	CustomPlanFields

	ID              string        `json:"id,omitempty"`
	WorkitemTypeID  string        `json:"workitem_type_id,omitempty"`
	Name            string        `json:"name,omitempty"`
	Description     string        `json:"description,omitempty"`
	WorkspaceID     string        `json:"workspace_id,omitempty"`
	Creator         string        `json:"creator,omitempty"`
	Created         string        `json:"created,omitempty"`
	Modified        string        `json:"modified,omitempty"`
	Status          string        `json:"status,omitempty"`
	Step            string        `json:"step,omitempty"`
	Owner           string        `json:"owner,omitempty"`
	Cc              string        `json:"cc,omitempty"`
	Begin           *string       `json:"begin,omitempty"`
	Due             *string       `json:"due,omitempty"`
	Size            *string       `json:"size,omitempty"`
	Priority        string        `json:"priority,omitempty"`
	Developer       string        `json:"developer,omitempty"`
	IterationID     string        `json:"iteration_id,omitempty"`
	TestFocus       string        `json:"test_focus,omitempty"`
	Type            string        `json:"type,omitempty"`
	Source          string        `json:"source,omitempty"`
	Module          string        `json:"module,omitempty"`
	Version         string        `json:"version,omitempty"`
	Completed       *string       `json:"completed,omitempty"`
	CategoryID      string        `json:"category_id,omitempty"`
	Path            string        `json:"path,omitempty"`
	ParentID        string        `json:"parent_id,omitempty"`
	ChildrenID      string        `json:"children_id,omitempty"`
	AncestorID      string        `json:"ancestor_id,omitempty"`
	Level           string        `json:"level,omitempty"`
	BusinessValue   *string       `json:"business_value,omitempty"`
	Effort          *string       `json:"effort,omitempty"`
	EffortCompleted string        `json:"effort_completed,omitempty"`
	Exceed          string        `json:"exceed,omitempty"`
	Remain          string        `json:"remain,omitempty"`
	ReleaseID       string        `json:"release_id,omitempty"`
	BugID           string        `json:"bug_id,omitempty"`
	TemplatedID     string        `json:"templated_id,omitempty"`
	CreatedFrom     string        `json:"created_from,omitempty"`
	Feature         string        `json:"feature,omitempty"`
	Label           string        `json:"label,omitempty"`
	Progress        string        `json:"progress,omitempty"`
	IsArchived      string        `json:"is_archived,omitempty"`
	TechRisk        *string       `json:"tech_risk,omitempty"`
	Flows           *string       `json:"flows,omitempty"`
	SecretRootID    string        `json:"secret_root_id,omitempty"`
	PriorityLabel   PriorityLabel `json:"priority_label,omitempty"`
}

// GetStories 获取需求
// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/story/get_stories.html
func (s *StoryService) GetStories(
	ctx context.Context, request *GetStoriesRequest, opts ...RequestOption,
) ([]*Story, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "stories", request, opts)
	if err != nil {
		return nil, nil, err
	}

	var items []struct {
		Story *Story `json:"story"`
	}
	resp, err := s.client.Do(req, &items)
	if err != nil {
		return nil, resp, err
	}

	stories := make([]*Story, 0, len(items))
	for _, item := range items {
		stories = append(stories, item.Story)
	}

	return stories, resp, nil
}

type GetStoriesCountRequest struct {
	CustomFieldsRequest
	CustomPlanFieldsRequest

	ID              *MultiType[int] `url:"id,omitempty"`               // ID	支持多ID查询,多个ID用逗号分隔
	Name            *string         `url:"name,omitempty"`             // 标题	支持模糊匹配
	Priority        *string         `url:"priority,omitempty"`         // 优先级。
	PriorityLabel   *PriorityLabel  `url:"priority_label,omitempty"`   // 优先级。推荐使用这个字段
	BusinessValue   *int            `url:"business_value,omitempty"`   // 业务价值
	Status          *string         `url:"status,omitempty"`           // 状态	支持枚举查询
	VStatus         *string         `url:"v_status,omitempty"`         // 状态(支持传入中文状态名称)
	WithVStatus     *string         `url:"with_v_status,omitempty"`    // 值=1可以返回中文状态
	Label           *string         `url:"label,omitempty"`            // 标签查询	支持枚举查询
	WorkitemTypeID  *string         `url:"workitem_type_id,omitempty"` // 需求类别ID	支持枚举查询
	Version         *string         `url:"version,omitempty"`          // 版本
	Module          *string         `url:"module,omitempty"`           // 模块
	Feature         *string         `url:"feature,omitempty"`          // 特性
	TestFocus       *string         `url:"test_focus,omitempty"`       // 测试重点
	Size            *int            `url:"size,omitempty"`             // 规模
	Owner           *string         `url:"owner,omitempty"`            // 处理人	支持模糊匹配
	CC              *string         `url:"cc,omitempty"`               // 抄送人	支持模糊匹配
	Creator         *string         `url:"creator,omitempty"`          // 创建人	支持多人员查询
	Developer       *string         `url:"developer,omitempty"`        // 开发人员
	Begin           *string         `url:"begin,omitempty"`            // 预计开始	支持时间查询
	Due             *string         `url:"due,omitempty"`              // 预计结束	支持时间查询
	Created         *string         `url:"created,omitempty"`          // 创建时间	支持时间查询
	Modified        *string         `url:"modified,omitempty"`         // 最后修改时间	支持时间查询
	Completed       *string         `url:"completed,omitempty"`        // 完成时间	支持时间查询
	IterationID     *string         `url:"iteration_id,omitempty"`     // 迭代ID	支持不等于查询
	Effort          *string         `url:"effort,omitempty"`           // 预估工时
	EffortCompleted *string         `url:"effort_completed,omitempty"` // 完成工时
	Remain          *float64        `url:"remain,omitempty"`           // 剩余工时
	Exceed          *float64        `url:"exceed,omitempty"`           // 超出工时
	CategoryID      *string         `url:"category_id,omitempty"`      // 需求分类	支持枚举查询
	ReleaseID       *string         `url:"release_id,omitempty"`       // 发布计划
	Source          *string         `url:"source,omitempty"`           // 需求来源
	Type            *string         `url:"type,omitempty"`             // 需求类型
	ParentID        *string         `url:"parent_id,omitempty"`        // 父需求
	ChildrenID      *string         `url:"children_id,omitempty"`      // 子需求	为空查询传：丨
	Description     *string         `url:"description,omitempty"`      // 详细描述	支持模糊匹配
	WorkspaceID     *int            `url:"workspace_id,omitempty"`     // 项目ID
}

// GetStoriesCount 获取需求数量
// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/story/get_stories_count.html
func (s *StoryService) GetStoriesCount(
	ctx context.Context, request *GetStoriesCountRequest, opts ...RequestOption,
) (int, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "stories/count", request, opts)
	if err != nil {
		return 0, nil, err
	}

	var response struct {
		Count int `json:"count,omitempty"`
	}
	resp, err := s.client.Do(req, &response)
	if err != nil {
		return 0, resp, err
	}

	return response.Count, resp, nil
}

// 获取保密需求
// 获取保密需求数量

type GetStoryCategoriesRequest struct {
	WorkspaceID *int            `url:"workspace_id,omitempty"` // [必须]项目ID
	ID          *MultiType[int] `url:"id,omitempty"`           // ID 支持多ID查询，多个ID用逗号分隔
	Name        *string         `url:"name,omitempty"`         // 需求分类名称	支持模糊匹配
	Description *string         `url:"description,omitempty"`  // 需求分类描述
	ParentID    *int            `url:"parent_id,omitempty"`    // 父分类ID
	Created     *string         `url:"created,omitempty"`      // 创建时间	支持时间查询
	Modified    *string         `url:"modified,omitempty"`     // 最后修改时间	支持时间查询
	Limit       *int            `url:"limit,omitempty"`        // 设置返回数量限制，默认为30
	Page        *int            `url:"page,omitempty"`         // 返回当前数量限制下第N页的数据，默认为1（第一页）
	Order       *Order          `url:"order,omitempty"`        //nolint:lll // 排序规则，规则：字段名 ASC或者DESC，然后 urlencode	如按创建时间逆序：order=created%20desc
	Fields      *string         `url:"fields,omitempty"`       // 设置获取的字段，多个字段间以','逗号隔开
}

type StoryCategory struct {
	ID          *string `json:"id"`
	WorkspaceID *string `json:"workspace_id"`
	Name        *string `json:"name"`
	Description *string `json:"description"`
	ParentID    *string `json:"parent_id"`
	Created     *string `json:"created"`
	Modified    *string `json:"modified"`
}

// GetStoryCategories 获取需求分类
// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/story/get_story_categories.html
func (s *StoryService) GetStoryCategories(
	ctx context.Context, request *GetStoryCategoriesRequest, opts ...RequestOption,
) ([]*StoryCategory, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "story_categories", request, opts)
	if err != nil {
		return nil, nil, err
	}

	var items []struct {
		Category *StoryCategory
	}
	resp, err := s.client.Do(req, &items)
	if err != nil {
		return nil, resp, err
	}

	categories := make([]*StoryCategory, 0, len(items))
	for _, item := range items {
		categories = append(categories, item.Category)
	}

	return categories, resp, nil
}

type GetStoryCategoriesCountRequest struct {
	WorkspaceID *int            `url:"workspace_id,omitempty"` // [必须]项目ID
	ID          *MultiType[int] `url:"id,omitempty"`           // ID 支持多ID查询，多个ID用逗号分隔
	Name        *string         `url:"name,omitempty"`         // 需求分类名称	支持模糊匹配
	Description *string         `url:"description,omitempty"`  // 需求分类描述
	ParentID    *int            `url:"parent_id,omitempty"`    // 父分类ID
	Created     *string         `url:"created,omitempty"`      // 创建时间	支持时间查询
	Modified    *string         `url:"modified,omitempty"`     // 最后修改时间	支持时间查询
}

// GetStoryCategoriesCount 获取需求分类数量
// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/story/get_story_categories_count.html
func (s *StoryService) GetStoryCategoriesCount(
	ctx context.Context, request *GetStoryCategoriesCountRequest, opts ...RequestOption,
) (int, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "story_categories/count", request, opts)
	if err != nil {
		return 0, nil, err
	}

	response := new(CountResponse)
	resp, err := s.client.Do(req, &response)
	if err != nil {
		return 0, resp, err
	}

	return response.Count, resp, nil
}

// -----------------------------------------------------------------------------
// 获取指定分类需求数量
// -----------------------------------------------------------------------------

// -----------------------------------------------------------------------------
// 获取需求变更历史
// -----------------------------------------------------------------------------

type GetStoryChangesRequest struct {
	ID               *MultiType[int] `url:"id,omitempty"`
	StoryID          *MultiType[int] `url:"story_id,omitempty"`           // 需求id	支持多ID查询
	WorkspaceID      *int            `url:"workspace_id,omitempty"`       // [必须]项目ID
	Creator          *string         `url:"creator,omitempty"`            // 创建人（操作人）
	Created          *string         `url:"created,omitempty"`            // 创建时间（变更时间）	支持时间查询
	ChangeType       *string         `url:"change_type,omitempty"`        // 变更类型	值范围见文档下方附录1
	ChangeSummary    *string         `url:"change_summary,omitempty"`     // 需求变更描述
	Comment          *string         `url:"comment,omitempty"`            // 评论
	EntityType       *string         `url:"entity_type,omitempty"`        // 变更的对象类型
	ChangeField      *string         `url:"change_field,omitempty"`       // 设置获取变更字段如（status）
	NeedParseChanges *int            `url:"need_parse_changes,omitempty"` // 设置field_changes字段是否返回（默认取 1。取 0 则不返回）
	Limit            *int            `url:"limit,omitempty"`              // 设置返回数量限制，默认为30，最大取 100
	Page             *int            `url:"page,omitempty"`               // 返回当前数量限制下第N页的数据，默认为1（第一页）
	Order            *Order          `url:"order,omitempty"`              // 排序规则，规则：字段名 ASC或者DESC
	Fields           *string         `url:"fields,omitempty"`             // 设置获取的字段，多个字段间以','逗号隔开
}

type StoryChange struct {
	ID             string  `json:"id,omitempty"`
	WorkspaceID    string  `json:"workspace_id,omitempty"`
	AppID          string  `json:"app_id,omitempty"`
	WorkitemTypeID string  `json:"workitem_type_id,omitempty"`
	Creator        string  `json:"creator,omitempty"`
	Created        string  `json:"created,omitempty"`
	ChangeSummary  *string `json:"change_summary,omitempty"`
	Comment        *string `json:"comment,omitempty"`
	Changes        string  `json:"changes,omitempty"`
	EntityType     string  `json:"entity_type,omitempty"`
	ChangeType     string  `json:"change_type,omitempty"`
	ChangeTypeText string  `json:"change_type_text,omitempty"`
	FieldChanges   []struct {
		Field             string `json:"field,omitempty"`
		ValueBefore       any    `json:"value_before,omitempty"`
		ValueAfter        any    `json:"value_after,omitempty"`
		ValueBeforeParsed string `json:"value_before_parsed,omitempty"`
		ValueAfterParsed  string `json:"value_after_parsed,omitempty"`
		FieldLabel        string `json:"field_label,omitempty"`
	} `json:"field_changes,omitempty"`
	StoryID string `json:"story_id,omitempty"`
}

// GetStoryChanges 获取需求变更历史
// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/story/get_story_changes.html
func (s *StoryService) GetStoryChanges(
	ctx context.Context, request *GetStoryChangesRequest, opts ...RequestOption,
) ([]*StoryChange, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "story_changes", request, opts)
	if err != nil {
		return nil, nil, err
	}

	items := make([]struct {
		WorkitemChange *StoryChange `json:"WorkitemChange"`
	}, 0)
	resp, err := s.client.Do(req, &items)
	if err != nil {
		return nil, resp, err
	}

	changes := make([]*StoryChange, len(items))
	for _, item := range items {
		changes = append(changes, item.WorkitemChange)
	}

	return changes, resp, nil
}

// -----------------------------------------------------------------------------
// 获取需求变更次数
// -----------------------------------------------------------------------------

// -----------------------------------------------------------------------------
// 获取需求自定义字段配置
// -----------------------------------------------------------------------------

type GetStoryCustomFieldsSettingsRequest struct {
	WorkspaceID *int `url:"workspace_id,omitempty"`
}

type StoryCustomFieldsSetting struct {
	ID          string          `json:"id"`
	WorkspaceID string          `json:"workspace_id"`
	EntryType   string          `json:"entry_type"`
	CustomField string          `json:"custom_field"`
	Type        string          `json:"type"`
	Name        string          `json:"name"`
	Options     json.RawMessage `json:"options"`
	Enabled     string          `json:"enabled"`
	Sort        string          `json:"sort"`
	Memo        string          `json:"memo"`
}

// GetStoryCustomFieldsSettings 获取需求自定义字段配置
// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/story/get_story_custom_fields_settings.html
func (s *StoryService) GetStoryCustomFieldsSettings(
	ctx context.Context, request *GetStoryCustomFieldsSettingsRequest, opts ...RequestOption,
) ([]*StoryCustomFieldsSetting, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "stories/custom_fields_settings", request, opts)
	if err != nil {
		return nil, nil, err
	}

	response := make([]struct {
		CustomFieldConfig *StoryCustomFieldsSetting `json:"CustomFieldConfig"`
	}, 0)
	resp, err := s.client.Do(req, &response)
	if err != nil {
		return nil, resp, err
	}

	settings := make([]*StoryCustomFieldsSetting, len(response))
	for _, item := range response {
		settings = append(settings, item.CustomFieldConfig)
	}

	return settings, resp, nil
}

// -----------------------------------------------------------------------------
// 获取需求与测试用例关联关系
// -----------------------------------------------------------------------------

// -----------------------------------------------------------------------------
// 获取需求前后置关系
// -----------------------------------------------------------------------------

// -----------------------------------------------------------------------------
// 批量新增或修改需求前后置关系
// -----------------------------------------------------------------------------

// -----------------------------------------------------------------------------
// 批量删除需求前后置关系
// -----------------------------------------------------------------------------

// -----------------------------------------------------------------------------
// 获取需求保密信息
// -----------------------------------------------------------------------------

// -----------------------------------------------------------------------------
// 批量修改保密信息
// -----------------------------------------------------------------------------

// -----------------------------------------------------------------------------
// 获取需求类别
// -----------------------------------------------------------------------------

// -----------------------------------------------------------------------------
// 更新需求
// -----------------------------------------------------------------------------
type UpdateStoryRequest struct {
	CustomFieldsRequest
	CustomPlanFieldsRequest

	ID              *int           `json:"id"`                           // 必须
	WorkspaceID     *int           `json:"workspace_id"`                 // 必须
	Name            *string        `json:"name,omitempty"`               // 标题
	Priority        *string        `url:"priority,omitempty"`            // 优先级。
	PriorityLabel   *PriorityLabel `url:"priority_label,omitempty"`      // 优先级。推荐使用这个字段
	BusinessValue   *int           `json:"business_value,omitempty"`     // 业务价值
	Status          *string        `json:"status,omitempty"`             // 状态
	VStatus         *string        `json:"v_status,omitempty"`           // 中文状态名称
	Version         *string        `json:"version,omitempty"`            // 版本
	Module          *string        `json:"module,omitempty"`             // 模块
	TestFocus       *string        `json:"test_focus,omitempty"`         // 测试重点
	Size            *int           `json:"size,omitempty"`               // 规模
	Owner           *string        `json:"owner,omitempty"`              // 处理人
	CurrentUser     *string        `json:"current_user,omitempty"`       // 变更人
	CC              *string        `json:"cc,omitempty"`                 // 抄送人
	Developer       *string        `json:"developer,omitempty"`          // 开发人员
	Begin           *string        `json:"begin,omitempty"`              // 预计开始
	Due             *string        `json:"due,omitempty"`                // 预计结束
	IterationID     *string        `json:"iteration_id,omitempty"`       // 迭代ID
	Effort          *string        `json:"effort,omitempty"`             // 预估工时
	EffortCompleted *string        `json:"effort_completed,omitempty"`   // 完成工时
	Remain          *float64       `json:"remain,omitempty"`             // 剩余工时
	Exceed          *float64       `json:"exceed,omitempty"`             // 超出工时
	CategoryID      *int           `json:"category_id,omitempty"`        // 需求分类ID
	ReleaseID       *int           `json:"release_id,omitempty"`         // 发布计划ID
	Source          *string        `json:"source,omitempty"`             // 来源
	Type            *string        `json:"type,omitempty"`               // 类型
	Description     *string        `json:"description,omitempty"`        // 详细描述
	IsAutoCloseTask *int           `json:"is_auto_close_task,omitempty"` // 自动关闭关联任务
	Label           *string        `json:"label,omitempty"`              // 标签
}

// UpdateStory 更新需求
// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/story/update_story.html
func (s *StoryService) UpdateStory(
	ctx context.Context, request *UpdateStoryRequest, opts ...RequestOption,
) (*Story, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodPost, "stories", request, opts)
	if err != nil {
		return nil, nil, err
	}

	var response struct {
		Story *Story `json:"story"`
	}
	resp, err := s.client.Do(req, &response)
	if err != nil {
		return nil, resp, err
	}

	return response.Story, resp, nil
}

// -----------------------------------------------------------------------------
// 更新需求的需求类别
// -----------------------------------------------------------------------------

// -----------------------------------------------------------------------------
// 获取需求所有字段及候选值
// -----------------------------------------------------------------------------

// -----------------------------------------------------------------------------
// 获取需求所有字段的中英文
// -----------------------------------------------------------------------------

// -----------------------------------------------------------------------------
// 获取需求模板列表
// -----------------------------------------------------------------------------

// -----------------------------------------------------------------------------
// 获取需求模板字段
// -----------------------------------------------------------------------------

// -----------------------------------------------------------------------------
// 更新需求分类
// -----------------------------------------------------------------------------

// -----------------------------------------------------------------------------
// 获取回收站下的需求
// -----------------------------------------------------------------------------

// -----------------------------------------------------------------------------
// 获取需求关联的缺陷
// -----------------------------------------------------------------------------

type GetStoryRelatedBugsRequest struct {
	WorkspaceID *int            `url:"workspace_id,omitempty"`
	StoryID     *MultiType[int] `url:"story_id,omitempty"`
}

type StoryRelatedBug struct {
	WorkspaceID int    `json:"workspace_id"`
	StoryID     string `json:"story_id"`
	BugID       string `json:"bug_id"`
}

// GetStoryRelatedBugs 获取需求关联的缺陷
// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/story/get_story_related_bugs.html
func (s *StoryService) GetStoryRelatedBugs(
	ctx context.Context, request *GetStoryRelatedBugsRequest, opts ...RequestOption,
) ([]*StoryRelatedBug, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "stories/get_related_bugs", request, opts)
	if err != nil {
		return nil, nil, err
	}

	bugs := make([]*StoryRelatedBug, 0)
	resp, err := s.client.Do(req, &bugs)
	if err != nil {
		return nil, resp, err
	}

	return bugs, resp, nil
}

// -----------------------------------------------------------------------------
// 解除需求缺陷关联关系
// -----------------------------------------------------------------------------

// -----------------------------------------------------------------------------
// 更新父需求
// -----------------------------------------------------------------------------

// -----------------------------------------------------------------------------
// 创建需求与缺陷关联关系
// -----------------------------------------------------------------------------

// -----------------------------------------------------------------------------
// 创建需求与测试用例关联关系
// -----------------------------------------------------------------------------

// -----------------------------------------------------------------------------
// 获取视图对应的需求列表
// -----------------------------------------------------------------------------

// -----------------------------------------------------------------------------
// 转换需求ID成列表queryToken
// -----------------------------------------------------------------------------

// -----------------------------------------------------------------------------
// 创建需求关联关系
// -----------------------------------------------------------------------------
