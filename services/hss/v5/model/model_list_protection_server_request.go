package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListProtectionServerRequest struct {
	Region string `json:"region"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Offset *int32 `json:"offset,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	HostName *string `json:"host_name,omitempty"`

	OsType *string `json:"os_type,omitempty"`

	HostIp *string `json:"host_ip,omitempty"`

	HostStatus *string `json:"host_status,omitempty"`

	LastDays *int32 `json:"last_days,omitempty"`
}

func (o ListProtectionServerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProtectionServerRequest struct{}"
	}

	return strings.Join([]string{"ListProtectionServerRequest", string(data)}, " ")
}
