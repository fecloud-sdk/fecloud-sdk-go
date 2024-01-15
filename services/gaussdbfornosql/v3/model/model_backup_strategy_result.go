package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BackupStrategyResult struct {
	StartTime *string `json:"start_time,omitempty"`

	KeepDays *string `json:"keep_days,omitempty"`
}

func (o BackupStrategyResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupStrategyResult struct{}"
	}

	return strings.Join([]string{"BackupStrategyResult", string(data)}, " ")
}
