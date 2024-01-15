package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type HostRaspProtectHistoryResponseInfo struct {
	HostIp *string `json:"host_ip,omitempty"`

	HostName *string `json:"host_name,omitempty"`

	AlarmTime *int64 `json:"alarm_time,omitempty"`

	ThreatType *string `json:"threat_type,omitempty"`

	AlarmLevel *int32 `json:"alarm_level,omitempty"`

	SourceIp *string `json:"source_ip,omitempty"`

	AttackedUrl *string `json:"attacked_url,omitempty"`
}

func (o HostRaspProtectHistoryResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HostRaspProtectHistoryResponseInfo struct{}"
	}

	return strings.Join([]string{"HostRaspProtectHistoryResponseInfo", string(data)}, " ")
}
