package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type QuotaResourcesResponseInfo struct {
	ResourceId *string `json:"resource_id,omitempty"`

	Version *string `json:"version,omitempty"`

	QuotaStatus *string `json:"quota_status,omitempty"`

	UsedStatus *string `json:"used_status,omitempty"`

	HostId *string `json:"host_id,omitempty"`

	HostName *string `json:"host_name,omitempty"`

	ChargingMode *string `json:"charging_mode,omitempty"`

	Tags *[]TagInfo `json:"tags,omitempty"`

	ExpireTime *int64 `json:"expire_time,omitempty"`

	SharedQuota *string `json:"shared_quota,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	EnterpriseProjectName *string `json:"enterprise_project_name,omitempty"`
}

func (o QuotaResourcesResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotaResourcesResponseInfo struct{}"
	}

	return strings.Join([]string{"QuotaResourcesResponseInfo", string(data)}, " ")
}
