package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type QueryUserResp struct {
	JobId *string `json:"job_id,omitempty"`

	IsGlobalPassword *string `json:"is_global_password,omitempty"`

	Message *string `json:"message,omitempty"`

	UserList *[]QueryUserDetailResp `json:"user_list,omitempty"`

	RolesList *[]QueryRoleDetailResp `json:"roles_list,omitempty"`

	IsSuccess *bool `json:"is_success,omitempty"`
}

func (o QueryUserResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryUserResp struct{}"
	}

	return strings.Join([]string{"QueryUserResp", string(data)}, " ")
}
