package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListInstancesBackupStrategyResult struct {
	StartTime string `json:"start_time"`

	KeepDays int32 `json:"keep_days"`
}

func (o ListInstancesBackupStrategyResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstancesBackupStrategyResult struct{}"
	}

	return strings.Join([]string{"ListInstancesBackupStrategyResult", string(data)}, " ")
}
