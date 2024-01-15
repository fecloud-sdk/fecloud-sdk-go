package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UserResponseInfo struct {
	AgentId *string `json:"agent_id,omitempty"`

	HostId *string `json:"host_id,omitempty"`

	HostName *string `json:"host_name,omitempty"`

	HostIp *string `json:"host_ip,omitempty"`

	UserName *string `json:"user_name,omitempty"`

	LoginPermission *bool `json:"login_permission,omitempty"`

	RootPermission *bool `json:"root_permission,omitempty"`

	UserGroupName *string `json:"user_group_name,omitempty"`

	UserHomeDir *string `json:"user_home_dir,omitempty"`

	Shell *string `json:"shell,omitempty"`

	ExpireTime *int64 `json:"expire_time,omitempty"`

	RecentScanTime *int64 `json:"recent_scan_time,omitempty"`
}

func (o UserResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserResponseInfo struct{}"
	}

	return strings.Join([]string{"UserResponseInfo", string(data)}, " ")
}
