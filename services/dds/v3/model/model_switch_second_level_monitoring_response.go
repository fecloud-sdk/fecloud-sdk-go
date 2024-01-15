package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type SwitchSecondLevelMonitoringResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SwitchSecondLevelMonitoringResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchSecondLevelMonitoringResponse struct{}"
	}

	return strings.Join([]string{"SwitchSecondLevelMonitoringResponse", string(data)}, " ")
}
