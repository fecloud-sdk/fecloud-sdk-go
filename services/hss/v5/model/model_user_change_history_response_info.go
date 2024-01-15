package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UserChangeHistoryResponseInfo struct {
	AgentId *string `json:"agent_id,omitempty"`

	ChangeType *string `json:"change_type,omitempty"`

	HostId *string `json:"host_id,omitempty"`

	HostName *string `json:"host_name,omitempty"`

	PrivateIp *string `json:"private_ip,omitempty"`

	LoginPermission *bool `json:"login_permission,omitempty"`

	RootPermission *bool `json:"root_permission,omitempty"`

	UserGroupName *string `json:"user_group_name,omitempty"`

	UserHomeDir *string `json:"user_home_dir,omitempty"`

	Shell *string `json:"shell,omitempty"`

	UserName *string `json:"user_name,omitempty"`

	ExpireTime *int64 `json:"expire_time,omitempty"`

	RecentScanTime *int64 `json:"recent_scan_time,omitempty"`
}

func (o UserChangeHistoryResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserChangeHistoryResponseInfo struct{}"
	}

	return strings.Join([]string{"UserChangeHistoryResponseInfo", string(data)}, " ")
}
