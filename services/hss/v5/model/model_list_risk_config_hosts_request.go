package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListRiskConfigHostsRequest struct {
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	CheckName string `json:"check_name"`

	Standard string `json:"standard"`

	HostName *string `json:"host_name,omitempty"`

	HostIp *string `json:"host_ip,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	Offset *int32 `json:"offset,omitempty"`
}

func (o ListRiskConfigHostsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRiskConfigHostsRequest struct{}"
	}

	return strings.Join([]string{"ListRiskConfigHostsRequest", string(data)}, " ")
}
