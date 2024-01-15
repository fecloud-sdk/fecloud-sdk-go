package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListHostProtectHistoryInfoRequest struct {
	Region string `json:"region"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	HostId string `json:"host_id"`

	StartTime int64 `json:"start_time"`

	EndTime int64 `json:"end_time"`

	Limit int32 `json:"limit"`

	Offset int32 `json:"offset"`
}

func (o ListHostProtectHistoryInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHostProtectHistoryInfoRequest struct{}"
	}

	return strings.Join([]string{"ListHostProtectHistoryInfoRequest", string(data)}, " ")
}
