package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// SwitchSecondLevelMonitoringRequest Request Object
type SwitchSecondLevelMonitoringRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *SwitchSecondLevelMonitoringRequestBody `json:"body,omitempty"`
}

func (o SwitchSecondLevelMonitoringRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchSecondLevelMonitoringRequest struct{}"
	}

	return strings.Join([]string{"SwitchSecondLevelMonitoringRequest", string(data)}, " ")
}
