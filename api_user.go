package tapd

import (
	"context"
	"net/http"
)

type UserRole struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// UserService is a service to interact with user related API
type UserService struct {
	client *Client
}

// GetRolesRequest represents a request to get roles
type GetRolesRequest struct {
	WorkspaceID *int `url:"workspace_id,omitempty"` // 项目 ID
}

// GetRoles 获取角色ID对照关系
//
// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/user/get_roles.html
func (s *UserService) GetRoles(
	ctx context.Context, request *GetRolesRequest, opts ...RequestOption,
) ([]*UserRole, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "roles", request, opts)
	if err != nil {
		return nil, nil, err
	}

	var items map[string]string
	resp, err := s.client.Do(req, &items)
	if err != nil {
		return nil, resp, err
	}

	roles := make([]*UserRole, 0, len(items))
	for id, name := range items {
		roles = append(roles, &UserRole{ID: id, Name: name})
	}

	return roles, resp, nil
}
