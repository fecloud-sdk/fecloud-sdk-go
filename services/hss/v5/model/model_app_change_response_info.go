package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type AppChangeResponseInfo struct {
	AgentId *string `json:"agent_id,omitempty"`

	VariationType *string `json:"variation_type,omitempty"`

	HostId *string `json:"host_id,omitempty"`

	AppName *string `json:"app_name,omitempty"`

	HostName *string `json:"host_name,omitempty"`

	HostIp *string `json:"host_ip,omitempty"`

	Version *string `json:"version,omitempty"`

	UpdateTime *int64 `json:"update_time,omitempty"`

	RecentScanTime *int64 `json:"recent_scan_time,omitempty"`
}

func (o AppChangeResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppChangeResponseInfo struct{}"
	}

	return strings.Join([]string{"AppChangeResponseInfo", string(data)}, " ")
}
