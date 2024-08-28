package tapd

import (
	"context"
	"net/http"
)

type WorkspaceService struct {
	client *Client
}

func NewWorkspaceService(client *Client) *WorkspaceService {
	return &WorkspaceService{client}
}

// 获取子项目信息
// 获取项目信息
// 获取指定项目成员
// 添加项目成员
// 获取公司项目列表
// 获取用户组ID对照关系
// 获取用户参与的项目列表
// 获取项目成员列表
// 获取项目自定义字段
// 更新项目信息
// 获取项目文档

type GetMemberActivityLogRequest struct {
	// [必须]项目 id 为公司id则查询所有项目
	WorkspaceID *int `url:"workspace_id,omitempty"`

	// [可选]为1则仅返回公司级活动日志 要求workspace_id=公司id & company_only=1
	CompanyOnly *int `url:"company_only,omitempty"`

	// [可选]设置返回数量限制，默认为20
	Limit *int `url:"limit,omitempty"`

	// [可选]返回当前数量限制下第N页的数据，默认为1（第一页）
	Page *int `url:"page,omitempty"`

	// [可选]起始时间，精确到分钟，格式为Y-m-d H:i 只能查最近半年内的数据
	StartTime *string `url:"start_time,omitempty"`

	// [可选]终止时间，精确到分钟，格式为Y-m-d H:i 只能查最近半年内的数据
	EndTime *string `url:"end_time,omitempty"`

	// [可选]操作人昵称
	Operator *string `url:"operator,omitempty"`

	// [可选]操作类型，默认为所有，可以填写add,delete,download,upload中的一个
	OperateType *OperateType `url:"operate_type,omitempty"`

	// [可选]操作对象，默认为所有，可以填写attachment,board,bug,document,iteration,
	// launch,member_activity_log,release,story,task,tcase,testplan,wiki中的一个
	OperatorObject *OperateObject `url:"operator_object,omitempty"`

	// [可选]请求IP条件，严格匹配
	IP *string `url:"ip,omitempty"`
}

type MemberActivityLog struct {
	ID            string        `json:"id,omitempty"`
	Action        string        `json:"action,omitempty"`
	Created       string        `json:"created,omitempty"`
	Creator       string        `json:"creator,omitempty"`
	ProjectName   string        `json:"project_name,omitempty"`
	OperateType   OperateType   `json:"operate_type,omitempty"`
	OperateObject OperateObject `json:"operate_object,omitempty"`
	Title         string        `json:"title,omitempty"`
	URL           string        `json:"url,omitempty"`
	IP            string        `json:"ip,omitempty"`
	UA            string        `json:"ua,omitempty"`
}

type GetMemberActivityLogResponse struct {
	PerPage      string               `json:"perPage"`
	TotalItems   int                  `json:"totalItems"`
	CurrentPage  string               `json:"currentPage"`
	Records      []*MemberActivityLog `json:"records"`
	OperateTypes struct {
		Add      string `json:"add"`
		Delete   string `json:"delete"`
		Upload   string `json:"upload"`
		Download string `json:"download"`
	} `json:"operate_types"`
	OperateObjects struct {
		Board             string `json:"board"`
		Story             string `json:"story"`
		Bug               string `json:"bug"`
		Iteration         string `json:"iteration"`
		Wiki              string `json:"wiki"`
		Document          string `json:"document"`
		Attachment        string `json:"attachment"`
		Task              string `json:"task"`
		Tcase             string `json:"tcase"`
		Testplan          string `json:"testplan"`
		Launch            string `json:"launch"`
		Release           string `json:"release"`
		MemberActivityLog string `json:"member_activity_log"`
	} `json:"operate_objects"`
}

// GetMemberActivityLog 获取成员活动日志
//
// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/workspace/member_activity_log.html
func (s *WorkspaceService) GetMemberActivityLog(
	ctx context.Context, request *GetMemberActivityLogRequest, opts ...RequestOption,
) (*GetMemberActivityLogResponse, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "workspaces/member_activity_log", request, opts)
	if err != nil {
		return nil, nil, err
	}

	response := new(GetMemberActivityLogResponse)
	resp, err := s.client.Do(req, &response)
	if err != nil {
		return nil, resp, err
	}

	return response, resp, nil
}
