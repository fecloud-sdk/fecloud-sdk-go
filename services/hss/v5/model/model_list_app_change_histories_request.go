package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListAppChangeHistoriesRequest struct {
	HostId *string `json:"host_id,omitempty"`

	HostIp *string `json:"host_ip,omitempty"`

	HostName *string `json:"host_name,omitempty"`

	AppName *string `json:"app_name,omitempty"`

	VariationType *string `json:"variation_type,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	SortKey *string `json:"sort_key,omitempty"`

	SortDir *string `json:"sort_dir,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	Offset *int32 `json:"offset,omitempty"`

	StartTime *int64 `json:"start_time,omitempty"`

	EndTime *int64 `json:"end_time,omitempty"`
}

func (o ListAppChangeHistoriesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAppChangeHistoriesRequest struct{}"
	}

	return strings.Join([]string{"ListAppChangeHistoriesRequest", string(data)}, " ")
}
