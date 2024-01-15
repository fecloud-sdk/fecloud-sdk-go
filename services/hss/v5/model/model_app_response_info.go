package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type AppResponseInfo struct {
	AgentId *string `json:"agent_id,omitempty"`

	HostId *string `json:"host_id,omitempty"`

	HostName *string `json:"host_name,omitempty"`

	HostIp *string `json:"host_ip,omitempty"`

	AppName *string `json:"app_name,omitempty"`

	Version *string `json:"version,omitempty"`

	UpdateTime *int64 `json:"update_time,omitempty"`

	RecentScanTime *int64 `json:"recent_scan_time,omitempty"`
}

func (o AppResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppResponseInfo struct{}"
	}

	return strings.Join([]string{"AppResponseInfo", string(data)}, " ")
}
