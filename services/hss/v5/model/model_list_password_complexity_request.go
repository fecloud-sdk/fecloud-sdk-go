package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListPasswordComplexityRequest struct {
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	HostName *string `json:"host_name,omitempty"`

	HostIp *string `json:"host_ip,omitempty"`

	HostId *string `json:"host_id,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	Offset *int32 `json:"offset,omitempty"`
}

func (o ListPasswordComplexityRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPasswordComplexityRequest struct{}"
	}

	return strings.Join([]string{"ListPasswordComplexityRequest", string(data)}, " ")
}
