package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BackupStrategyForItemResponse struct {
	StartTime string `json:"start_time"`

	KeepDays int32 `json:"keep_days"`
}

func (o BackupStrategyForItemResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupStrategyForItemResponse struct{}"
	}

	return strings.Join([]string{"BackupStrategyForItemResponse", string(data)}, " ")
}
