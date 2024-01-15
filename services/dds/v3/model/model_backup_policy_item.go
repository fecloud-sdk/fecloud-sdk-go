package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BackupPolicyItem struct {
	KeepDays int32 `json:"keep_days"`

	StartTime *string `json:"start_time,omitempty"`

	Period *string `json:"period,omitempty"`
}

func (o BackupPolicyItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupPolicyItem struct{}"
	}

	return strings.Join([]string{"BackupPolicyItem", string(data)}, " ")
}
