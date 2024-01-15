package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListRiskConfigsRequest struct {
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	CheckName *string `json:"check_name,omitempty"`

	GroupId *string `json:"group_id,omitempty"`

	Severity *string `json:"severity,omitempty"`

	Standard *string `json:"standard,omitempty"`

	HostId *string `json:"host_id,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	Offset *int32 `json:"offset,omitempty"`
}

func (o ListRiskConfigsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRiskConfigsRequest struct{}"
	}

	return strings.Join([]string{"ListRiskConfigsRequest", string(data)}, " ")
}
