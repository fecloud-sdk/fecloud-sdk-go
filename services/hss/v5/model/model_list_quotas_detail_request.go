package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListQuotasDetailRequest struct {
	Region *string `json:"region,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Version *string `json:"version,omitempty"`

	Category *string `json:"category,omitempty"`

	QuotaStatus *string `json:"quota_status,omitempty"`

	UsedStatus *string `json:"used_status,omitempty"`

	HostName *string `json:"host_name,omitempty"`

	ResourceId *string `json:"resource_id,omitempty"`

	ChargingMode *string `json:"charging_mode,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	Offset *int32 `json:"offset,omitempty"`
}

func (o ListQuotasDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQuotasDetailRequest struct{}"
	}

	return strings.Join([]string{"ListQuotasDetailRequest", string(data)}, " ")
}
