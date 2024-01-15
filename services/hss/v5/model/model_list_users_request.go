package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListUsersRequest struct {
	HostId *string `json:"host_id,omitempty"`

	UserName *string `json:"user_name,omitempty"`

	HostName *string `json:"host_name,omitempty"`

	PrivateIp *string `json:"private_ip,omitempty"`

	LoginPermission *bool `json:"login_permission,omitempty"`

	RootPermission *bool `json:"root_permission,omitempty"`

	UserGroup *string `json:"user_group,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	Offset *int32 `json:"offset,omitempty"`
}

func (o ListUsersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUsersRequest struct{}"
	}

	return strings.Join([]string{"ListUsersRequest", string(data)}, " ")
}
