package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type AutoLaunchChangeResponseInfo struct {
	AgentId *string `json:"agent_id,omitempty"`

	VariationType *string `json:"variation_type,omitempty"`

	Type *int32 `json:"type,omitempty"`

	HostId *string `json:"host_id,omitempty"`

	HostName *string `json:"host_name,omitempty"`

	HostIp *string `json:"host_ip,omitempty"`

	Path *string `json:"path,omitempty"`

	Hash *string `json:"hash,omitempty"`

	RunUser *string `json:"run_user,omitempty"`

	Name *string `json:"name,omitempty"`

	RecentScanTime *int64 `json:"recent_scan_time,omitempty"`
}

func (o AutoLaunchChangeResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AutoLaunchChangeResponseInfo struct{}"
	}

	return strings.Join([]string{"AutoLaunchChangeResponseInfo", string(data)}, " ")
}
