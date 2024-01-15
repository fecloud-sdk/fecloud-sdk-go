package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DurationStrategies struct {
	Strategy string `json:"strategy"`

	EstimatedUpgradeDuration int32 `json:"estimated_upgrade_duration"`
}

func (o DurationStrategies) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DurationStrategies struct{}"
	}

	return strings.Join([]string{"DurationStrategies", string(data)}, " ")
}
