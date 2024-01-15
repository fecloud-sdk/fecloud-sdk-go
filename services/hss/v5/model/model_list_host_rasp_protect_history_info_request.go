package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListHostRaspProtectHistoryInfoRequest struct {
	Region string `json:"region"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	HostId string `json:"host_id"`

	StartTime int64 `json:"start_time"`

	EndTime int64 `json:"end_time"`

	Limit int32 `json:"limit"`

	Offset int32 `json:"offset"`

	AlarmLevel *int32 `json:"alarm_level,omitempty"`

	Severity *string `json:"severity,omitempty"`

	ProtectStatus *string `json:"protect_status,omitempty"`
}

func (o ListHostRaspProtectHistoryInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHostRaspProtectHistoryInfoRequest struct{}"
	}

	return strings.Join([]string{"ListHostRaspProtectHistoryInfoRequest", string(data)}, " ")
}
