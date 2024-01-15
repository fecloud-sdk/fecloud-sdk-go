package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowBackupPolicyResult struct {
	KeepDays int32 `json:"keep_days"`

	StartTime *string `json:"start_time,omitempty"`

	Period *string `json:"period,omitempty"`
}

func (o ShowBackupPolicyResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBackupPolicyResult struct{}"
	}

	return strings.Join([]string{"ShowBackupPolicyResult", string(data)}, " ")
}
