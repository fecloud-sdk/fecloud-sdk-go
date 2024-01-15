package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type SwitchSecondLevelMonitoringRequestBody struct {
	Enabled bool `json:"enabled"`
}

func (o SwitchSecondLevelMonitoringRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchSecondLevelMonitoringRequestBody struct{}"
	}

	return strings.Join([]string{"SwitchSecondLevelMonitoringRequestBody", string(data)}, " ")
}
