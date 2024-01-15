package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type EventUserResponseInfo struct {
	UserId *int32 `json:"user_id,omitempty"`

	UserGid *int32 `json:"user_gid,omitempty"`

	UserName *string `json:"user_name,omitempty"`

	UserGroupName *string `json:"user_group_name,omitempty"`

	UserHomeDir *string `json:"user_home_dir,omitempty"`

	LoginIp *string `json:"login_ip,omitempty"`

	ServiceType *string `json:"service_type,omitempty"`

	ServicePort *int32 `json:"service_port,omitempty"`

	LoginMode *int32 `json:"login_mode,omitempty"`

	LoginLastTime *int64 `json:"login_last_time,omitempty"`

	LoginFailCount *int32 `json:"login_fail_count,omitempty"`

	PwdHash *string `json:"pwd_hash,omitempty"`

	PwdWithFuzzing *string `json:"pwd_with_fuzzing,omitempty"`

	PwdUsedDays *int32 `json:"pwd_used_days,omitempty"`

	PwdMinDays *int32 `json:"pwd_min_days,omitempty"`

	PwdMaxDays *int32 `json:"pwd_max_days,omitempty"`

	PwdWarnLeftDays *int32 `json:"pwd_warn_left_days,omitempty"`
}

func (o EventUserResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventUserResponseInfo struct{}"
	}

	return strings.Join([]string{"EventUserResponseInfo", string(data)}, " ")
}
