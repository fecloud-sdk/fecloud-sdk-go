package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListWtpProtectHostRequest struct {
	Region string `json:"region"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	HostName *string `json:"host_name,omitempty"`

	HostId *string `json:"host_id,omitempty"`

	PublicIp *string `json:"public_ip,omitempty"`

	PrivateIp *string `json:"private_ip,omitempty"`

	GroupName *string `json:"group_name,omitempty"`

	OsType *string `json:"os_type,omitempty"`

	ProtectStatus *string `json:"protect_status,omitempty"`

	AgentStatus *string `json:"agent_status,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	Offset *int32 `json:"offset,omitempty"`
}

func (o ListWtpProtectHostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWtpProtectHostRequest struct{}"
	}

	return strings.Join([]string{"ListWtpProtectHostRequest", string(data)}, " ")
}
