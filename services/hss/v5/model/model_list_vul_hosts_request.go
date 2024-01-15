package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListVulHostsRequest struct {
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	VulId string `json:"vul_id"`

	Type string `json:"type"`

	HostName *string `json:"host_name,omitempty"`

	HostIp *string `json:"host_ip,omitempty"`

	Status *string `json:"status,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	Offset *int32 `json:"offset,omitempty"`

	AssetValue *string `json:"asset_value,omitempty"`

	GroupName *string `json:"group_name,omitempty"`

	HandleStatus *string `json:"handle_status,omitempty"`

	SeverityLevel *string `json:"severity_level,omitempty"`

	IsAffectBusiness *bool `json:"is_affect_business,omitempty"`
}

func (o ListVulHostsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVulHostsRequest struct{}"
	}

	return strings.Join([]string{"ListVulHostsRequest", string(data)}, " ")
}
