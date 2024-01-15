package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListAutoLaunchsRequest struct {
	HostId *string `json:"host_id,omitempty"`

	HostName *string `json:"host_name,omitempty"`

	Name *string `json:"name,omitempty"`

	HostIp *string `json:"host_ip,omitempty"`

	Type *string `json:"type,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	Offset *int32 `json:"offset,omitempty"`
}

func (o ListAutoLaunchsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAutoLaunchsRequest struct{}"
	}

	return strings.Join([]string{"ListAutoLaunchsRequest", string(data)}, " ")
}
