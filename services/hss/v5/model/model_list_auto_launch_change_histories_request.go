package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListAutoLaunchChangeHistoriesRequest struct {
	HostId *string `json:"host_id,omitempty"`

	HostIp *string `json:"host_ip,omitempty"`

	HostName *string `json:"host_name,omitempty"`

	AutoLaunchName *string `json:"auto_launch_name,omitempty"`

	Type *int32 `json:"type,omitempty"`

	VariationType *string `json:"variation_type,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	SortKey *string `json:"sort_key,omitempty"`

	SortDir *string `json:"sort_dir,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	Offset *int32 `json:"offset,omitempty"`

	StartTime *int64 `json:"start_time,omitempty"`

	EndTime *int64 `json:"end_time,omitempty"`
}

func (o ListAutoLaunchChangeHistoriesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAutoLaunchChangeHistoriesRequest struct{}"
	}

	return strings.Join([]string{"ListAutoLaunchChangeHistoriesRequest", string(data)}, " ")
}
