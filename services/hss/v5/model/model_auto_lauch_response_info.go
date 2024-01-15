package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type AutoLauchResponseInfo struct {
	AgentId *string `json:"agent_id,omitempty"`

	HostId *string `json:"host_id,omitempty"`

	HostName *string `json:"host_name,omitempty"`

	HostIp *string `json:"host_ip,omitempty"`

	Name *string `json:"name,omitempty"`

	Type *int32 `json:"type,omitempty"`

	Path *string `json:"path,omitempty"`

	Hash *string `json:"hash,omitempty"`

	RunUser *string `json:"run_user,omitempty"`

	RecentScanTime *int64 `json:"recent_scan_time,omitempty"`
}

func (o AutoLauchResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AutoLauchResponseInfo struct{}"
	}

	return strings.Join([]string{"AutoLauchResponseInfo", string(data)}, " ")
}
