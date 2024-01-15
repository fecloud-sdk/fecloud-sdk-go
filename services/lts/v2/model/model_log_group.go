package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type LogGroup struct {
	CreationTime *int64 `json:"creation_time,omitempty"`

	LogGroupName *string `json:"log_group_name,omitempty"`

	LogGroupId *string `json:"log_group_id,omitempty"`

	TtlInDays *int32 `json:"ttl_in_days,omitempty"`

	Tag map[string]string `json:"tag,omitempty"`
}

func (o LogGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LogGroup struct{}"
	}

	return strings.Join([]string{"LogGroup", string(data)}, " ")
}
