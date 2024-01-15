package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type WtpProtectHostResponseInfo struct {
	HostName *string `json:"host_name,omitempty"`

	HostId *string `json:"host_id,omitempty"`

	PublicIp *string `json:"public_ip,omitempty"`

	PrivateIp *string `json:"private_ip,omitempty"`

	GroupName *string `json:"group_name,omitempty"`

	OsBit *string `json:"os_bit,omitempty"`

	OsType *string `json:"os_type,omitempty"`

	ProtectStatus *string `json:"protect_status,omitempty"`

	RaspProtectStatus *string `json:"rasp_protect_status,omitempty"`

	AntiTamperingTimes *int64 `json:"anti_tampering_times,omitempty"`

	DetectTamperingTimes *int64 `json:"detect_tampering_times,omitempty"`

	LastDetectTime *int64 `json:"last_detect_time,omitempty"`

	ScheduledShutdownStatus *string `json:"scheduled_shutdown_status,omitempty"`

	AgentStatus *string `json:"agent_status,omitempty"`
}

func (o WtpProtectHostResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WtpProtectHostResponseInfo struct{}"
	}

	return strings.Join([]string{"WtpProtectHostResponseInfo", string(data)}, " ")
}
