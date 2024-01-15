package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ApplyHistoryInfo struct {
	InstanceId string `json:"instance_id"`

	InstanceName string `json:"instance_name"`

	AppliedAt string `json:"applied_at"`

	ApplyResult string `json:"apply_result"`

	FailureReason *string `json:"failure_reason,omitempty"`
}

func (o ApplyHistoryInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplyHistoryInfo struct{}"
	}

	return strings.Join([]string{"ApplyHistoryInfo", string(data)}, " ")
}
