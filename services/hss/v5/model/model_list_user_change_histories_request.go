package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListUserChangeHistoriesRequest struct {
	UserName *string `json:"user_name,omitempty"`

	HostId *string `json:"host_id,omitempty"`

	RootPermission *bool `json:"root_permission,omitempty"`

	HostName *string `json:"host_name,omitempty"`

	PrivateIp *string `json:"private_ip,omitempty"`

	ChangeType *string `json:"change_type,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	Offset *int32 `json:"offset,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	StartTime *int64 `json:"start_time,omitempty"`

	EndTime *int64 `json:"end_time,omitempty"`
}

func (o ListUserChangeHistoriesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUserChangeHistoriesRequest struct{}"
	}

	return strings.Join([]string{"ListUserChangeHistoriesRequest", string(data)}, " ")
}
