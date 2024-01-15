package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BackupStrategyOption struct {
	StartTime string `json:"start_time"`

	KeepDays *string `json:"keep_days,omitempty"`
}

func (o BackupStrategyOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupStrategyOption struct{}"
	}

	return strings.Join([]string{"BackupStrategyOption", string(data)}, " ")
}
